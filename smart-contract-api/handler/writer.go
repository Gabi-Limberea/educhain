package handler

import (
	"encoding/json"
	"net/http"
)

type writer struct {
	http.ResponseWriter
	*json.Encoder
}

func newWriter(w http.ResponseWriter) *writer {
	return &writer{ResponseWriter: w, Encoder: json.NewEncoder(w)}
}

// Object encodes the object as JSON and sends it as a response
func (w *writer) Object(httpCode int, obj interface{}) {
	w.WriteHeader(httpCode)
	w.Encode(obj)
}

// Error wraps an error into a JSON object and sends it as a response
func (w *writer) Error(httpCode int, err error) {
	jsonErr := map[string]string{
		"error": err.Error(),
	}

	w.Object(httpCode, jsonErr)
}

// Message wraps a message into a JSON object and sends it as a response
func (w *writer) Message(httpCode int, msg string) {
	jsonMsg := map[string]string{
		"message": msg,
	}

	w.Object(httpCode, jsonMsg)
}
