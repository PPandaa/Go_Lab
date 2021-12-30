package main

import (
	"GoLab/database/mongodb"
	"GoLab/guard"
	"GoLab/server"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load("local.env")
	if err != nil {
		guard.Logger.Fatal("Loading ENV File: " + err.Error())
	}

	server.Check()

	mongodb.Set()
	// redisdb.Set()
	// miniodb.Set()
	// dependency.Set()

	mongodb.Connect()
	// redisdb.Connect()
	// miniodb.Connect()

	// socketLab.Set()

}

// var wg sync.WaitGroup

func main() {

	// wg.Add(1)

	guard.Logger.Info("GoLab Server Active")
	// mongodbLab.RemoveAllCollection()

	// wg.Wait()

}
