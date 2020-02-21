package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/veeruns/Go-Lambda/GRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.OpRequest) (*calculatorpb.OpResponse, error) {
	firstOperand := req.GetOperands().GetFirstNumber()
	secondOperand := req.GetOperands().GetSecondNumber()
	operation := req.GetOperation().GetOperation()

	fmt.Printf("Operands are %d and %d and operation is %s\n", firstOperand, secondOperand, operation)
	op := firstOperand + secondOperand
	res := calculatorpb.OpResponse{
		Result: op,
	}
	return &res, nil
}

func (*server) Generic(ctx context.Context, req *calculatorpb.OpRequest) (*calculatorpb.OpResponse, error) {
	firstOperand := req.GetOperands().GetFirstNumber()
	secondOperand := req.GetOperands().GetSecondNumber()
	operation := req.GetOperation().GetOperation()

	fmt.Printf("Operands are %d and %d and operation is %s\n", firstOperand, secondOperand, operation)
	var result int32
	switch operation {
	case "+":
		result = firstOperand + secondOperand
	case "*":
		result = firstOperand * secondOperand
	case "-":
		result = firstOperand - secondOperand
	case "/":
		result = firstOperand / secondOperand
	default:
		result = firstOperand + secondOperand
	}
	//op := firstOperand + secondOperand
	res := calculatorpb.OpResponse{
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
	calculatorpb.RegisterSumServiceServer(derivedserver, &server{})
	calculatorpb.RegisterGenericServiceServer(derivedserver, &server{})
	//calculatorpb.RegisterGreetServiceServer(derivedserver, &server{})

	if err := derivedserver.Serve(lis); err != nil {
		log.Fatalf("Failed to start serve %v\n", err)
	}
}
