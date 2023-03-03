package service

import (
	"github.com/cbotte21/auth-go/internal/handler"
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

func (api *Api) Start() error {
	return http.ListenAndServe(":"+strconv.Itoa(api.port), api.router)
}

func (api *Api) RegisterHandlers() { //Add all API handlers here
	api.router.HandleFunc("/", handler.IndexHandler).Methods("GET")
	//User lifecycle
	api.router.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	api.router.HandleFunc("/signup", handler.SignupHandler).Methods("POST")
	api.router.HandleFunc("/verify", handler.VerifyHandler).Methods("POST")

}
