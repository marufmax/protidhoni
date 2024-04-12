package mqtt

import (
	"fmt"
	"os"
)

func (m *MQTT) Publish(message Message) {
	if token := m.client.Publish(message.Topic.Name, message.QoS, false, message.Content); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Printf("Published message: %s\n", message.Content)
}
