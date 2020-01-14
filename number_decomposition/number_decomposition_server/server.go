package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"strconv"
	"time"

	"github.com/VIBHOR94/grpc/number_decomposition/number_decompositionpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Decomposition(req *number_decompositionpb.NumberDecompositionRequest, stream number_decompositionpb.NumberDecompositionService_DecompositionServer) error {
	fmt.Printf("Decomposition function was invoked with %v\n", req)
	Number := req.GetNumber()
	if Number < 2 {
		res := &number_decompositionpb.NumberDecompositionResponse{
			Result: "No prime factors avaialable",
		}
		stream.Send(res)
	} else {
		checkNum := int32(2)
		for checkNum <= Number {
			if IsPrimeSqrt(int(checkNum)) {
				for Number%checkNum == 0 {
					res := &number_decompositionpb.NumberDecompositionResponse{
						Result: strconv.Itoa(int(checkNum)),
					}
					Number /= checkNum
					stream.Send(res)
					time.Sleep(1000 * time.Millisecond)
				}
			}
			checkNum++
		}
	}
	return nil
}

func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	number_decompositionpb.RegisterNumberDecompositionServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
