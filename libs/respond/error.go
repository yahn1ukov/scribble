package respond

import (
	"encoding/json"
	"net/http"
)

type error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(error{
		Code:    code,
		Message: message,
	})
}
