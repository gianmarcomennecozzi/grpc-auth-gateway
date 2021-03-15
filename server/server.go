package server

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"

	pb "github.com/gianmarcomennecozzi/grpc-auth-gateway/proto/todo"
)

type server struct {
	m     sync.RWMutex
	todos map[string]string //map name with id
	pb.UnimplementedTodoServer
}

func NewServer() *server {
	return &server{
		todos: make(map[string]string),
	}
}

func (s *server) AddTodo(ctx context.Context, request *pb.AddTodoRequest) (*pb.AddTodoResponse, error) {
	s.m.Lock()
	defer s.m.Unlock()

	_, ok := s.todos[request.Name]
	if ok {
		return nil, fmt.Errorf("todo already exists")
	}

	id := uuid.New().String()
	s.todos[request.Name] = id
	log.Printf("created new todo [%s]: %s", request.Name, id)
	return &pb.AddTodoResponse{Name: request.Name, Id: id}, nil
}

func (s *server) GetTodos(_ *pb.Empty, srv pb.Todo_GetTodosServer) error {
	s.m.RLock()
	defer s.m.RUnlock()

	for n, id := range s.todos {
		err := srv.Send(&pb.AddTodoResponse{Name: n, Id: id})
		if err != nil {
			return err
		}
	}
	return nil
}
