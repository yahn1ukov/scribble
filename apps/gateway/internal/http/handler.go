package http

import (
	"github.com/yahn1ukov/scribble/libs/grpc"
	"github.com/yahn1ukov/scribble/libs/respond"
	filepb "github.com/yahn1ukov/scribble/proto/file"
	"google.golang.org/grpc/codes"
	"net/http"
)

type Handler struct {
	fileClient filepb.FileServiceClient
}

func NewHandler(fileClient filepb.FileServiceClient) *Handler {
	return &Handler{
		fileClient: fileClient,
	}
}

func (h *Handler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	noteID := r.PathValue("noteId")
	id := r.PathValue("fileId")

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
