package redisdb

import (
	"fmt"

	"GoLab/guard"

	"github.com/go-redis/redis/v8"
)

func Subscriber() {

	channels := []string{"*"}
	sub := Client.PSubscribe(CTX, channels...)
	for {
		msg, err := sub.ReceiveMessage(CTX)
		if err != nil {
			guard.Logger.Panic(err.Error())
		}
		subHandler(msg)
	}

}

func subHandler(msg *redis.Message) {

	fmt.Println(msg)

}
