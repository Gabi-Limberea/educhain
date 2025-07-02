package pgsql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/lib/pq"

	_ "github.com/lib/pq"
)

const (
	connStr                       = "user=%s password=%s host=%s port=%s dbname=%s sslmode=disable"
	uniqueConstraintViolationCode = "23505"
)

type DB struct {
	conn *sql.DB
}

// NewDB creates a new pgsql.DB object
func NewDB() *DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("postgres", fmt.Sprintf(connStr, user, password, host, port, dbName))
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("connected to database", "db-name", dbName, "host", host, "port", port, "user", user)
	return &DB{conn: db}
}

// Close closes the connection to the database
func (d *DB) Close() {
	err := d.conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// Query returns the rows retrieved from the database
func (d *DB) Query(rawStmt string) (*sql.Rows, error) {
	stmt, err := d.conn.Prepare(rawStmt)
	if err != nil {
		return nil, err
	}

	return stmt.Query()
}

// Exec returns a summary of the query on the database
func (d *DB) Exec(rawStmt string) (sql.Result, error) {
	stmt, err := d.conn.Prepare(rawStmt)
	if err != nil {
		return nil, err
	}

	return stmt.Exec()
}

// TranslatePQError transforms the original Postgresql error into a human-readable one to be redirected to the user
func TranslatePQError(err error) (int, error) {
	if err == nil {
		return http.StatusOK, nil
	}

	var pqErr *pq.Error
	ok := errors.As(err, &pqErr)
	if !ok {
		return http.StatusInternalServerError, err
	}

	if pqErr.Code == uniqueConstraintViolationCode {
		obj, _ := strings.CutSuffix(pqErr.Table, "s")
		return http.StatusConflict, fmt.Errorf("%s already exists", obj)
	}

	return http.StatusInternalServerError, err
}
