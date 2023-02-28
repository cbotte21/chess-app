/*
*	Author: Cody Botte
*	Purpose: An authentication microservice using gorilla/mux. Users saved via mongodb, passwords encrypted using
* 		 bcrypt, and users identified using json web tokens.
 */

package main

import (
	"fmt"
	"github.com/cbotte21/auth-go/internal"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	//Verify env variables exist
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load enviroment variables!")
	}
	_, uriPresent := os.LookupEnv("auth_mongo_uri")

	if !uriPresent {
		fmt.Println("could not find {auth_mongo_uri} environment variable")
		os.Exit(1)
	}

	api, res := service.NewApi(5000)
	if !res || api.Start() != nil { //Start API Listener
		log.Fatal("Failed to initialize API.")
	}
}
