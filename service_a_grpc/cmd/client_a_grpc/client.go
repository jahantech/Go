package main

import (
	"log"
	accumulator "service_a_grpc/api/accumulator"
	internal "service_a_grpc/internal"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(internal.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := accumulator.NewAccumulatorClient(conn)

	_, err = c.Accumulate(context.Background(), &accumulator.AccumulatorRequest{Number: 2})
	if err != nil {
		log.Fatalf("Error when calling Accumulate: %s", err)
	}

}
