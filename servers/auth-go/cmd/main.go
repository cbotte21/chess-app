/*
*	Author: Cody Botte
*	Purpose: An authentication microservice using gorilla/mux. Users saved via mongodb, passwords encrypted using
* 		 bcrypt, and users identified using json web tokens.
 */

package main

import (
	"github.com/cbotte21/auth-go/internal"
	"log"
)

func main() {
	api, res := service.NewApi(5000)
	if !res || api.Start() != nil { //Start API Listener
		log.Fatal("Failed to initialize API.")
	}
}
