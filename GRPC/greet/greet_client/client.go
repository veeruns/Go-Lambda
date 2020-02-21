package main

import (
	"fmt"
	"log"

	"github.com/veeruns/Go-Lambda/GRPC/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Mello")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {

		log.Fatalf("Could not connect to server %v\n", err)
	}
	defer conn.Close()
	greetclient := greetpb.GreetServiceClient(conn)

}
