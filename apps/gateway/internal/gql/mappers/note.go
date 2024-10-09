package mappers

import (
	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/gql/gqlmodels"
	notepb "github.com/yahn1ukov/scribble/proto/note"
)

func (m *Mapper) GRPCNoteToNote(note *notepb.Note) gqlmodels.Note {
	return gqlmodels.Note{
		ID:        uuid.MustParse(note.Id),
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt.AsTime(),
	}
}
