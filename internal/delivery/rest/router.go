package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/nottee-project/note_service/internal/bootstrap"
	"github.com/nottee-project/note_service/internal/delivery/rest/handler"
	mw_auth "github.com/nottee-project/note_service/internal/delivery/rest/middleware"
)

const prefix = "api/v1"

func RegisterRoutes(e *echo.Echo) error {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	noteSrv, err := bootstrap.CreateNoteService()
	if err != nil {
		return err
	}

	n := e.Group(prefix+"/note")
	n.Use(mw_auth.AuthMiddleware(authServiceURL))

	noteHandler := handler.NewNoteHandler(noteSrv)

	n.POST("/create", noteHandler.CreateNote)
	n.PUT("/update/:id", noteHandler.UpdateNote)
	n.POST("/list", noteHandler.ListNotes)
	n.GET("/get/:id", noteHandler.GetNote)
	n.DELETE("/delete/:id", noteHandler.DeleteNote)

	// e.POST("/webhook", handler.TelegramWebhookHandler)

	e.GET("/test", noteHandler.TestNote)

	return nil
}
