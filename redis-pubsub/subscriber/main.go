package main

import (
	"context"
	"fmt"
	"github.com/jun98427/redis-pubsub/pkg/pubsub"
)

func main() {
	var topic string
	fmt.Print("Enter topic: ")
	_, err := fmt.Scanln(&topic)
	if err != nil {
		return
	}

	client := pubsub.RedisClient()
	sub := client.Subscribe(context.Background(), topic)
	defer func() {
		_ = sub.Close()
	}()

	for {
		select {
		case msg := <-sub.Channel():
			fmt.Printf("enter : %s\n", msg.Payload)
			if msg.Payload == "exit" {
				return
			}
		}
	}
}
