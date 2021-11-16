package redisdb

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func Subscriber() {

	channels := []string{"TopicA"}
	sub := DB.Subscribe(CTX, channels...)
	for {
		msg, err := sub.ReceiveMessage(CTX)
		if err != nil {
			panic(err)
		}
		subHandler(msg)
	}

}

func subHandler(subMsg *redis.Message) {

	fmt.Println(subMsg)

}
