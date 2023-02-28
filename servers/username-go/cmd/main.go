package main

import (
	"github.com/cbotte21/username-go/internal"
	"log"
)

//TODO: Should really be kept in a relational database

func main() {
	api, res := service.NewApi(5000)
	if !res || api.Start() != nil { //Start API Listener
		log.Fatal("Failed to initialize API.")
	}
}
