package bootstrap

import (
	"github.com/nottee-project/note_service/internal/config"
	note_store "github.com/nottee-project/note_service/internal/adapter/store/note"
	"github.com/nottee-project/note_service/internal/adapter/store"
	"github.com/pkg/errors"
	"github.com/nottee-project/note_service/internal/service/note"
)

func CreateNoteService() (*note.NoteService, error) {
	cfg, err := config.NewConfig()
	if err != nil {
			return nil, errors.Wrap(err, "CreateConfig")
	}

	dbStore, err := store.New(cfg.Database)
	if err != nil {
			return nil, err
	}

	noteStore := &note_store.NoteStore{Store: dbStore}
	return note.New(noteStore)
}