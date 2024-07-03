package adapters

import (
	"errors"
	"net/http"

	"github.com/yahn1ukov/scribble/libs/respond"
	"github.com/yahn1ukov/scribble/services/note/internal/core/domain"
	"github.com/yahn1ukov/scribble/services/note/internal/core/dto"
	"github.com/yahn1ukov/scribble/services/note/internal/core/ports"
	"google.golang.org/grpc/codes"
)

type HTTPHandler struct {
	service        ports.Service
	notebookClient *NotebookGRPCClient
	fileClient     *FileGRPCClient
}

func NewHTTPHandler(service ports.Service, notebookClient *NotebookGRPCClient, fileClient *FileGRPCClient) *HTTPHandler {
	return &HTTPHandler{
		service:        service,
		notebookClient: notebookClient,
		fileClient:     fileClient,
	}
}

func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	notebookId := r.PathValue("notebookId")

	if err := r.ParseMultipartForm(1 << 20); err != nil {
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

	id, err := h.service.Create(ctx, notebookId, &in)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if grpcErr := h.fileClient.UploadAll(ctx, id, in.Files); grpcErr != nil {
		respond.Error(w, http.StatusBadRequest, grpcErr.Message())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	notebookId := r.PathValue("notebookId")

	if grpcErr := h.notebookClient.Exists(ctx, notebookId); grpcErr != nil {
		if grpcErr.Code() == codes.NotFound {
			respond.Error(w, http.StatusNotFound, grpcErr.Message())
			return
		}
		respond.Error(w, http.StatusBadRequest, grpcErr.Message())
		return
	}

	notes, err := h.service.GetAll(ctx, notebookId)
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

	if grpcErr := h.notebookClient.Exists(ctx, notebookId); grpcErr != nil {
		if grpcErr.Code() == codes.NotFound {
			respond.Error(w, http.StatusNotFound, grpcErr.Message())
			return
		}
		respond.Error(w, http.StatusBadRequest, grpcErr.Message())
		return
	}

	note, err := h.service.Get(ctx, id, notebookId)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			respond.Error(w, http.StatusNotFound, err.Error())
			return
		}
		respond.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	files, grpcErr := h.fileClient.GetAll(ctx, id)
	if grpcErr != nil {
		respond.Error(w, http.StatusBadRequest, grpcErr.Message())
		return
	}

	note.Files = files

	respond.JSON(w, http.StatusOK, note)
}

func (h *HTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	notebookId := r.PathValue("notebookId")
	id := r.PathValue("noteId")

	if err := r.ParseMultipartForm(1 << 20); err != nil {
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

	if err := h.service.Update(ctx, id, notebookId, &in); err != nil {
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

	if grpcErr := h.notebookClient.Exists(ctx, notebookId); grpcErr != nil {
		if grpcErr.Code() == codes.NotFound {
			respond.Error(w, http.StatusNotFound, grpcErr.Message())
			return
		}
		respond.Error(w, http.StatusBadRequest, grpcErr.Message())
		return
	}

	if err := h.service.Delete(ctx, id, notebookId); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			respond.Error(w, http.StatusNotFound, err.Error())
			return
		}
		respond.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HTTPHandler) Upload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	notebookId := r.PathValue("notebookId")
	id := r.PathValue("noteId")

	if err := r.ParseMultipartForm(1 << 20); err != nil {
		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, file, _ := r.FormFile("file")

	if grpcErr := h.notebookClient.Exists(ctx, notebookId); grpcErr != nil {
		if grpcErr.Code() == codes.NotFound {
			respond.Error(w, http.StatusNotFound, grpcErr.Message())
			return
		}
		respond.Error(w, http.StatusBadRequest, grpcErr.Message())
		return
	}

	if grpcErr := h.fileClient.Upload(ctx, id, file); grpcErr != nil {
		respond.Error(w, http.StatusBadRequest, grpcErr.Message())
		return
	}

	w.WriteHeader(http.StatusOK)
}
