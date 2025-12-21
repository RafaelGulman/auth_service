package main

import (
	"auth/service/internal/auth_service"
	"auth/service/internal/server"

	"github.com/gorilla/mux"
)

func main() {
	methodsStructs := make(map[string]server.HttpMethod)

	methodsStructs["CreateAccount"] = *server.NewHttpMethod(server.POST, auth_service.CreateAccount, "/accounts")
	methodsStructs["GetAccounts"] = *server.NewHttpMethod(server.GET, auth_service.GetAccounts, "/accounts")

	mainServer := server.NewController(*mux.NewRouter(), methodsStructs)

	mainServer.Start(":8080")
}
