package handler

import (
	"net/http"

	tgbotapi "github.com/matterbridge/telegram-bot-api/v6"
	"github.com/labstack/echo/v4"
)

func TelegramWebhookHandler(c echo.Context) error {
	var update tgbotapi.Update
	if err := c.Bind(&update); err != nil {
		c.Logger().Errorf("Bind error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if update.Message != nil {
		c.Logger().Infof("Received message: %s", update.Message.Text)

		bot, err := tgbotapi.NewBotAPI("7720869175:AAEv76LlBHuRTQ5D_bVgPuIwVR0mv5iMrzQ")
		if err != nil {
			c.Logger().Errorf("Bot init error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to initialize bot"})
		}

		webAppButton := tgbotapi.NewInlineKeyboardButtonWebApp("Открыть мини-приложение", tgbotapi.WebAppInfo{
			URL: "https://29eb-46-226-163-65.ngrok-free.app",
		})
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать в приложение для заметок!")
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(webAppButton),
		)

		if _, err := bot.Send(msg); err != nil {
			c.Logger().Errorf("Send message error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to send message"})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
