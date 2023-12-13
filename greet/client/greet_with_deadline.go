package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dev-hyunsang/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "HyunSang",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			// 강의 상에서는 status.DeadlineExceeded로 코딩하고 있으나, status 패키지 오류가 존재함.
			// codes 패키지를 사용하여 해결하면 됨.
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline execeeded!")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
