package muxLab

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func Server() {

	r := router()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	handler := c.Handler(r)
	log.Print("Server Start" + "\n")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
