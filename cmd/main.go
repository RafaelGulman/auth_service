package main

import (
	"auth/service/internal/auth_service"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/auth", auth_service.CreateAccount).Methods("POST")
}
