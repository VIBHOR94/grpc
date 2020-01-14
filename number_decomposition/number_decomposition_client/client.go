package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/VIBHOR94/grpc/number_decomposition/number_decompositionpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connetc %v", err)
	}
	defer cc.Close()

	c := number_decompositionpb.NewNumberDecompositionServiceClient(cc)
	doServerStreaming(c)
}

func doServerStreaming(c number_decompositionpb.NumberDecompositionServiceClient) {
	req := &number_decompositionpb.NumberDecompositionRequest{
		Number: 898,
	}
	resStream, err := c.Decomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Decomposition RPC: %v", c)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream %v", err)
		}
		log.Printf("Response from Decomposition: %v", msg)
	}
}
