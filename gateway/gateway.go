package gateway

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	pb "github.com/gianmarcomennecozzi/grpc-auth-gateway/proto/todo"
	"github.com/rs/zerolog/log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const port = "8080"

func Run(dialAddr string) error {

	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterTodoHandler(context.Background(), gwmux, conn)
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}

	gatewayAddr := "0.0.0.0:" + port
	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				gwmux.ServeHTTP(w, r)
				return
			}
			log.Print("error! wrong path")
		}),
	}

	log.Printf("Serving gRPC-Gateway on http://", gatewayAddr)
	return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServe())
}
