package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"marufalom.com/websocket/internal/mqtt"
	"net/http"
	"time"
)

const topicSingleChat = "messages/%s/chat"

func SendSingleHandler(c echo.Context) error {
	receiver := c.Param("receiverId")

	var message mqtt.Message
	if err := c.Bind(&message); err != nil {
		return err
	}

	message.Topic = mqtt.Topic{Name: fmt.Sprintf(topicSingleChat, receiver)}
	message.QoS = 2

	return c.JSON(http.StatusOK, map[string]interface{}{
		"sent":     true,
		"receiver": receiver,
		"sent_at":  time.Now().Unix(),
	})
}
