package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	tracepb "micro-gogrpc-gateway/server/proto/gen/go"
	trace "micro-gogrpc-gateway/server/traceService"
	"net"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGPRCGatewayAPI()
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	tracepb.RegisterTraceServiceServer(s, &trace.Service{})
	log.Fatalf("err: %v", s.Serve(lis))
}

func startGPRCGatewayAPI() {
	log.SetFlags(log.Lshortfile)
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux()
	err := tracepb.RegisterTraceServiceHandlerFromEndpoint(
		c,
		mux,
		"localhost:8081",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalf("cannot start grpc gateway %v:", err)
	}
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("cannot listen and server, err: %v", err)
	}
}
