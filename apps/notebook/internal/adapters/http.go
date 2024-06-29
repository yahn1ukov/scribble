package adapters

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/core/domain"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/core/dto"
	"github.com/yahn1ukov/scribble/apps/notebook/internal/core/ports"
	"github.com/yahn1ukov/scribble/libs/respond"
)

type HTTPHandler struct {
	service ports.Service
}

func NewHTTPHandler(service ports.Service) *HTTPHandler {
	return &HTTPHandler{
		service: service,
	}
}

func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var in dto.CreateInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.service.Create(ctx, &in); err != nil {
		if errors.Is(err, domain.ErrAlreadyExists) {
			respond.Error(w, http.StatusConflict, err.Error())
			return
		}
		respond.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	notebooks, err := h.service.GetAll(ctx)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	respond.JSON(w, http.StatusOK, notebooks)
}

func (h *HTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("notebookId")

	uuid, err := uuid.Parse(id)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	var in dto.UpdateInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.service.Update(ctx, uuid, &in); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			respond.Error(w, http.StatusNotFound, err.Error())
			return
		}
		respond.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("notebookId")

	uuid, err := uuid.Parse(id)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.service.Delete(ctx, uuid); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			respond.Error(w, http.StatusNotFound, err.Error())
			return
		}
		respond.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
