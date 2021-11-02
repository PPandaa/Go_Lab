package main

import (
	"GoLab/database/mongodb"
	"GoLab/dependency"
	"GoLab/mongodbLab"
	"GoLab/server"
	"log"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Error Loading ENV File: " + err.Error())
	}

	server.Set()
	mongodb.Set()
	dependency.Set()
	mongodb.Connect()

}

func main() {

	log.Print("Server Active \n")
	mongodbLab.PipeTest()

}
