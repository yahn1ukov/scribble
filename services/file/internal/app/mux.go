package app

import (
	"net/http"

	"github.com/yahn1ukov/scribble/services/file/internal/adapters"
)

func NewMux(handler *adapters.HTTPHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /files/{fileId}", handler.Get)
	mux.HandleFunc("DELETE /files/{fileId}", handler.Delete)

	return mux
}
