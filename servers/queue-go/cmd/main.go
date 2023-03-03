package main

import (
	hive "github.com/cbotte21/hive-go/pb"
	"github.com/cbotte21/judicial-go/internal"
	pb "github.com/cbotte21/judicial-go/pb"
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
	//Get port
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		log.Fatalf("could not parse {port} enviroment variable")
	}

	//Setup tcp listener
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Failed to listen on port: %d", port)
	}
	grpcServer := grpc.NewServer()

	//Register handler(s) to attach
	hiveClient := hive.NewHiveServiceClient(getConn())
	//Initialize judicial
	jury := internal.NewJudicial(&hiveClient)

	pb.RegisterJudicialServiceServer(grpcServer, &jury)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize grpc server.")
	}
}

func getConn() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
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
