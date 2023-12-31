package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/dev-hyunsang/grpc-go-course/calculator/proto"
)

var addr string = "localhost:50051"

/* 2023/12/03 23:09:35 doSum was invoked
2023/12/03 23:09:35 Sum: 2 */

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	// doSum(c)
	// doPrimes(c)
	// doAvg(c)
	// doMax(c)
	doSqrt(c, -2)
}
