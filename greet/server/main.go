package main

import (
	"log"
	"net"

	pb "github.com/dev-hyunsang/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

const (
	addr = "0.0.0.0:50051"
)

type Server struct {
	pb.GreetServiceServer
}

/*
2023/12/03 23:21:50 doGreet was invoked
2023/12/03 23:21:50 Greeting: Hello HyunSang
*/
func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
