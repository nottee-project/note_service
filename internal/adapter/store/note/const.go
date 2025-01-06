package tm

import "github.com/pkg/errors"

const (
	tableNameNotes = "notes"

	fieldNameNoteId    = "id"
	fieldNameUserId    = "user_id"
	fieldNameTitle     = "title"
	fieldNameBody      = "body"
	fieldNameCreatedAt = "created_at"
	fieldNameUpdatedAt = "updated_at"
)

var NoteNotFound = errors.New("this note not found")
