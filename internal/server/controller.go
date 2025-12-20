package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// this type like enum in C++ define fixed arguments
type HttpType string

const (
	GET    HttpType = "GET"
	POST   HttpType = "POST"
	PUT    HttpType = "PUT"
	DELETE HttpType = "DELETE"
)

type HttpMethod struct {
	responseType HttpType
	method       func(http.ResponseWriter, *http.Request)
}

func NewHttpMethod(typeMethod HttpType, method func(http.ResponseWriter, *http.Request)) *HttpMethod {
	return &HttpMethod{
		responseType: typeMethod,
		method:       method,
	}
}

type Controller struct {
	Rout   mux.Router
	Action map[string]func(http.ResponseWriter, *http.Request) //init all action for REST pattern. ["actionName"]function
}

// Init router with map of func, for work with REST
func NewController(r mux.Router, actionSet map[string]func(http.ResponseWriter, *http.Request)) *Controller {
	return &Controller{
		Rout:   r,
		Action: actionSet,
	}
}
