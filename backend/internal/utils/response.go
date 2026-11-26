package utils

import (
	"encoding/json"
	"net/http"
)

// RespondJSON writes any payload as JSON with the given HTTP status code.
func RespondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// RespondOK writes a standard 200 OK JSON envelope.
func RespondOK(w http.ResponseWriter, payload any) {
	RespondJSON(w, http.StatusOK, payload)
}

// RespondCreated writes a standard 201 Created JSON envelope.
func RespondCreated(w http.ResponseWriter, payload any) {
	RespondJSON(w, http.StatusCreated, payload)
}

// RespondError writes a standardized error JSON payload.
func RespondError(w http.ResponseWriter, status int, code, message string) {
	RespondJSON(w, status, map[string]any{
		"success": false,
		"code":    code,
		"error":   message,
	})
}
