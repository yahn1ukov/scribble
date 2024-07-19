package respond

import (
	"encoding/json"
	"net/http"
)

type errorOutput struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorOutput{
		Code:    code,
		Message: err.Error(),
	})
}
