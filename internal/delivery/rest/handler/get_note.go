package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (t *NoteHandler) GetNote(c echo.Context) error {
	noteID := c.Param("id")
	if noteID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Note ID is required",
		})
	}

	note, err := t.NoteSrv.GetNote(context.Background(), noteID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get note",
		})
	}

	return c.JSON(http.StatusCreated, note)
}
