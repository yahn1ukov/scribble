package app

import (
	"net/http"

	"github.com/yahn1ukov/scribble/apps/notebook/internal/adapters"
)

func NewMux(handler *adapters.HTTPHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /notebooks", handler.Create)
	mux.HandleFunc("GET /notebooks", handler.GetAll)
	mux.HandleFunc("PATCH /notebooks/{id}", handler.Update)
	mux.HandleFunc("DELETE /notebooks/{id}", handler.Delete)

	return mux
}
