package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/dev-hyunsang/grpc-go-course/greet/proto"
)

/*
2023/12/11 21:06:05 doGreetEveryone was invoked
2023/12/11 21:06:05 Sending request: first_name:"HyunSang"
2023/12/11 21:06:05 Recevied: Hello HyunSang!
2023/12/11 21:06:06 Sending request: first_name:"John"
2023/12/11 21:06:06 Recevied: Hello John!
2023/12/11 21:06:07 Sending request: first_name:"Tomo"
2023/12/11 21:06:07 Recevied: Hello Tomo!
*/
func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "HyunSang"},
		{FirstName: "John"},
		{FirstName: "Tomo"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Recevied: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
