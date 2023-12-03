package main

import (
	"context"
	"log"

	pb "github.com/dev-hyunsang/grpc-go-course/greet/proto"
)

/*
2023/12/03 23:21:48 Listening on 0.0.0.0:50051
2023/12/03 23:21:50 Greet function was invoked with first_name:"HyunSang"
*/
func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "HyunSang",
	})
	if err != nil {
		log.Fatalf("Colud not greet %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
