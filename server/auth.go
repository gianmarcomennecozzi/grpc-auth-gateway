package server

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	MissingTokenErr = errors.New("no security token provided")
	UnknownTokenErr = errors.New("unknown security token")
)

func getToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", MissingTokenErr
	}

	if len(md["token"]) == 0 {
		return "", MissingTokenErr
	}

	token := md["token"][0]
	if token == "" {
		return "", MissingTokenErr
	}

	return token, nil
}

func GetServer(authCode string) *grpc.Server {
	var opts []grpc.ServerOption

	streamInterceptor := func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

		token, authErr := getToken(stream.Context())
		if authErr != nil {
			return authErr
		}

		if token != authCode {
			return UnknownTokenErr
		}

		return handler(srv, stream)
	}

	unaryInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		token, authErr := getToken(ctx)
		if authErr != nil {
			return nil, authErr
		}

		if token != authCode {
			return nil, UnknownTokenErr
		}

		return handler(ctx, req)
	}

	opts = append([]grpc.ServerOption{
		grpc.StreamInterceptor(streamInterceptor),
		grpc.UnaryInterceptor(unaryInterceptor),
	}, opts...)

	return grpc.NewServer(opts...)
}
