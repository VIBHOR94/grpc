package main

import (
	"context"
	"fmt"
	"log"

	"github.com/VIBHOR94/grpc/sum/sumpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	c := sumpb.NewSumServiceClient(cc)
	doUnary(c)
}

func doUnary(c sumpb.SumServiceClient) {
	req := &sumpb.SumRequest{
		Addition: &sumpb.Numbers{
			FirstNumber:  58,
			SecondNumber: 112,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", c)
	}
	log.Printf("Response from Sum: %v", res.Result)
}
