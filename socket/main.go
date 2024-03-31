package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	setupApi()
}

func setupApi() {
	manager := NewManager()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, from socket world!")
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/chat/ws", manager.serveWs)

	e.Logger.Fatal(e.Start(":9055"))
}
