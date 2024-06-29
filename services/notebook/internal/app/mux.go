package app

import (
	"net/http"

	"github.com/yahn1ukov/scribble/services/notebook/internal/adapters"
)

func NewMux(handler *adapters.HTTPHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /notebooks", handler.Create)
	mux.HandleFunc("GET /notebooks", handler.GetAll)
	mux.HandleFunc("PATCH /notebooks/{notebookId}", handler.Update)
	mux.HandleFunc("DELETE /notebooks/{notebookId}", handler.Delete)

	return mux
}
