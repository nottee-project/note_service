package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	models "github.com/nottee-project/note_service/internal/models/note"
)

func (t *NoteHandler) UpdateNote(c echo.Context) error {
	// userID, ok := c.Get("user_id").(string)
	// if !ok || userID == "" {
	// 	return c.JSON(http.StatusUnauthorized, map[string]string{
	// 		"error": "Unauthorized",
	// 	})
	// }

	noteID := c.Param("id")
	if noteID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Note ID is required",
		})
	}

	var params models.Note
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	params.Id = noteID

	note, err := t.NoteSrv.UpdateNote(context.Background(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update note",
		})
	}

	return c.JSON(http.StatusOK, note)
}
