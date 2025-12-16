package main

import (
	"github.com/gorilla/mux"
)

func main() {
	var acc Account
	router := mux.NewRouter()
	router.HandleFunc("/auth", CreateAccount).Methods("POST")
}
