package grpc

import (
	"context"
	"errors"
	"net"
	"strconv"

	"github.com/LidorAlmkays/MineServerForge/config"
	"github.com/LidorAlmkays/MineServerForge/internal/api"
	"github.com/LidorAlmkays/MineServerForge/internal/api/GRPC/pb"
	"github.com/LidorAlmkays/MineServerForge/internal/api/GRPC/servers"
	"github.com/LidorAlmkays/MineServerForge/internal/application/serverfeaturedatamanager"
	"github.com/LidorAlmkays/MineServerForge/pkg/logger"
	"google.golang.org/grpc"
)

type Server struct {
	ctx        context.Context
	cfg        *config.Config
	l          logger.Logger
	grpcServer *grpc.Server
	lis        net.Listener
	sFeatures  serverfeaturedatamanager.ServerFeaturesDataManager
}

func NewServer(ctx context.Context, cfg *config.Config, l logger.Logger, sFeatures serverfeaturedatamanager.ServerFeaturesDataManager) api.BaseServer {
	var opts []grpc.ServerOption = []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	return &Server{ctx: ctx, cfg: cfg, l: l, grpcServer: grpcServer, sFeatures: sFeatures}
}

func (s *Server) ListenAndServe() error {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(s.cfg.ServiceConfig.GrpcPort))
	if err != nil {
		return err
	}
	s.lis = lis

	s.createRoutes()

	s.l.Message("Server ready to receive GRPC requests, on port: " + strconv.Itoa(s.cfg.ServiceConfig.GrpcPort))
	if err := s.grpcServer.Serve(lis); err != nil {
		return errors.New("failed to serve gRPC server over port " + strconv.Itoa(s.cfg.ServiceConfig.GrpcPort) + ", the error: " + err.Error())
	}
	return nil
}

func (s *Server) createRoutes() {
	pb.RegisterUploadFeatureDataServer(s.grpcServer, servers.NewFeatureDataServer(s.l, s.sFeatures))
}

func (s *Server) Shutdown() error {
	s.l.Message("Shutting down gRPC server...")
	s.grpcServer.GracefulStop()
	return nil
}
