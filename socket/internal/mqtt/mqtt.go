package mqtt

import (
	"fmt"
	paho "github.com/eclipse/paho.mqtt.golang"
	"log/slog"
)

type MQTT struct {
	client paho.Client
	logger *slog.Logger
}

type Message struct {
	Topic   Topic
	Content string
	QoS     byte
}

type Topic struct {
	Name string
}

func NewMQTT() *MQTT {
	opts := paho.NewClientOptions()
	// TODO:: Move to config
	opts.AddBroker("tcp://mqtt:1883")
	opts.SetUsername("admin")
	opts.SetPassword("pass")
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectionLostHandler

	client := paho.NewClient(opts)

	return &MQTT{
		client: client,
	}
}

func (m *MQTT) StartServer() {
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

var connectHandler paho.OnConnectHandler = func(client paho.Client) {
	fmt.Println("Connected")
}

var connectionLostHandler paho.ConnectionLostHandler = func(client paho.Client, err error) {
	fmt.Printf("Connection Lost: %s\n", err.Error())
}
