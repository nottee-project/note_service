package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	models "github.com/nottee-project/note_service/internal/models/note"
)

func (t *NoteHandler) CreateNote(c echo.Context) error {
	userId, ok := c.Get("user_id").(string)
	if !ok || userId == "" {
		log.Println("user_id is missing or invalid")
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
	}

	log.Printf("user_id: %s", userId)

	var params models.CreateNoteParams
	if err := c.Bind(&params); err != nil {
		log.Printf("Failed to bind request: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	log.Printf("CreateNoteParams: %+v", params)

	if params.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Title is required",
		})
	}

	params.UserId = userId
	log.Printf("Final CreateNoteParams: %+v", params)

	note, err := t.NoteSrv.CreateNote(context.Background(), params)
	if err != nil {
		log.Printf("Failed to create note: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create note",
		})
	}

	return c.JSON(http.StatusCreated, note)
}
