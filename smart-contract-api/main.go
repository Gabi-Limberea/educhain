package main

import (
	"log"
	"log/slog"
	"net/http"
	"smart-contract-api/handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	h := handler.NewHandler()
	defer h.Close()

	router.HandleFunc("/healthz", h.HealthCheck).Methods(http.MethodGet)
	router.HandleFunc("/providers", h.ProviderRegister).Methods(http.MethodPost)
	router.HandleFunc("/providers/token", h.ProviderTokenGenerate).Methods(http.MethodPost)
	router.HandleFunc("/account", h.GetProviderInfo).Methods(http.MethodGet)
	router.HandleFunc("/students", h.StudentsGet).Methods(http.MethodGet)
	router.HandleFunc("/students", h.StudentsEnroll).Methods(http.MethodPost)
	router.HandleFunc("/students/{student_id}", h.StudentGetByID).Methods(http.MethodGet)
	router.HandleFunc("/students/{student_id}", h.NotImplemented).Methods(http.MethodPut)
	router.HandleFunc("/students/{student_id}/diplomas", h.DiplomasGet).Methods(http.MethodGet)
	router.HandleFunc("/students/{student_id}/diplomas", h.DiplomaUpload).Methods(http.MethodPost)
	router.HandleFunc("/diplomas", h.DiplomasBulkUpload).Methods(http.MethodPost)

	slog.Info("listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
