package redisdbLab

import "GoLab/database/redisdb"

func Pub() {

	redisdb.Client.Publish(redisdb.CTX, "III", "Peter")

}
