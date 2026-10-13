package main

import (
	"fmt"
	"time"

	"github.com/Darpan313/pubsub/pubsub"
)

func main() {
	broker := pubsub.NewBroker()
	topic := "system-alert"

	// 1. Create Subscriber
	sub := broker.Subscribe(topic, 10)

	go func() {
		for msg := range sub.Messages() {
			fmt.Printf("Received Alert on '%s': %v\n", msg.Topic, msg.Payload)
		}
		fmt.Println("Subscriber shut down.")
	}()

	// 3. Publish messages
	broker.Publish(topic, "High CPU Usage")
	broker.Publish(topic, "Database Connection Lost")

	// Allow time for messages to be processed
	time.Sleep(100 * time.Millisecond)

	// 4. Cleanup
	broker.UnSubscribe(sub)
	time.Sleep(50 * time.Millisecond) // Wait for shutdown print
}