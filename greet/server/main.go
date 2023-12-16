package main

import (
	"log"
	"net"

	pb "github.com/dev-hyunsang/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	opts := []grpc.ServerOption{}
	tls := true

	if tls {
		cretFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(cretFile, keyFile)
		if err != nil {
			log.Fatalf("Failed loading certficates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
