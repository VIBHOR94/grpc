package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"google.golang.org/grpc/codes"

	square_rootpb "github.com/VIBHOR94/grpc/squareroot/squarerootpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) ComputeSquareRoot(ctx context.Context, req *square_rootpb.SquareRootRequest) (*square_rootpb.SquarerootResponse, error) {
	fmt.Println("Recieved SquareRoot RPC")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Recieved a negatice number: %v", number),
		)
	}
	return &square_rootpb.SquarerootResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	square_rootpb.RegisterSquareRootServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
