/*
*	Author: Cody Botte
*	Purpose: An authentication microservice using gorilla/mux. Users saved via mongodb, passwords encrypted using
* 		 bcrypt, and users identified using json web tokens.
 */

package main

import (
	"github.com/cbotte21/auth-go/internal"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	//Verify enviroment variables exist
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load enviroment variables")
	}
	verifyEnvVariable("mongo_uri")
	verifyEnvVariable("port")
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

func verifyEnvVariable(name string) {
	_, uriPresent := os.LookupEnv(name)
	if !uriPresent {
		log.Fatalf("could not find {" + name + "} environment variable")
	}
}
