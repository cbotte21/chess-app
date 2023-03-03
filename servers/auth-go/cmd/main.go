/*
*	Author: Cody Botte
*	Purpose: An authentication microservice using gorilla/mux. Users saved via mongodb, passwords encrypted using
* 		 bcrypt, and users identified using json web tokens.
 */

package main

import (
	"github.com/cbotte21/auth-go/internal"
	"github.com/cbotte21/microservice-common/pkg/enviroment"
	"log"
	"os"
	"strconv"
)

func main() {
	//Verify enviroment variables exist
	enviroment.VerifyEnvVariable("mongo_uri")
	enviroment.VerifyEnvVariable("port")
	enviroment.VerifyEnvVariable("jwt_secret")
	//Get port
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		log.Fatalf("could not parse {auth_port} enviroment variable")
	}
	//Start API
	api, res := service.NewApi(port)
	if !res || api.Start() != nil { //Start API Listener
		log.Fatal("Failed to initialize API.")
	}
}
