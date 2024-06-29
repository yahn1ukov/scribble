package app

import (
	"net/http"

	"github.com/yahn1ukov/scribble/services/note/internal/adapters"
)

func NewMux(handler *adapters.HTTPHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /notebooks/{notebookId}/notes", handler.Create)
	mux.HandleFunc("GET /notebooks/{notebookId}/notes", handler.GetAll)
	mux.HandleFunc("GET /notebooks/{notebookId}/notes/{noteId}", handler.Get)
	mux.HandleFunc("PATCH /notebooks/{notebookId}/notes/{noteId}", handler.Update)
	mux.HandleFunc("DELETE /notebooks/{notebookId}/notes/{noteId}", handler.Delete)

	return mux
}
