package main

import (
	hive "github.com/cbotte21/hive-go/pb"
	"github.com/cbotte21/judicial-go/internal"
	pb "github.com/cbotte21/judicial-go/pb"
	"github.com/cbotte21/microservice-common/pkg/enviroment"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//Verify enviroment variables exist
	enviroment.VerifyEnvVariable("port")
	enviroment.VerifyEnvVariable("hive_port")

	//Get port
	port := enviroment.GetEnvVariable("port")

	//Setup tcp listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port: %d", port)
	}
	grpcServer := grpc.NewServer()

	//Register handler(s) to attach
	hiveClient := hive.NewHiveServiceClient(getHiveConn())
	//Initialize judicial
	jury := internal.NewJudicial(&hiveClient)

	pb.RegisterJudicialServiceServer(grpcServer, &jury)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize grpc server.")
	}
}

func getHiveConn() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+enviroment.GetEnvVariable("hive_port"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	return conn
}
