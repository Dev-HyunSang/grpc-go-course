package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dev-hyunsang/grpc-go-course/greet/proto"
)

/* 2023/12/04 21:36:46 doLogGreet was invoked
2023/12/04 21:36:46 Sending req: first_name:"HyunSang"
2023/12/04 21:36:47 Sending req: first_name:"John"
2023/12/04 21:36:48 Sending req: first_name:"Tomo"
2023/12/04 21:36:49 LongGreet: Hello HyunSang!
Hello John!
Hello Tomo! */

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLogGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "HyunSang"},
		{FirstName: "John"},
		{FirstName: "Tomo"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreetL %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while reciving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
