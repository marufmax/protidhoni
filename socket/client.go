package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	connection *websocket.Conn
	manager    *Manager
}

type ClientList map[*Client]bool

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
	}
}

func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	for {
		messageType, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message %v\n", err)
			}
			break
		}

		log.Println(messageType)
		log.Println(string(payload))
	}
}

//func cleanupConnection(c *Client) {
//	c.manager.removeClient(c)
//}
