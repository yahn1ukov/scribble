package mappers

import (
	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/gqlmodels"
	filepb "github.com/yahn1ukov/scribble/proto/file"
)

func (m *Mapper) GRPCFileToFile(file *filepb.File) gqlmodels.File {
	return gqlmodels.File{
		ID:          uuid.MustParse(file.Id),
		Name:        file.Name,
		Size:        file.Size,
		ContentType: file.ContentType,
		CreatedAt:   file.CreatedAt.AsTime(),
	}
}
