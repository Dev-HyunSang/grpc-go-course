package main

import (
	"context"
	"log"

	pb "github.com/dev-hyunsang/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Hyunsang",
	})
	if err != nil {
		log.Fatalf("Colud not greet %v\n", err)
	}

	log.Println("Greeting: %s\n", res.Result)
}
