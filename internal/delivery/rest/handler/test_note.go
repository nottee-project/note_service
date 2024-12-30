package handler

import "github.com/labstack/echo/v4"

func (h *NoteHandler) TestNote(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"message": "Note tested successfully",
	})
}
