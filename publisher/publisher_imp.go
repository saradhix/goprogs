package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttFunc mqtt.MessageHandler = func(client *mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Topic: %s\n", msg.Topic())
	fmt.Printf("Msg: %s\n", msg.Payload())
}

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883/")
	opts.SetClientID("clientE")
	opts.SetDefaultPublishHandler(mqttFunc)

	// This is the line I change
	opts.SetCleanSession(false)

	c := mqtt.NewClient(opts)
	token := c.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token = c.Subscribe("together/example1", 1, nil)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		panic(1)
	}

	forever := make(chan bool)
	<-forever
}
