package handler

import "github.com/labstack/echo/v4"

type (
	MessageHandler interface {
		SendMessage(c echo.Context) error
	}

	messageHandler struct {
	}
)
