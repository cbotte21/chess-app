package main

import (
	"github.com/cbotte21/hive-go/internal"
	"github.com/cbotte21/hive-go/internal/jwtParser"
	"github.com/cbotte21/hive-go/internal/playerbase"
	"github.com/cbotte21/hive-go/pb"
	judicial "github.com/cbotte21/judicial-go/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	//Verify enviroment variables exist
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load enviroment variables")
	}
	verifyEnvVariable("port")
	verifyEnvVariable("secret")
	//Get port
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		log.Fatalf("could not parse {auth_port} enviroment variable")
	}

	//Setup tcp listener
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Failed to listen on port: %d", port)
	}
	grpcServer := grpc.NewServer()

	//Register handlers to attach
	playerBase := playerbase.PlayerBase{}
	jwtRedeemer := jwtParser.NewJwtSecret(os.Getenv("secret"))
	judicialClient := judicial.NewJudicialServiceClient(getConn())
	//Initialize hive
	hive := internal.NewHive(&playerBase, &jwtRedeemer, &judicialClient)

	pb.RegisterHiveServiceServer(grpcServer, &hive)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize grpc server.")
	}
}

func getConn() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	return conn
}

func verifyEnvVariable(name string) {
	_, uriPresent := os.LookupEnv(name)
	if !uriPresent {
		log.Fatalf("could not find {" + name + "} environment variable")
	}
}
