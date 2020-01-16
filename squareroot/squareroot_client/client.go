package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	square_rootpb "github.com/VIBHOR94/grpc/squareroot/squarerootpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connetc %v", err)
	}
	defer cc.Close()

	c := square_rootpb.NewSquareRootServiceClient(cc)
	doErrorUnary(c)
}

func doErrorUnary(c square_rootpb.SquareRootServiceClient) {
	fmt.Println("Starting to do a SquareRoot Unary RPC...")
	doErrorCall(c, 10)
	doErrorCall(c, -7)
}

func doErrorCall(c square_rootpb.SquareRootServiceClient, n int32) {
	res, err := c.ComputeSquareRoot(context.Background(), &square_rootpb.SquareRootRequest{Number: n})

	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number.")
			}
		} else {
			log.Fatalf("Big error calling Squareroot: %v", err)
		}
	} else {
		fmt.Printf("Result of square root of %v: %v\n", n, res.GetResult())
	}
}
