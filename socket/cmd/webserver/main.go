package main

import (
	"fmt"
	paho "github.com/eclipse/paho.mqtt.golang"
	"marufalom.com/websocket/internal/api"
)

type webapp struct {
	app *api.API
}

func main() {
	setupMQTT()
	setupApi()
}

func setupApi() {
	app := api.NewApi()
	app.StartServer(9055)
}

func setupMQTT() {
	opts := paho.NewClientOptions()
	opts.AddBroker("tcp://mqtt:1883")
	opts.SetUsername("admin")
	opts.SetPassword("pass")
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectionLostHandler

	client := paho.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

var connectHandler paho.OnConnectHandler = func(client paho.Client) {
	fmt.Println("Connected")
}

var connectionLostHandler paho.ConnectionLostHandler = func(client paho.Client, err error) {
	fmt.Printf("Connection Lost: %s\n", err.Error())
}
