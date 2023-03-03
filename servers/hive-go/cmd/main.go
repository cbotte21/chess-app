package main

import (
	"github.com/cbotte21/hive-go/internal"
	"github.com/cbotte21/hive-go/internal/playerbase"
	"github.com/cbotte21/hive-go/pb"
	judicial "github.com/cbotte21/judicial-go/pb"
	"github.com/cbotte21/microservice-common/pkg/enviroment"
	"github.com/cbotte21/microservice-common/pkg/jwtParser"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//Verify enviroment variables exist
	enviroment.VerifyEnvVariable("port")
	enviroment.VerifyEnvVariable("jwt_secret")
	enviroment.VerifyEnvVariable("judicial_port")

	port := enviroment.GetEnvVariable("port")
	//Setup tcp listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port: %d", port)
	}
	grpcServer := grpc.NewServer()

	//Register handlers to attach
	playerBase := playerbase.PlayerBase{}
	jwtRedeemer := jwtParser.JwtSecret(enviroment.GetEnvVariable("secret"))
	judicialClient := judicial.NewJudicialServiceClient(getJudicialConn(port))
	//Initialize hive
	hive := internal.NewHive(&playerBase, &jwtRedeemer, &judicialClient)

	pb.RegisterHiveServiceServer(grpcServer, &hive)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize grpc server.")
	}
}

func getJudicialConn(port string) *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+enviroment.GetEnvVariable("judicial_port"), grpc.WithInsecure()) //TODO: variable for this port
	if err != nil {
		log.Fatalf(err.Error())
	}
	return conn
}
