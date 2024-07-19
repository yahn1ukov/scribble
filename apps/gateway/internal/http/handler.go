package http

import (
	"net/http"

	"github.com/yahn1ukov/scribble/apps/gateway/internal/grpc/clients"
	"github.com/yahn1ukov/scribble/libs/respond"
	"google.golang.org/grpc/codes"
)

type Handler struct {
	grpc *clients.Client
}

func NewHandler(grpc *clients.Client) *Handler {
	return &Handler{
		grpc: grpc,
	}
}

func (h *Handler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("fileId")
	noteID := r.PathValue("noteId")

	file, grpcErr := h.grpc.DownloadFile(ctx, id, noteID)
	if grpcErr != nil {
		if grpcErr.Code() == codes.NotFound {
			respond.Error(w, http.StatusNotFound, grpcErr.Error())
			return
		}
		respond.Error(w, http.StatusBadRequest, grpcErr.Error())
		return
	}

	respond.File(w, file.Name, file.ContentType, file.Content)
}
