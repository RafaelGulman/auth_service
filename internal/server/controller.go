package server

import (
	"fmt"
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
	httpPath     string
}

func NewHttpMethod(typeMethod HttpType, method func(http.ResponseWriter, *http.Request), httpPath string) *HttpMethod {
	return &HttpMethod{
		responseType: typeMethod,
		method:       method,
		httpPath:     httpPath,
	}
}

type Controller struct {
	rout   mux.Router
	action map[string]HttpMethod //init all action for REST pattern. ["actionName"]function
}

// Init router with map of func, for work with REST
func NewController(r mux.Router, actionSet map[string]HttpMethod) *Controller {
	return &Controller{
		rout:   r,
		action: actionSet,
	}
}

func (c *Controller) Start(ip string) {
	for _, f := range c.action {
		c.rout.HandleFunc(f.httpPath, f.method).Methods(string(f.responseType))
	}

	fmt.Printf("Сервер запущен на %s", ip)
	http.ListenAndServe(ip, &c.rout)
}
