package main

import (
	"fmt"
	"log"
	"net"

	"github.com/veeruns/Go-Lambda/GRPC/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello GRPC server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Fail to listen %v\n", err)
	}

	derivedserver := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(derivedserver, &server{})
	if err := derivedserver.Serve(lis); err != nil {
		log.Fatalf("Failed to start serve %v\n", err)
	}
}
