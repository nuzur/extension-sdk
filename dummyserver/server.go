package dummyserver

import (
	"context"
	"errors"
	"fmt"
	"net"

	pb "github.com/nuzur/extension-sdk/proto_deps/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedNuzurProductServer
	grpcServer *grpc.Server
}

const grpcPort = "5505"

func New() (*Server, error) {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			meta, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, errors.New("missing metadata")
			}
			if len(meta["authorization"]) != 1 {
				return nil, errors.New("[sdk] authorization metadata not found ")
			}
			return handler(ctx, req)
		}),
	)

	// proto server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		return nil, err
	}

	// register grpc servers
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	server := &Server{
		grpcServer: s,
	}
	pb.RegisterNuzurProductServer(s, server)
	reflection.Register(s)

	go s.Serve(lis)
	return server, nil
}

func (s *Server) Stop() {
	s.grpcServer.Stop()
}

func (s *Server) Address() *string {
	addr := fmt.Sprintf("localhost:%s", grpcPort)
	return &addr
}
