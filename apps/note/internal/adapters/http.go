package adapters

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/apps/note/internal/core/domain"
	"github.com/yahn1ukov/scribble/apps/note/internal/core/dto"
	"github.com/yahn1ukov/scribble/apps/note/internal/core/ports"
	"github.com/yahn1ukov/scribble/libs/respond"
	"google.golang.org/grpc/codes"
)

type HTTPHandler struct {
	service        ports.Service
	notebookClient *NotebookGRPCClient
}

func NewHTTPHandler(service ports.Service, notebookClient *NotebookGRPCClient) *HTTPHandler {
	return &HTTPHandler{
		service:        service,
		notebookClient: notebookClient,
	}
}

func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	notebookId := r.PathValue("notebookId")

	notebookUuid, err := uuid.Parse(notebookId)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = r.ParseMultipartForm(4 << 20); err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	var in dto.CreateInput
	in.Title = r.FormValue("title")
	in.Body = r.FormValue("body")
	in.Files = r.MultipartForm.File["files"]

	if grpcErr := h.notebookClient.Exists(ctx, notebookId); grpcErr != nil {
		if grpcErr.Code() == codes.NotFound {
			respond.Error(w, http.StatusNotFound, grpcErr.Message())
			return
		}
		respond.Error(w, http.StatusBadRequest, grpcErr.Message())
		return
	}

	if err = h.service.Create(ctx, notebookUuid, &in); err != nil {
		respond.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	notebookId := r.PathValue("notebookId")

	notebookUuid, err := uuid.Parse(notebookId)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if grpcErr := h.notebookClient.Exists(ctx, notebookId); grpcErr != nil {
		if grpcErr.Code() == codes.NotFound {
			respond.Error(w, http.StatusNotFound, grpcErr.Message())
			return
		}
		respond.Error(w, http.StatusBadRequest, grpcErr.Message())
		return
	}

	notes, err := h.service.GetAll(ctx, notebookUuid)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	respond.JSON(w, http.StatusOK, notes)
}

func (h *HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	notebookId := r.PathValue("notebookId")
	id := r.PathValue("noteId")

	notebookUuid, err := uuid.Parse(notebookId)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if grpcErr := h.notebookClient.Exists(ctx, notebookId); grpcErr != nil {
		if grpcErr.Code() == codes.NotFound {
			respond.Error(w, http.StatusNotFound, grpcErr.Message())
			return
		}
		respond.Error(w, http.StatusBadRequest, grpcErr.Message())
		return
	}

	note, err := h.service.Get(ctx, uuid, notebookUuid)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			respond.Error(w, http.StatusNotFound, err.Error())
			return
		}
		respond.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	respond.JSON(w, http.StatusOK, note)
}

func (h *HTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	notebookId := r.PathValue("notebookId")
	id := r.PathValue("noteId")

	notebookUuid, err := uuid.Parse(notebookId)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = r.ParseMultipartForm(4 << 20); err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	var in dto.UpdateInput
	in.Title = r.FormValue("title")
	in.Body = r.FormValue("body")
	in.Files = r.MultipartForm.File["files"]

	if grpcErr := h.notebookClient.Exists(ctx, notebookId); grpcErr != nil {
		if grpcErr.Code() == codes.NotFound {
			respond.Error(w, http.StatusNotFound, grpcErr.Message())
			return
		}
		respond.Error(w, http.StatusBadRequest, grpcErr.Message())
		return
	}

	if err := h.service.Update(ctx, uuid, notebookUuid, &in); err != nil {
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
	notebookId := r.PathValue("notebookId")
	id := r.PathValue("noteId")

	notebookUuid, err := uuid.Parse(notebookId)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if grpcErr := h.notebookClient.Exists(ctx, notebookId); grpcErr != nil {
		if grpcErr.Code() == codes.NotFound {
			respond.Error(w, http.StatusNotFound, grpcErr.Message())
			return
		}
		respond.Error(w, http.StatusBadRequest, grpcErr.Message())
		return
	}

	if err := h.service.Delete(ctx, uuid, notebookUuid); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			respond.Error(w, http.StatusNotFound, err.Error())
			return
		}
		respond.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
