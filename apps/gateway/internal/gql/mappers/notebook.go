package mappers

import (
	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/gqlmodels"
	notebookpb "github.com/yahn1ukov/scribble/proto/notebook"
)

func (m *Mapper) GRPCNotebookToNotebook(notebook *notebookpb.Notebook) gqlmodels.Notebook {
	return gqlmodels.Notebook{
		ID:          uuid.MustParse(notebook.Id),
		Title:       notebook.Title,
		Description: notebook.Description,
		CreatedAt:   notebook.CreatedAt.AsTime(),
	}
}
