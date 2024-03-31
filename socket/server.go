package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, from socket world!")
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/chat/ws", serveWs)

	e.Logger.Fatal(e.Start(":9055"))
}

func serveWs(c echo.Context) error {
	log.Println("Received new connection")
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		log.Fatal(err)

		return err
	}

	log.Printf("client connected")

	defer ws.Close()

	return nil
}
