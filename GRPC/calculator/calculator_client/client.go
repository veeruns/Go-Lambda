package main

import (
	"context"
	"fmt"
	"log"

	"github.com/veeruns/Go-Lambda/GRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Mello")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {

		log.Fatalf("Could not connect to server %v\n", err)
	}
	defer conn.Close()
	//greetclient := greetpb.NewGreetServiceClient(conn)
	calcclient := calculatorpb.NewSumServiceClient(conn)
	generciclient := calculatorpb.NewGenericServiceClient(conn)
	doUnary(calcclient)
	doUnaryGeneric(generciclient, 10, 5, "/")
	doUnaryGeneric(generciclient, 10, 5, "+")

}

func doUnary(con calculatorpb.SumServiceClient) {
	fmt.Println("Starting to do unary api to greet client")
	req := &calculatorpb.OpRequest{
		Operands: &calculatorpb.Operands{
			FirstNumber:  5,
			SecondNumber: 10,
		},
		Operation: &calculatorpb.Operation{
			Operation: "+",
		},
	}
	resp, err := con.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Sum %v\n", err)
	}

	fmt.Printf("Response from GRPC Calculator server is %d\n", resp.Result)
}

func doUnaryGeneric(con calculatorpb.GenericServiceClient, firstoperand, secondoperand int32, operation string) {
	fmt.Printf("Starting to do unary api to Generic client %d and %d and %s", firstoperand, secondoperand, operation)
	req := &calculatorpb.OpRequest{
		Operands: &calculatorpb.Operands{
			FirstNumber:  firstoperand,
			SecondNumber: secondoperand,
		},
		Operation: &calculatorpb.Operation{
			Operation: operation,
		},
	}
	resp, err := con.Generic(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Sum %v\n", err)
	}

	fmt.Printf("Response from GRPC Calculator server is %d\n", resp.Result)
}
