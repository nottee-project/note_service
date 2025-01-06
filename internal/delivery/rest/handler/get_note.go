package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (t *NoteHandler) GetNote(c echo.Context) error {
	userId, ok := c.Get("user_id").(string)
	if !ok || userId == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
	}

	noteId := c.Param("id")
	if noteId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Note ID is required",
		})
	}

	note, err := t.NoteSrv.GetNote(context.Background(), noteId, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get note",
		})
	}

	return c.JSON(http.StatusCreated, note)
}
