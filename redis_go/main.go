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

	//Set and Get
	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		log.Fatal("couldn't set value")
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		log.Fatal("couldn't get value from redis")
	}

	fmt.Println(val)

	//Increment

	err = client.IncrBy(ctx, "counter", 5).Err()
	if err != nil {
		log.Fatal("couldn't increment a value")
	}

	ctr, _ := client.Get(ctx, "counter").Result()
	fmt.Println("counter is: ", ctr)

	//Pipeline

	pipeline := client.Pipeline()

	pipeline.Incr(ctx, "pipeline_val")
	pipeline.Incr(ctx, "pipeline_val")
	pipeline.Incr(ctx, "pipeline_val")

	pipeline.Exec(ctx)

	pipeline_val, _ := client.Get(ctx, "pipeline_val").Result()
	fmt.Println("pipeline val is: ", pipeline_val)

	//Publish

	err = client.Publish(ctx, "mychannel", "how are you?").Err()
	if err != nil {
		log.Fatal("couldn't publish to a redis channel")
	}

}
