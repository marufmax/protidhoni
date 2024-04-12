package mqttt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

type Message struct {
	topic   string
	content string
	qos     uint8
}

func Publish(client mqtt.Client, message Message) {
	if token := client.Publish(message.topic, message.qos, false, message.content); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Printf("Published message: %s\n", message.content)
}
