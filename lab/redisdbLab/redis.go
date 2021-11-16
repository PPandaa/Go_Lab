package redisdbLab

import (
	"GoLab/database/redisdb"
	"fmt"
)

func GetChannels(patternA ...string) {

	pattern := "*"
	if len(patternA) != 0 {
		pattern = patternA[0]
	}

	channels := redisdb.DB.PubSubChannels(redisdb.CTX, pattern)
	for i, v := range channels.Val() {
		fmt.Println(i, v)
	}

}
