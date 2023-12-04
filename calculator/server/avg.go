package main

import (
	"io"
	"log"

	pb "github.com/dev-hyunsang/grpc-go-course/calculator/proto"
)

/*
2023/12/04 22:05:39 Receiving number: 3
2023/12/04 22:05:39 Receiving number: 5
2023/12/04 22:05:39 Receiving number: 9
2023/12/04 22:05:39 Receiving number: 54
2023/12/04 22:05:39 Receiving number: 23
*/

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")

	sum := 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving number: %d\n", req.Number)
		sum += int(req.Number)
		count++
	}

}
