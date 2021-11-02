package muxLab

import (
	"github.com/gorilla/mux"
)

func router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", hello).Methods("GET", "OPTIONS")
	return router

}
