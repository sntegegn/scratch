// redis set/get/pubsub- https://redis.uptrace.dev/guide/go-redis-pubsub.html
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("In main")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		log.Fatal("couldn't set value")
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		log.Fatal("couldn't get value from redis")
	}

	fmt.Println(val)

	err = client.Publish(ctx, "mychannel", "how are you?").Err()
	if err != nil {
		log.Fatal("couldn't publish to a redis channel")
	}

}
