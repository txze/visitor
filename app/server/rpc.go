package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"visitor/service/rpc"
	"visitor/visitorpb"
)

func startRpcServer() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	visitorpb.RegisterVisitorServer(s, &rpc.Service{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
