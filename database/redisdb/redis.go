package redisdb

import (
	"GoLab/guard"
	"GoLab/server"
	"GoLab/tool"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/go-redis/redis/v8"
)

var (
	DB          *redis.Client
	CTX         = context.Background()
	RedisdbInfo infoStruct
)

type infoStruct struct {
	URL      string
	Password string
}

func Set() {

	logString := "RedisDB Info." + "\n"

	if server.Location == server.Cloud {
		ensaasService := os.Getenv("ENSAAS_SERVICES")
		if !tool.IsEmptyString(ensaasService) {
			tempReader := strings.NewReader(ensaasService)
			m, _ := simplejson.NewFromReader(tempReader)
			redisdb := m.Get("redis").GetIndex(0).Get("credentials").MustMap()
			RedisdbInfo.URL = redisdb["host"].(string) + ":" + strconv.Itoa(redisdb["port"].(int))
			RedisdbInfo.Password = redisdb["password"].(string)
		} else {
			RedisdbInfo.URL = os.Getenv("REDIS_URL")
			RedisdbInfo.Password = os.Getenv("REDIS_PASSWORD")
		}
	} else {
		RedisdbInfo.URL = os.Getenv("REDIS_URL")
		redisdbPasswordFile := os.Getenv("REDIS_PASSWORD_FILE")
		if !tool.IsEmptyString(redisdbPasswordFile) {
			redisPassword, err := ioutil.ReadFile(redisdbPasswordFile)
			if err != nil {
				guard.Logger.Sugar().Fatalw("RedisDB Password File", "FilePath", redisdbPasswordFile)
			} else {
				RedisdbInfo.Password = string(redisPassword)
			}
		} else {
			RedisdbInfo.Password = os.Getenv("REDIS_PASSWORD")
		}
	}

	logString += "  URL: " + RedisdbInfo.URL + "\n" +
		"  PASSWORD: " + RedisdbInfo.Password + "\n"

	fmt.Print(logString + "\n")

}

func Connect() {

	DB = redis.NewClient(&redis.Options{
		Addr:     RedisdbInfo.URL,
		Password: RedisdbInfo.Password,
	})
	_, err := DB.Ping(CTX).Result()
	if err != nil {
		guard.Logger.Fatal("RedisDB Login Fail -> " + err.Error())
	}
	guard.Logger.Info("RedisDB Connect Success")

}
