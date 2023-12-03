package main

import (
	"context"
	"log"

	pb "github.com/dev-hyunsang/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invkoed with %v\n", in)

	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
