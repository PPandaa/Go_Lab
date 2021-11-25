package redisdbLab

import (
	"GoLab/database/redisdb"
	"context"
)

func Pub() {

	ctx := context.Background()
	redisdb.Client.Publish(ctx, "III", "Peter")

}
