package note

import (
	"context"

	models "github.com/nottee-project/note_service/internal/models/note"
)

type NoteRepository interface {
	CreateNote(ctx context.Context, note models.CreateNoteParams) (models.Note, error)
	UpdateNote(ctx context.Context, note models.Note) (models.UpdateNoteResponse, error)
	GetNote(ctx context.Context, note_id, user_id string) (models.GetNoteResponse, error)
	ListNotes(ctx context.Context, listNotesParams models.ListNotesParams) ([]models.GetNoteResponse, error)
	DeleteNote(ctx context.Context, note_id, user_id string) error
}
