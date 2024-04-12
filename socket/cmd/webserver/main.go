package main

import (
	"marufalom.com/websocket/internal/api"
	"marufalom.com/websocket/internal/mqtt"
)

type webapp struct {
	Api  *api.API
	Mqtt *mqtt.MQTT
}

var App webapp

func main() {
	App = webapp{
		Api:  setupApi(),
		Mqtt: setupMQTT(),
	}
	App.Api.StartServer(9055)
	App.Mqtt.StartServer()
}

func setupApi() *api.API {
	app := api.NewApi()

	return app
}

func setupMQTT() *mqtt.MQTT {
	return mqtt.NewMQTT()
}
