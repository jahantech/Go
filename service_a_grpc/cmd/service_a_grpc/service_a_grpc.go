package main

import (
	"context"
	"fmt"
	"log"
	"net"
	apb "service_a_grpc/api/accumulator"
	internal "service_a_grpc/internal"

	"google.golang.org/grpc"
)

type server struct {
	apb.UnimplementedAccumulatorServer
}

func (s *server) Accumulate(ctx context.Context, in *apb.AccumulatorRequest) (*apb.AccumulatorResponse, error) {
	internal.Count = internal.Count + in.GetNumber()
	fmt.Println(internal.Count)
	return &apb.AccumulatorResponse{Number: in.GetNumber()}, nil
}

func main() {
	lis, err := net.Listen("tcp", internal.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Running Server....")
	s := grpc.NewServer()
	apb.RegisterAccumulatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
