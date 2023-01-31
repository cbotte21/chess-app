package service

import (
	"github.com/cbotte21/username-go/internal/handler"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Api struct {
	port   int
	router *mux.Router
}

func NewApi(port int) (*Api, bool) {
	api := &Api{}
	api.port = port
	api.router = mux.NewRouter()
	api.RegisterHandlers()
	return api, true
}

func (api *Api) Start() error { //maybe change return to bool
	return http.ListenAndServe(":"+strconv.Itoa(api.port), api.router)
}

func (api *Api) RegisterHandlers() { //Add all API handlers here
	prefix := "/api"

	api.router.HandleFunc(prefix+"/", handler.GetHandler).Methods("GET")
	api.router.HandleFunc(prefix+"/", handler.SetHandler).Methods("POST")
}
