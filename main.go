package main

import (
	"log"
	"net"

	pb "github.com/gianmarcomennecozzi/grpc-auth-gateway/proto"
	"github.com/gianmarcomennecozzi/grpc-auth-gateway/server"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoServer(s, server.NewServer())
	log.Printf("server listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
