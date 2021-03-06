package main

import (
	"log"
	"net"

	"github.com/gianmarcomennecozzi/grpc-auth-gateway/gateway"

	"github.com/gianmarcomennecozzi/grpc-auth-gateway/proto/todo"

	"github.com/gianmarcomennecozzi/grpc-auth-gateway/server"
)

const (
	endpoint = "0.0.0.0:50051"
	secToken = "passcode"
)

func main() {
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := server.GetServer(secToken)
	todo.RegisterTodoServer(s, server.NewServer())

	log.Printf("Serving gRPC on %s", endpoint)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + endpoint)
	log.Fatalln(err)
}
