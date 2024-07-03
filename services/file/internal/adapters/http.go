package adapters

import (
	"errors"
	"net/http"

	"github.com/yahn1ukov/scribble/libs/respond"
	"github.com/yahn1ukov/scribble/services/file/internal/core/domain"
	"github.com/yahn1ukov/scribble/services/file/internal/core/ports"
)

type HTTPHandler struct {
	service ports.Service
}

func NewHTTPHandler(service ports.Service) *HTTPHandler {
	return &HTTPHandler{
		service: service,
	}
}

func (h *HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("fileId")

	file, err := h.service.Get(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			respond.Error(w, http.StatusNotFound, err.Error())
			return
		}
		respond.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	respond.File(w, http.StatusOK, file.Name, file.ContentType, file.Content)
}

func (h *HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("fileId")

	if err := h.service.Delete(ctx, id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			respond.Error(w, http.StatusNotFound, err.Error())
			return
		}
		respond.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
