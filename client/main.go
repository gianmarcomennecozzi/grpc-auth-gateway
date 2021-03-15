package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gianmarcomennecozzi/grpc-auth-gateway/proto/todo"

	"google.golang.org/grpc"
)

const endpoint = "localhost:50051"

func main() {

	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := todo.NewTodoClient(conn)
	resp, err := client.AddTodo(context.Background(), &todo.AddTodoRequest{Name: "random"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New id created for [%s]: %s", resp.Name, resp.Id)
}
