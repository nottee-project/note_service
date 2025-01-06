package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	models "github.com/nottee-project/note_service/internal/models/note"
)

func (t *NoteHandler) ListNotes(c echo.Context) error {
	var params models.ListNotesParams

	userId, ok := c.Get("user_id").(string)
	if !ok || userId == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
	}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	params.UserId = userId

	notes, err := t.NoteSrv.ListNotes(context.Background(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve notes",
		})
	}

	return c.JSON(http.StatusOK, notes)
}
