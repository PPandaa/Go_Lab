package pkg

// import (
// 	"GoLab/database/mongodb"
// 	"GoLab/database/redisdb"
// 	"GoLab/guard"
// 	"GoLab/tool"

// 	"context"
// 	"strings"

// 	"github.com/go-redis/redis/v8"
// 	"gopkg.in/mgo.v2/bson"
// )

// type SubscribeMessageStruct struct {
// 	ChannelFrom string
// 	ChannelType string
// 	ChannelID   string
// 	Payload     map[string]interface{}
// }

// var (
// 	ctx = context.Background()
// )

// func subChannelDivider(channel string) (string, string, string) {

// 	splitStrings := strings.Split(channel, ".")

// 	return splitStrings[0], splitStrings[1], splitStrings[2] + "." + splitStrings[3]

// }

// func Redis_Subscriber() {

// 	channels := []string{"ifp-organizer.*"}
// 	sub := redisdb.Client.PSubscribe(ctx, channels...)
// 	for {
// 		msg, err := sub.ReceiveMessage(ctx)
// 		if err != nil {
// 			guard.Logger.Panic(err.Error())
// 		}
// 		subHandler(msg)
// 	}

// }

// func subHandler(msg *redis.Message) {

// 	var subMsg SubscribeMessageStruct
// 	subMsg.ChannelFrom, subMsg.ChannelType, subMsg.ChannelID = subChannelDivider(msg.Channel)
// 	subMsg.Payload = tool.ConvertStringToMap(msg.Payload)
// 	switch subMsg.ChannelType {
// 	case "Group":
// 		guard.Logger.Sugar().Infow("Sub", "Channel", msg.Channel, "GroupID", subMsg.Payload["id"])
// 		groupHandler(subMsg.ChannelID, subMsg.Payload)
// 	case "Machine":
// 		guard.Logger.Sugar().Infow("Sub", "Channel", msg.Channel, "GroupID", subMsg.Payload["groupId"], "MachineID", subMsg.Payload["id"])
// 		machineHandler(subMsg.ChannelID, subMsg.Payload)
// 	}

// }

// func groupHandler(id string, payload map[string]interface{}) {

// 	collection := mongodb.DB.C(mongodb.Test)
// 	if len(payload) == 0 {
// 		Desk_DeleteGroup(id)
// 	} else {
// 		existGroup := map[string]interface{}{}
// 		collection.Pipe([]bson.M{{"$match": bson.M{"GroupID": payload["id"].(string)}}}).One(&existGroup)
// 		if len(existGroup) == 0 {
// 			Desk_GetGroup("redis")
// 		} else {
// 			group := map[string]interface{}{"GroupName": payload["name"].(string), "TimeZone": payload["timeZone"].(string)}
// 			Desk_SetGroup(id, group)
// 		}
// 	}

// }

// func machineHandler(id string, payload map[string]interface{}) {

// 	if len(payload) == 0 {
// 		Desk_DeleteMachine(id)
// 		Desk_DeleteStation(id)
// 	} else {
// 		Desk_GetMachine("redis", payload["groupId"].(string))
// 		Desk_GetStation("redis", payload["groupId"].(string))
// 	}

// }
