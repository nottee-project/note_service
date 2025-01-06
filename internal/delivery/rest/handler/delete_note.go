package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (t *NoteHandler) DeleteNote(c echo.Context) error {
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

	err := t.NoteSrv.DeleteNote(context.Background(), noteId, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete note",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Note deleted successfully",
	})
}
