package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/LidorAlmkays/MineServerForge/config"
	"github.com/LidorAlmkays/MineServerForge/pkg/logger"
)


type Server struct {
	ctx     context.Context
	mux     *http.ServeMux
	cfg     *config.Config
	l       logger.Logger
}

func NewServer(ctx context.Context, cfg *config.Config, l logger.Logger) *Server {
	mux := http.NewServeMux()
	return &Server{mux: mux, cfg: cfg, l: l, ctx:ctx}
}

func (s *Server) ListenAndServe() error {
	handler := s.addRoutes()
	s.l.Message("Server ready to receive REST requests, on port: " + strconv.Itoa(s.cfg.ServiceConfig.Port))
	err := http.ListenAndServe(":"+strconv.Itoa(s.cfg.ServiceConfig.Port), handler)
	if err != nil {
		return err
	}
	return nil
}
