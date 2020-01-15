package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/VIBHOR94/grpc/average/averagepb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) ComputeAverage(stream averagepb.AverageService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked with stream\n")
	numberInputs := []int32{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result := fetchAverage(numberInputs)
			return stream.SendAndClose(&averagepb.AverageResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v:", err)
		}

		numberInputs = append(numberInputs, req.GetNumber())
	}
}

func fetchAverage(arrayOfInputs []int32) float64 {
	if len(arrayOfInputs) == 0 {
		return 0.0
	}
	sum := int32(0)
	for _, value := range arrayOfInputs {
		sum += value
	}
	return float64(sum) / float64(len(arrayOfInputs))
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	averagepb.RegisterAverageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
