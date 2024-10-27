package server

import (
	"context"
	"fmt"
	"net"
	"os"

	pb "github.com/nuzur/extension-sdk/idl/gen"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Params struct {
	fx.In
	Logger    *zap.Logger
	Lifecycle fx.Lifecycle
	Server    pb.NuzurExtensionServer
}

func New(params Params) *grpc.Server {
	log := params.Logger
	s := grpc.NewServer()
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// GRPC port from config
			port := os.Getenv("PORT")

			// proto server
			lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
			if err != nil {
				log.Error("failed to listen: %v", zap.Error(err))
				return err
			}

			// register grpc servers
			grpc_health_v1.RegisterHealthServer(s, health.NewServer())
			pb.RegisterNuzurExtensionServer(s, params.Server)
			reflection.Register(s)

			log.Info("GRPC Server listening at %v", zap.Any("addr", lis.Addr()))

			go s.Serve(lis)

			return nil

		},
		OnStop: func(ctx context.Context) error {
			s.Stop()
			return nil
		},
	})

	return s
}