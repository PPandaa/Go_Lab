package redisdbLab

import (
	"GoLab/database/redisdb"
	"GoLab/guard"
	"fmt"
)

func GetChannels(patternA ...string) {

	pattern := "*"
	if len(patternA) != 0 {
		pattern = patternA[0]
	}

	channels, err := redisdb.DB.PubSubChannels(redisdb.CTX, pattern).Result()
	if err != nil {
		guard.Logger.Panic(err.Error())
	} else {
		for i, v := range channels {
			fmt.Println(i, v)
		}
	}

}
