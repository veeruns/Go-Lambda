package main

import (
	"context"
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
	greetclient := greetpb.NewGreetServiceClient(conn)

	doUnary(greetclient)

}

func doUnary(con greetpb.GreetServiceClient) {
	fmt.Println("Starting to do unary api to greet server")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Veeru",
			LastName:  "Natarajan",
		},
	}
	resp, err := con.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling greet %v\n", err)
	}

	fmt.Printf("Response from GRPC server is %s\n", resp.Result)
}
