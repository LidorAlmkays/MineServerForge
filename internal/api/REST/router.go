package rest

import (
	"net/http"

	"github.com/LidorAlmkays/MineServerForge/dtos"
	"github.com/LidorAlmkays/MineServerForge/internal/api/REST/handlers"
	"github.com/LidorAlmkays/MineServerForge/internal/api/REST/middleware"
	"github.com/rs/cors"
)

func (s *Server) addRoutes() http.Handler {
	s.l.Message("Setting up http routes")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	h := handlers.NewHandler(s.cfg, s.ctx, s.l)
	s.mux.Handle("POST /minecraft/create", middleware.DecodeJSONBody[dtos.CreateMinecraftServerDTO]("dto", http.HandlerFunc(h.RunMinecraftServer)))
	return c.Handler(s.mux)
}
