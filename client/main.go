package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/gianmarcomennecozzi/grpc-auth-gateway/proto"
	"google.golang.org/grpc"
)

const endpoint = "localhost:50051"

func main() {

	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewTodoClient(conn)
	resp, err := client.AddTodo(context.Background(), &pb.AddTodoRequest{Name: "random"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New id created for [%s]: %s", resp.Name, resp.Id)
}
