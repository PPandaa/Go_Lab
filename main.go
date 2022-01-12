package main

import (
	"GoLab/dependency"
	"GoLab/guard"
	"GoLab/lab/mongodbLab"
	"GoLab/server"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load("local.env")
	if err != nil {
		guard.Logger.Fatal("loading env file: " + err.Error())
	}

	server.Up()

	dependency.Set()

}

// var wg sync.WaitGroup

func main() {

	// wg.Add(1)

	guard.Logger.Info(server.AppNameL + "-" + server.ServiceName + " active")
	mongodbLab.RemoveAllCollection()

	// wg.Wait()

}
