package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/VIBHOR94/grpc/sum/sumpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Printf("Sum function was involved with %v\n", req)
	firstNumber := req.GetAddition().GetFirstNumber()
	secondNumber := req.GetAddition().GetSecondNumber()
	result := firstNumber + secondNumber
	res := &sumpb.SumResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
