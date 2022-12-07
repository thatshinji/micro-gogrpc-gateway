package main

import (
	"google.golang.org/grpc"
	"log"
	tracepb "micro-gogrpc-gateway/server/proto/gen/go"
	trace "micro-gogrpc-gateway/server/traceService"
	"net"
)

func main() {
	log.SetFlags(log.Lshortfile)
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	tracepb.RegisterTraceServiceServer(s, &trace.Service{})
	log.Fatalf("err: %v", s.Serve(lis))
}
