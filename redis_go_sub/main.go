package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	client.Subscribe(ctx, "mychannel")
	pubsub := client.Subscribe(ctx, "mychannel")

	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println("The messgae is: ", msg)
	}
}
