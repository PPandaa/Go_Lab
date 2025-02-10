package mongodb

import (
	"GoLab/guard"
	"context"
	"log"

	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB            *mongo.Database
	MongodbInfo   InfoStruct
	MongodbClient *mongo.Client
	valueFrom     string
)

type InfoStruct struct {
	URL          string
	Database     string
	Username     string
	Password     string
	AuthDatabase string
}

func Set() {
	logString := "  MongoDB Info." + "\n"

	valueFrom = "Hardcode"

	MongodbInfo = InfoStruct{
		URL:          "localhost:27017",
		AuthDatabase: "admin",
		Username:     "root",
		Password:     "rootPassword",
	}

	logString += "    FROM: " + valueFrom + "\n" +
		"      URL: " + MongodbInfo.URL + "\n" +
		"      Database: " + MongodbInfo.AuthDatabase + "\n" +
		"      Username: " + MongodbInfo.Username + "\n" +
		"      Password: " + MongodbInfo.Password + "\n"

	fmt.Print(logString + "\n")
}

func Connect() {
	// 設置 MongoDB 連線
	mongodb_client_options := options.Client().ApplyURI("mongodb://" + MongodbInfo.Username + ":" + MongodbInfo.Password + "@" + MongodbInfo.URL)

	// 建立連線
	new_client, err := mongo.Connect(context.TODO(), mongodb_client_options)
	if err != nil {
		guard.Logger.Error("mongodb connect fail - " + err.Error())
		for err != nil {
			guard.Logger.Info("mongodb retry connect")
			new_client, err = mongo.Connect(context.TODO(), mongodb_client_options)
			time.Sleep(5 * time.Second)
		}
	}

	MongodbClient = new_client

	// 測試連線
	err = MongodbClient.Ping(context.TODO(), nil)
	if err != nil {
		guard.Logger.Error("mongodb connect check fail")
	}

	guard.Logger.Info("mongodb connect check success")

	DB = MongodbClient.Database("hth")

	collection := DB.Collection("test")

	// 插入測試數據
	user := bson.D{{Key: "name", Value: "Alice"}, {Key: "age", Value: 25}, {Key: "email", Value: "alice@example.com"}}
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("插入成功，ID:", insertResult.InsertedID)
}
