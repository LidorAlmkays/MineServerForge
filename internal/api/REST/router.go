package rest

import (
	"net/http"

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

	// h := handlers.NewHandler(s.cfg, s.ctx, s.l)
	// s.mux.HandleFunc("POST /user/register", h.RegisterUser)
	return c.Handler(s.mux)
}
