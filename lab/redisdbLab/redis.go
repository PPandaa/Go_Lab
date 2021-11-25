package redisdbLab

import (
	"GoLab/database/redisdb"
	"GoLab/guard"
	"context"
	"fmt"
)

func GetChannels(patternA ...string) {

	ctx := context.Background()

	pattern := "*"
	if len(patternA) != 0 {
		pattern = patternA[0]
	}

	channels, err := redisdb.Client.PubSubChannels(ctx, pattern).Result()
	if err != nil {
		guard.Logger.Panic(err.Error())
	} else {
		for i, v := range channels {
			fmt.Println(i, v)
		}
	}

}
