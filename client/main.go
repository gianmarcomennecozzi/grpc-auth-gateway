package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/gianmarcomennecozzi/grpc-auth-gateway/proto/todo"

	"google.golang.org/grpc"
)

const endpoint = "localhost:50051"

type Creds struct {
	Token    string
	Insecure bool
}

func (c Creds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"token": c.Token,
	}, nil
}

func (c Creds) RequireTransportSecurity() bool {
	return !c.Insecure
}

func main() {

	conn, err := grpc.Dial(endpoint, grpc.WithInsecure(), grpc.WithPerRPCCredentials(Creds{Insecure: true, Token: "passcode"}))
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewTodoClient(conn)
	ctx := context.Background()

	resp, err := client.AddTodo(ctx, &pb.AddTodoRequest{Name: "random"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New todo created for [%s]: %s", resp.Name, resp.Id)

	stream, err := client.GetTodos(ctx, &pb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		r, err := stream.Recv()
		if err != nil {
			break
		}
		log.Print(r)
	}
}
