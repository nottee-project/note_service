package handler

import (
	"github.com/nottee-project/note_service/internal/service/note"
)

type NoteHandler struct {
	NoteSrv *note.NoteService
}

func NewNoteHandler(NoteSrv *note.NoteService) NoteHandler {
	return NoteHandler{
		NoteSrv: NoteSrv,
	}
}
