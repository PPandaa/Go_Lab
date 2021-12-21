package main

import (
	"GoLab/guard"

	"github.com/joho/godotenv"
)

// var wg sync.WaitGroup

func init() {

	err := godotenv.Load("local.env")
	if err != nil {
		guard.Logger.Fatal("Loading ENV File: " + err.Error())
	}

	// server.Set()
	// mongodb.Set()
	// redisdb.Set()
	// miniodb.Set()
	// dependency.Set()

	// mongodb.Connect()
	// redisdb.Connect()
	// miniodb.Connect()

	// socketLab.Set()

}

func main() {

	// wg.Add(1)

	guard.Logger.Info("GoLab Server Active")
	// mongodbLab.RemoveAllCollection()

	// wg.Wait()

}
