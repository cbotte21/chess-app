package main

import (
	hive "github.com/cbotte21/hive-go/pb"
	"github.com/cbotte21/judicial-go/internal"
	pb "github.com/cbotte21/judicial-go/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

const (
	PORT int = 9001
)

func main() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Fatalf("Failed to listen on port: %d", PORT)
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
