package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/nottee-project/note_service/internal/bootstrap"
	"github.com/nottee-project/note_service/internal/delivery/rest/handler"
	mw_auth "github.com/nottee-project/note_service/internal/delivery/rest/middleware"
)

const prefix = "/api/v1"

func RegisterRoutes(e *echo.Echo, authServiceURL string) error {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	noteSrv, err := bootstrap.CreateNoteService()
	if err != nil {
		return err
	}

	n := e.Group(prefix+"/note")
	n.Use(mw_auth.AuthMiddleware(authServiceURL))

	noteHandler := handler.NewNoteHandler(noteSrv)

	n.POST("", noteHandler.CreateNote)
	n.PUT("/:id", noteHandler.UpdateNote)
	n.GET("", noteHandler.ListNotes)
	n.GET("/:id", noteHandler.GetNote)
	n.DELETE("/:id", noteHandler.DeleteNote)

	// e.POST("/webhook", handler.TelegramWebhookHandler)

	e.GET("/test", noteHandler.TestNote)

	return nil
}
