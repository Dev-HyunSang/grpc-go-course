package main

import (
	"log"
	"net"

	pb "github.com/dev-hyunsang/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

const (
	addr = "0.0.0.0:50051"
)

type Server struct {
	pb.CalculatorServiceServer
}

/* 2023/12/03 23:08:57 Listening on 0.0.0.0:50051
2023/12/03 23:09:35 Sum function was invkoed with first_number:1 second_number:1 */

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
