package handlers

import (
	"net/http"

	"github.com/yahn1ukov/scribble/apps/gateway/internal/http/middlewares"
	"github.com/yahn1ukov/scribble/libs/grpc"
	"github.com/yahn1ukov/scribble/libs/respond"
	filepb "github.com/yahn1ukov/scribble/proto/file"
	"google.golang.org/grpc/codes"
)

type Handler struct {
	middleware *middlewares.Middleware
	fileClient filepb.FileServiceClient
}

func New(middleware *middlewares.Middleware, fileClient filepb.FileServiceClient) *Handler {
	return &Handler{
		middleware: middleware,
		fileClient: fileClient,
	}
}

func (h *Handler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if userID := h.middleware.GetUserIDFromCtx(ctx); userID == "" {
		respond.Error(w, http.StatusUnauthorized, middlewares.ErrUnauthorized)
		return
	}

	id := r.PathValue("fileId")
	noteID := r.PathValue("noteId")

	file, grpcErr := h.fileClient.DownloadFile(
		ctx,
		&filepb.DownloadFileRequest{
			Id:     id,
			NoteId: noteID,
		},
	)
	if grpcErr != nil {
		err := grpc.ParseError(grpcErr)

		if err.Code() == codes.NotFound {
			respond.Error(w, http.StatusNotFound, err.Error())
			return
		}

		respond.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	respond.File(w, file.Name, file.ContentType, file.Content)
}
