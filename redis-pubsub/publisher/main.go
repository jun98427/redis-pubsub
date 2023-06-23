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
	for {
		var msg string
		fmt.Print("Enter message: ")
		_, err := fmt.Scanln(&msg)
		if err != nil {
			return
		}

		err = client.Publish(context.Background(), topic, msg).Err()
		if err != nil {
			fmt.Printf("error: %s", err.Error())
		}

		if msg == "exit" {
			return
		}
	}

}
