package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/VIBHOR94/grpc/average/averagepb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connetc %v", err)
	}
	defer cc.Close()

	c := averagepb.NewAverageServiceClient(cc)

	doClientStreaming(c)

}

func doClientStreaming(c averagepb.AverageServiceClient) {
	requests := []*averagepb.AverageRequest{
		&averagepb.AverageRequest{
			Number: 15,
		},
		&averagepb.AverageRequest{
			Number: 18,
		},
		&averagepb.AverageRequest{
			Number: 26,
		},
		&averagepb.AverageRequest{
			Number: 999,
		},
	}
	stream, err := c.ComputeAverage(context.Background())

	if err != nil {
		log.Fatalf("error while calling ComputeAverage: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from ComputeAverage: %v", err)
	}
	fmt.Printf("ComputeAverage Response: %v\n", res)
}
