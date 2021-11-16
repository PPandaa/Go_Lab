package redisdbLab

import "GoLab/database/redisdb"

func Pub() {

	redisdb.DB.Publish(redisdb.CTX, "III", "Peter")

}
