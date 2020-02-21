package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/veeruns/Go-Lambda/GRPC/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstname := req.GetGreeting().GetFirstName()
	lastname := req.GetGreeting().GetLastName()

	result := "Hello" + firstname + " " + lastname

	res := greetpb.GreetResponse{
		Result: result,
	}

	return &res, nil

}
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
