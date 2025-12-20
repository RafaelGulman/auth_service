package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

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
