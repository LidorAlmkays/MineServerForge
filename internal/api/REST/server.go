package rest

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/LidorAlmkays/MineServerForge/config"
	"github.com/LidorAlmkays/MineServerForge/internal/api"
	"github.com/LidorAlmkays/MineServerForge/internal/application"
	"github.com/LidorAlmkays/MineServerForge/pkg/logger"
)

type Server struct {
	ctx        context.Context
	mux        *http.ServeMux
	cfg        *config.Config
	l          logger.Logger
	httpServer *http.Server
	s          application.ServerConfigDataManager
}

func NewServer(ctx context.Context, cfg *config.Config, l logger.Logger, s application.ServerConfigDataManager) api.BaseServer {
	mux := http.NewServeMux()
	return &Server{ctx: ctx, mux: mux, cfg: cfg, l: l, s: s}
}

func (s *Server) ListenAndServe() error {

	s.httpServer = &http.Server{
		Addr: ":" + strconv.Itoa(s.cfg.ServiceConfig.HttpPort),
	}
	handler, err := s.addRoutes()
	if err != nil {
		return err
	}
	s.httpServer.Handler = handler

	s.l.Message("Server ready to receive REST requests, on port: " + strconv.Itoa(s.cfg.ServiceConfig.HttpPort))
	err = s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return errors.New("failed to serve rest server: " + err.Error())
	}
	return nil
}

func (s *Server) Shutdown() error {
	s.l.Message("Shutting down REST server...")
	return s.httpServer.Shutdown(s.ctx) // Gracefully shut down HTTP server
}
