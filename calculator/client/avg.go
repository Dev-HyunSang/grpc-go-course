package main

import (
	"context"
	"log"

	pb "github.com/dev-hyunsang/grpc-go-course/calculator/proto"
)

/*
2023/12/04 22:05:39 doAvg was invoked
2023/12/04 22:05:39 Sending number: 3
2023/12/04 22:05:39 Sending number: 5
2023/12/04 22:05:39 Sending number: 9
2023/12/04 22:05:39 Sending number: 54
2023/12/04 22:05:39 Sending number: 23
2023/12/04 22:05:39 Avg: 18.800000
*/

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while opening the stream: %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)

		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving server: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}
