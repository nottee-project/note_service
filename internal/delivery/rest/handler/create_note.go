package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	models "github.com/nottee-project/note_service/internal/models/note"
)

func (t *NoteHandler) CreateNote(c echo.Context) error {
	var params models.CreateNoteParams

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if params.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Title is required",
		})
	}

	note, err := t.NoteSrv.CreateNote(context.Background(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create note",
		})
	}

	return c.JSON(http.StatusCreated, note)
}
