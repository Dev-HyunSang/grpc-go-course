package main

import (
	"context"
	"log"

	pb "github.com/dev-hyunsang/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
2023/12/12 23:40:49 doSqrt was invoked
2023/12/12 23:40:49 Sqrt: 3.162278
parppppp
Error Message for '-2'
2023/12/12 23:41:54 doSqrt was invoked
2023/12/12 23:41:54 Error message from server: Received a negative number: -2
2023/12/12 23:41:54 Error code from server: InvalidArgument
2023/12/12 23:41:54 We probably sent a negative number!
*/
func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", e)
		}

	}

	log.Printf("Sqrt: %f\n", res.Result)
}
