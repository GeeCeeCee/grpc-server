package main

import (
	"fmt"
	"log"
	"net"

	"github.com/GeeCeeCee/grpc-server/person"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Server started successfully")
	}

	s := person.Server{}

	grpcServer := grpc.NewServer()

	person.RegisterSayHelloServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
