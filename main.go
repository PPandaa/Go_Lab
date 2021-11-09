package main

import (
	"GoLab/guard"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load("local.env")
	if err != nil {
		guard.Logger.Fatal("Error Loading ENV File: " + err.Error())
	}

	// server.Set()
	// mongodb.Set()
	// dependency.Set()
	// mongodb.Connect()
	// socketLab.Set()

}

func main() {

	guard.Logger.Info("GoLab Server Active")

}
