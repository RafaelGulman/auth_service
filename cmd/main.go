package main

import (
	"auth/service/internal/auth_service"
	"auth/service/internal/server"

	"github.com/gorilla/mux"
)

func main() {
	//define functions which will be use
	methodsStructs := make(map[string]server.HttpMethod)
	methodsStructs["CreateAccount"] = *server.NewHttpMethod(server.POST, auth_service.CreateAccount, "/CreateAccount")
	methodsStructs["GetAccounts"] = *server.NewHttpMethod(server.GET, auth_service.GetAccounts, "/GetAccounts")
	methodsStructs["EnterAccount"] = *server.NewHttpMethod(server.POST, auth_service.EnterAccount, "/EnterAccount")

	//Init and start service
	mainServer := server.NewController(*mux.NewRouter(), methodsStructs)
	mainServer.Start(":8081")
}
