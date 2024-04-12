package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"net/http"
)

type API struct {
	app    *echo.Echo
	logger *slog.Logger
}

func NewApi() *API {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"connected": true,
		})
	})

	a := &API{
		app: e,
	}

	return a
}

func (a *API) StartServer(port int) {
	a.app.HideBanner = true
	a.app.Use(middleware.Logger())
	a.app.Use(middleware.Recover())

	a.app.Logger.Fatal(a.app.Start(fmt.Sprintf(":%d", port)))
}
