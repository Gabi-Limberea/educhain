package handler

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"smart-contract-api/eth"
	"smart-contract-api/ipfs"
	"smart-contract-api/models"
	"smart-contract-api/pgsql"
	"strings"
	"time"

	"github.com/Gabi-Limberea/educhain/ipfs-api/constants"
	"github.com/Gabi-Limberea/educhain/ipfs-api/pkg/zip"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

const (
	defaultTokenTTL = 30
)

type Handler struct {
	db         *pgsql.DB
	ethClient  *eth.Client
	ipfsClient *ipfs.Client
	signingKey []byte
}

func NewHandler() *Handler {
	signingKey := os.Getenv("JWT_SIGNING_KEY")
	if signingKey == "" {
		log.Fatal("JWT_SIGNING_KEY environment variable not set")
	}

	return &Handler{
		db:         pgsql.NewDB(),
		ethClient:  eth.NewClient(),
		ipfsClient: ipfs.NewClient(),
		signingKey: []byte(signingKey),
	}
}

func (h *Handler) Close() {
	h.db.Close()
	h.ethClient.Close()
}

func (h *Handler) getJWT(r *http.Request) (*AuthToken, error) {
	var auth *AuthToken
	header, ok := r.Header["Authorization"]
	if !ok {
		slog.Error("no JWT token used")
		return nil, fmt.Errorf("no JWT token used")
	}

	strToken, _ := strings.CutPrefix(header[0], "Bearer ")
	token, err := jwt.Parse(
		strToken, func(token *jwt.Token) (interface{}, error) {
			return h.signingKey, nil
		},
	)
	if err != nil {
		slog.Error("failed to parse token", "error", err)
		return nil, fmt.Errorf("failed to parse token")
	}

	if !token.Valid {
		slog.Error("invalid token")
		return nil, fmt.Errorf("invalid token")
	}

	auth, err = FromJWT(token)
	if err != nil {
		slog.Error("failed to convert token", "error", err)
		return nil, fmt.Errorf("invalid token")
	}

	return auth, nil
}

func (h *Handler) checkAuthorization(r *http.Request) (int64, bool) {
	tok, err := h.getJWT(r)
	if err != nil {
		slog.Error("failed to retrieve auth token", "error", err)
		return 0, false
	}

	pid := tok.PID
	selectStmt := fmt.Sprintf("select pid from providers where pid = %d", pid)

	res, err := h.db.Query(selectStmt)
	if err != nil {
		slog.Error("failed to query provider", "error", err)
		return pid, false
	}
	defer res.Close()

	if !res.Next() {
		if err = res.Err(); err != nil {
			slog.Error("failed to query provider", "error", err)
			return pid, false
		}

		slog.Error("authorization failed")
		return pid, false
	}

	return pid, true
}

func (h *Handler) NotImplemented(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *Handler) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// ProviderRegister registers a new certification provider into the system
func (h *Handler) ProviderRegister(w http.ResponseWriter, r *http.Request) {
	wr := newWriter(w)

	var provider models.Provider
	var err error

	if err = json.NewDecoder(r.Body).Decode(&provider); err != nil {
		slog.Error("failed to decode request body", "error", err)
		wr.Error(http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if !provider.IsValid() {
		slog.Error("invalid request body")
		wr.Error(http.StatusBadRequest, fmt.Errorf("invalid request body"))
		return
	}

	provider.ContractAddress, err = h.ethClient.NewContract()
	if err != nil {
		slog.Error("failed to deploy new provider smart contract", "error", err)
		wr.Error(http.StatusBadRequest, fmt.Errorf("failed to deploy new provider smart contract"))
		return
	}

	insertStmt := fmt.Sprintf(
		"insert into providers(email, passwd, contract_address, name, address, phone_number, website) values %s;",
		provider.ToSQL(),
	)
	_, err = h.db.Exec(insertStmt)
	if err != nil {
		slog.Error("failed to insert provider", "error", err)
		wr.Error(pgsql.TranslatePQError(err))
		return
	}

	wr.Message(http.StatusCreated, "provider registered successfully")
}

func (h *Handler) GetProviderInfo(w http.ResponseWriter, r *http.Request) {
	wr := newWriter(w)

	tok, err := h.getJWT(r)
	if err != nil {
		slog.Error("failed to retrieve auth token", "error", err)
		wr.Message(http.StatusUnauthorized, "unauthorized")
		return
	}

	selectStmt := fmt.Sprintf(
		"select email, contract_address, name, address, phone_number, website from providers where pid = %d",
		tok.PID,
	)

	res, err := h.db.Query(selectStmt)
	if err != nil {
		slog.Error("failed to query provider", "error", err)
		wr.Error(pgsql.TranslatePQError(err))
		return
	}
	defer res.Close()

	if !res.Next() {
		if err = res.Err(); err != nil {
			slog.Error("failed to query provider", "error", err)
			wr.Error(pgsql.TranslatePQError(err))
			return
		}

		slog.Error("unauthorized request", "method", r.Method, "path", r.URL.Path)
		wr.Error(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	provider := &models.Provider{}
	if err = res.Scan(
		&provider.Email, &provider.ContractAddress, &provider.OrganizationInfo.Name,
		&provider.OrganizationInfo.Address, &provider.OrganizationInfo.PhoneNumber,
		&provider.OrganizationInfo.Website,
	); err != nil {
		slog.Error("failed to parse query result", "error", err)
		wr.Error(pgsql.TranslatePQError(err))
		return
	}

	wr.Object(http.StatusOK, provider)
}

// ProviderTokenGenerate generates a new JWT token for provider authorization
func (h *Handler) ProviderTokenGenerate(w http.ResponseWriter, r *http.Request) {
	wr := newWriter(w)

	email, password, ok := r.BasicAuth()
	if !ok {
		slog.Error("basic auth required")
		wr.Error(http.StatusUnauthorized, fmt.Errorf("basic auth required"))
		return
	}

	selectStmt := fmt.Sprintf(
		"select pid from providers where email = '%s' and passwd = '%s';", email,
		base64.StdEncoding.EncodeToString([]byte(password)),
	)
	res, err := h.db.Query(selectStmt)
	if err != nil {
		slog.Error("failed to query provider", "error", err)
		wr.Error(pgsql.TranslatePQError(err))
		return
	}
	defer res.Close()

	if !res.Next() {
		if err = res.Err(); err != nil {
			slog.Error("failed to query provider", "error", err)
			wr.Error(pgsql.TranslatePQError(err))
			return
		}

		slog.Error("wrong credentials")
		wr.Error(http.StatusUnauthorized, fmt.Errorf("wrong credentials"))
		return
	}

	var pid int64
	if err = res.Scan(&pid); err != nil {
		slog.Error("failed to parse query result", "error", err)
		wr.Error(pgsql.TranslatePQError(err))
		return
	}

	authToken := &AuthToken{}
	json.NewDecoder(r.Body).Decode(&authToken)
	if authToken.TTL == 0 {
		authToken.TTL = defaultTokenTTL
	}
	authToken.PID = pid
	authToken.IssuedAt = time.Now()
	authToken.ExpirationTime = authToken.IssuedAt.Add(time.Hour * 24 * authToken.TTL)
	authToken.Issuer = "EduChain"

	tok, err := authToken.ToJWT().SignedString(h.signingKey)
	if err != nil {
		slog.Error("failed to sign token", "error", err)
		wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to generate token"))
		return
	}

	wr.Object(http.StatusOK, tok)
}

// StudentsEnroll registers a batch of students for the provider
func (h *Handler) StudentsEnroll(w http.ResponseWriter, r *http.Request) {
	wr := newWriter(w)
	pid, authorized := h.checkAuthorization(r)
	if !authorized {
		slog.Info("unauthorized request", "method", r.Method, "path", r.URL.Path)
		wr.Message(http.StatusUnauthorized, "unauthorized")
		return
	}

	var students []models.Student
	if err := json.NewDecoder(r.Body).Decode(&students); err != nil {
		slog.Error("failed to parse body", "error", err)
		wr.Error(http.StatusBadRequest, fmt.Errorf("failed to parse body"))
		return
	}
	if len(students) == 0 {
		slog.Error("no students provided")
		wr.Error(http.StatusBadRequest, fmt.Errorf("no students provided"))
		return
	}

	studentValues := make([]string, len(students))
	for i, student := range students {
		student.Provider = pid
		studentValues[i] = student.ToSQL()
	}

	insertStmt := fmt.Sprintf(
		"insert into students(external_id, wallet_address, provider) values %s",
		strings.Join(studentValues, ", "),
	)
	if _, err := h.db.Exec(insertStmt); err != nil {
		slog.Error("failed to insert students", "error", err)
		wr.Error(pgsql.TranslatePQError(err))
		return
	}

	wr.Message(http.StatusCreated, "successfully enrolled students")
}

// StudentsGet retrieves all students registered for the provider
func (h *Handler) StudentsGet(w http.ResponseWriter, r *http.Request) {
	wr := newWriter(w)
	pid, authorized := h.checkAuthorization(r)
	if !authorized {
		slog.Info("unauthorized request", "method", r.Method, "path", r.URL.Path)
		wr.Error(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	selectStmt := fmt.Sprintf(
		"select external_id, wallet_address  from students where provider = %d", pid,
	)
	res, err := h.db.Query(selectStmt)
	if err != nil {
		slog.Error("failed to query students", "error", err)
		wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to query students"))
		return
	}
	defer res.Close()

	var students = make([]models.Student, 0)
	for res.Next() {
		var student models.Student
		if err = res.Scan(
			&student.ExternalID, &student.WalletAddress,
		); err != nil {
			slog.Error("failed to parse query results", "error", err)
			wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to retrieve students"))
			return
		}

		students = append(students, student)
	}

	wr.Object(http.StatusOK, students)
}

func (h *Handler) getStudent(externalID string, pid int64) (*sql.Rows, error) {
	selectStmt := fmt.Sprintf(
		"select external_id, wallet_address from students where external_id = '%s' and provider = %d",
		externalID, pid,
	)
	res, err := h.db.Query(selectStmt)
	if err != nil {
		slog.Error("failed to query students", "error", err)
		return nil, fmt.Errorf("failed to query students: %w", err)
	}

	return res, nil
}

func (h *Handler) studentExists(externalID string, pid int64) (bool, error) {
	res, err := h.getStudent(externalID, pid)
	if err != nil {
		return false, err
	}
	defer res.Close()

	if !res.Next() {
		if err = res.Err(); err != nil {
			slog.Error("failed to query student", "error", err)
			return false, fmt.Errorf("failed to query student: %w", err)
		}
		return false, nil
	}

	return true, nil
}

func (h *Handler) StudentGetByID(w http.ResponseWriter, r *http.Request) {
	wr := newWriter(w)
	pid, authorized := h.checkAuthorization(r)
	if !authorized {
		slog.Info("unauthorized request", "method", r.Method, "path", r.URL.Path)
		wr.Error(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	externalID, ok := mux.Vars(r)["student_id"]
	if !ok {
		slog.Error("no student ID provided")
		wr.Error(http.StatusBadRequest, fmt.Errorf("no student ID provided"))
		return
	}

	res, err := h.getStudent(externalID, pid)
	if err != nil {
		slog.Error("failed to query student", "error", err)
		wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to retrieve student"))
		return
	}
	defer res.Close()

	if !res.Next() {
		if err = res.Err(); err != nil {
			slog.Error("failed to query student", "error", err)
			wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to query student"))
			return
		}

		slog.Info("no student found")
		wr.Error(http.StatusNotFound, fmt.Errorf("student not found"))
		return
	}

	var student models.Student
	if err = res.Scan(&student.ExternalID, &student.WalletAddress); err != nil {
		slog.Error("failed to parse query result", "error", err)
		wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to retrieve student"))
		return
	}

	wr.Object(http.StatusOK, student)
}

func (h *Handler) DiplomasGet(w http.ResponseWriter, r *http.Request) {
	wr := newWriter(w)
	pid, authorized := h.checkAuthorization(r)
	if !authorized {
		slog.Info("unauthorized request", "method", r.Method, "path", r.URL.Path)
		wr.Error(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	externalID, ok := mux.Vars(r)["student_id"]
	if !ok {
		slog.Error("no student ID provided")
		wr.Error(http.StatusBadRequest, fmt.Errorf("no student ID provided"))
		return
	}

	if exists, err := h.studentExists(externalID, pid); err != nil {
		slog.Error("failed to check student existence", "err", err)
		wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to retrieve student"))
		return
	} else if !exists {
		slog.Info("no student found")
		wr.Error(http.StatusNotFound, fmt.Errorf("student not found"))
		return
	}

	selectStmt := fmt.Sprintf(
		"select did, student from diplomas where student = '%s' and provider = %d", externalID, pid,
	)
	res, err := h.db.Query(selectStmt)
	if err != nil {
		slog.Error("failed to query students", "error", err)
		wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to retrieve diplomas"))
		return
	}
	defer res.Close()

	diplomas := make([]models.Diploma, 0)
	for res.Next() {
		diploma := models.Diploma{}
		if err = res.Scan(&diploma.TokenID, &diploma.Student); err != nil {
			slog.Error("failed to parse query result", "error", err)
			wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to retrieve diplomas"))
			return
		}

		diplomas = append(diplomas, diploma)
	}

	wr.Object(http.StatusOK, diplomas)
}

func (h *Handler) uploadDiplomaForStudent(
	student models.Student, pid int64, file io.ReadCloser, fileName string,
) (*models.Diploma, error) {
	uploadedFile, err := h.ipfsClient.UploadFile(file, fileName)
	if err != nil {
		slog.Error("error uploading file", "msg", err)
		return nil, fmt.Errorf("failed to upload file")
	}

	selectStmt := fmt.Sprintf(
		"select name, address, contract_address from providers where pid = %d", pid,
	)
	res, err := h.db.Query(selectStmt)
	if err != nil {
		slog.Error("failed to query provider", "error", err)
		return nil, fmt.Errorf("failed to query provider: %w", err)
	}

	defer res.Close()
	res.Next()

	var provider models.Provider
	if err = res.Scan(
		&provider.OrganizationInfo.Name, &provider.OrganizationInfo.Address,
		&provider.ContractAddress,
	); err != nil {
		slog.Error("failed to parse query result", "error", err)
		return nil, fmt.Errorf("something went wrong")
	}

	metadataFile, err := eth.GenerateNFTMetadataFile(uploadedFile, provider)
	if err != nil {
		slog.Error("failed to generate metadata file", "error", err)
		return nil, fmt.Errorf("something went wrong")
	}

	uploadedMetadata, err := h.ipfsClient.UploadFile(
		metadataFile, fmt.Sprintf("%s.json", uploadedFile.Hash),
	)
	if err != nil {
		slog.Error("failed to upload metadata file", "error", err)
		return nil, fmt.Errorf("something went wrong")
	}

	tokenID, err := h.ethClient.MintNFT(
		fmt.Sprintf("ipfs.io/ipfs/%s", uploadedMetadata.Hash), provider.ContractAddress,
		student.WalletAddress,
	)
	if err != nil {
		slog.Error("failed to mint nft", "error", err)
		return nil, fmt.Errorf("failed to mint nft")
	}

	return &models.Diploma{
		TokenID:  *tokenID,
		Student:  student.ExternalID,
		Provider: pid,
	}, nil
}

func (h *Handler) DiplomaUpload(w http.ResponseWriter, r *http.Request) {
	wr := newWriter(w)
	pid, authorized := h.checkAuthorization(r)
	if !authorized {
		slog.Info("unauthorized request", "method", r.Method, "path", r.URL.Path)
		wr.Error(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	externalID, ok := mux.Vars(r)["student_id"]
	if !ok {
		slog.Error("no student ID provided")
		wr.Error(http.StatusBadRequest, fmt.Errorf("no student ID provided"))
		return
	}

	rows, err := h.getStudent(externalID, pid)
	if err != nil {
		slog.Error("failed to check student existence", "error", err)
		wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to retrieve student"))
		return
	}

	defer rows.Close()
	if !rows.Next() {
		if err = rows.Err(); err != nil {
			slog.Error("failed to query student", "error", err)
			wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to query student"))
			return
		}

		slog.Info("no student found")
		wr.Error(http.StatusNotFound, fmt.Errorf("student not found"))
		return
	}

	var student models.Student
	if err = rows.Scan(&student.ExternalID, &student.WalletAddress); err != nil {
		slog.Error("failed to parse query result", "error", err)
		wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to query student"))
		return
	}

	if err := r.ParseMultipartForm(constants.MB512); err != nil {
		slog.Error("error parsing form", "msg", err)
		wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to parse form"))
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		slog.Error("error parsing form", "msg", err)
		wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to parse form"))
		return
	}

	switch format := fileHeader.Header.Get("Content-Type"); format {
	case "image/png", "image/jpeg", "application/pdf":
		break
	default:
		slog.Error("wrong file type", "format", format)
		wr.Error(
			http.StatusBadRequest,
			fmt.Errorf("wrong file type, formats accepted are: jpeg, png, pdf"),
		)
		return
	}

	defer file.Close()

	diploma, err := h.uploadDiplomaForStudent(student, pid, file, fileHeader.Filename)
	if err != nil {
		slog.Error("failed to upload diploma for student", "error", err)
		wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to upload diploma for student"))
	}

	insertStmt := fmt.Sprintf(
		"insert into diplomas (did, student, provider) values %s", diploma.ToSQL(),
	)
	_, err = h.db.Exec(insertStmt)
	if err != nil {
		slog.Error("failed to insert diploma", "error", err)
		wr.Error(pgsql.TranslatePQError(err))
		return
	}

	wr.Object(http.StatusCreated, diploma)
}

func (h *Handler) DiplomasBulkUpload(w http.ResponseWriter, r *http.Request) {
	wr := newWriter(w)
	pid, authorized := h.checkAuthorization(r)
	if !authorized {
		slog.Info("unauthorized request", "method", r.Method, "path", r.URL.Path)
		wr.Error(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	if err := r.ParseMultipartForm(constants.MB512); err != nil {
		slog.Error("error parsing form", "msg", err)
		wr.Error(http.StatusBadRequest, fmt.Errorf("failed to parse form"))
		return
	}

	archive, archiveHeader, err := r.FormFile("file")
	if err != nil {
		slog.Error("error parsing form", "msg", err)
		wr.Error(http.StatusBadRequest, fmt.Errorf("failed to parse form"))
		return
	}

	if archiveHeader.Header.Get("Content-Type") != "application/zip" {
		slog.Error("wrong file type", "type", archiveHeader.Header.Get("Content-Type"))
		wr.Error(http.StatusBadRequest, fmt.Errorf("wrong file type, accepted formats: zip"))
		return
	}

	files, err := zip.UnzipArchive(archive, archiveHeader.Size)
	if err != nil {
		slog.Error("error unzipping archive", "error", err)
		wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to unzip archive"))
		return
	}

	diplomas := make([]models.Diploma, len(files))
	sqlDiplomas := make([]string, len(files))
	for i, file := range files {
		rows, err := h.getStudent(file.Name, pid)
		if err != nil {
			slog.Error("failed to check student existence", "error", err)
			wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to retrieve student"))
			return
		}

		if !rows.Next() {
			if err = rows.Err(); err != nil {
				slog.Error("failed to query student", "error", err)
				wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to query student"))
				return
			}

			slog.Info("no student found")
			wr.Error(http.StatusNotFound, fmt.Errorf("student not found"))
			return
		}

		var student models.Student
		if err = rows.Scan(&student.ExternalID, &student.WalletAddress); err != nil {
			slog.Error("failed to parse query result", "error", err)
			wr.Error(http.StatusInternalServerError, fmt.Errorf("failed to query student"))
			return
		}
		rows.Close()

		diploma, err := h.uploadDiplomaForStudent(student, pid, file, file.Name)
		if err != nil {
			slog.Error("failed to upload diploma for student", "error", err)
			wr.Error(
				http.StatusInternalServerError, fmt.Errorf("failed to upload diploma for student"),
			)
		}

		diplomas[i] = *diploma
		sqlDiplomas[i] = diploma.ToSQL()
	}

	insertStmt := fmt.Sprintf(
		"insert into diplomas (did, student, provider) values %s", strings.Join(sqlDiplomas, ", "),
	)
	_, err = h.db.Exec(insertStmt)
	if err != nil {
		slog.Error("failed to insert diplomas", "error", err)
		wr.Error(pgsql.TranslatePQError(err))
		return
	}

	wr.Object(http.StatusCreated, diplomas)
}
