package muxLab

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	logString := r.Method + " " + r.URL.Path + "\n"
	w.WriteHeader(200)
	msg := "This is Mux Lab"
	w.Write([]byte(msg))
	logString += "- Response" + "\n" + "200" + "\n"
	log.Print(logString + "\n")

}
