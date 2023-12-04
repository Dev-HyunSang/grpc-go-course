package main

import (
	"context"
	"io"
	"log"

	pb "github.com/dev-hyunsang/grpc-go-course/greet/proto"
)

/* 2023/12/04 16:44:49 doGreetsManyTimes function was invoked
2023/12/04 16:44:49 GreetManyTimes: Hello HyunSang, number 0
2023/12/04 16:44:49 GreetManyTimes: Hello HyunSang, number 1
2023/12/04 16:44:49 GreetManyTimes: Hello HyunSang, number 2
2023/12/04 16:44:49 GreetManyTimes: Hello HyunSang, number 3
2023/12/04 16:44:49 GreetManyTimes: Hello HyunSang, number 4
2023/12/04 16:44:49 GreetManyTimes: Hello HyunSang, number 5
2023/12/04 16:44:49 GreetManyTimes: Hello HyunSang, number 6
2023/12/04 16:44:49 GreetManyTimes: Hello HyunSang, number 7
2023/12/04 16:44:49 GreetManyTimes: Hello HyunSang, number 8
2023/12/04 16:44:49 GreetManyTimes: Hello HyunSang, number 9
*/

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetsManyTimes function was invoked")

	req := &pb.GreetRequest{
		FirstName: "HyunSang",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		mes, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("GreetManyTimes: %s\n", mes.Result)
	}
}
