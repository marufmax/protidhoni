package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"marufalom.com/websocket/internal/api/handler"
	"net/http"
)

type API struct {
	app    *echo.Echo
	logger *slog.Logger
}

func NewApi() *API {
	e := echo.New()
	a := &API{
		app: e,
	}

	return a
}

func (a *API) setupHandlers() {
	a.app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"connected": true,
			"health":    "ok",
		})
	})

	messageRoutes := a.app.Group("/v1/messages")
	messageRoutes.POST("/single/:receiverId", handler.SendSingleHandler)
}

func (a *API) StartServer(port int) {
	a.app.HideBanner = true
	a.app.Use(middleware.Logger())
	a.app.Use(middleware.Recover())
	a.setupHandlers()

	a.app.Logger.Fatal(a.app.Start(fmt.Sprintf(":%d", port)))
}
