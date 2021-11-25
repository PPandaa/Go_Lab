package redisdb

import (
	"fmt"

	"GoLab/guard"

	"github.com/go-redis/redis/v8"
)

func Subscriber() {

	channels := []string{"*"}
	sub := Client.PSubscribe(ctx, channels...)
	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			guard.Logger.Panic(err.Error())
		}
		subHandler(msg)
	}

}

func subHandler(msg *redis.Message) {

	fmt.Println(msg)

}
