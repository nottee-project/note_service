package tm

import "github.com/pkg/errors"

const (
	tableNameNote = "note"

	fieldNameNoteId    = "id"
	fieldNameTitle     = "title"
	fieldNameBody      = "body"
	fieldNameCreatedAt = "created_at"
	fieldNameUpdatedAt = "updated_at"
)

var NoteNotFound = errors.New("this note not found")
