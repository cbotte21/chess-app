package enviroment

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var inited = false

func initEnv() {
	if !inited {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("could not load enviroment variables")
		}
		inited = true
	}
}

func VerifyEnvVariable(name string) {
	initEnv()

	_, uriPresent := os.LookupEnv(name)
	if !uriPresent {
		log.Fatalf("could not find {" + name + "} environment variable")
	}
}

func GetEnvVariable(name string) string {
	return os.Getenv(name)
}
