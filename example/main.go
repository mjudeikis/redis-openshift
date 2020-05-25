package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Go Redis Example")

	server := os.Getenv("REDIS_SERVER")
	passwd := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     server + ":6379",
		Password: passwd,
		DB:       0,
	})

	for {
		pong, err := client.Ping().Result()
		fmt.Println(pong, err)

		err = client.Set("name", "MJ", 0).Err()
		if err != nil {
			fmt.Println(err)
		}

		val, err := client.Get("name").Result()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(val)

		err = client.Del("name").Err()
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Second * 10)
	}

}
