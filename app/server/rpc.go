package server

import (
	"github.com/txze/visitorpb/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"visitor/service/rpc"

	"github.com/sirupsen/logrus"
)

func startRpcServer() {
	listen, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	visitorpb.RegisterVisitorServer(s, &rpc.Service{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	logrus.Println("rpc start")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
