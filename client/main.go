package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	tracepb "micro-gogrpc-gateway/server/proto/gen/go"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect server: %v", err)
	}

	tsClient := tracepb.NewTraceServiceClient(conn)

	r, err := tsClient.GetTrace(context.Background(), &tracepb.GetTraceRequest{
		Id: "trace0",
	})
	if err != nil {
		log.Fatalf("cannot call GetTrace: %v", err)
	}
	fmt.Println("r", r)
}
