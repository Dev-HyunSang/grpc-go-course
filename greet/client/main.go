package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/dev-hyunsang/grpc-go-course/greet/proto"
)

var (
	addr string = "localhost:50051"
	tls  bool   = true
)

func main() {
	opts := []grpc.DialOption{}

	if tls {
		cretFile := "ssl/server.crt"
		creds, err := credentials.NewClientTLSFromFile(cretFile, "")
		if err != nil {
			log.Fatalf("Failed loading certficates: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))

	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	// doGreetWithDeadline(c, 5*time.Second)
}
