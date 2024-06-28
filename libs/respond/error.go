package respond

import (
	"encoding/json"
	"net/http"
)

type errorOutput struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorOutput{
		StatusCode: statusCode,
		Message:    message,
	})
}
