package rest

import (
	"net/http"

	"github.com/LidorAlmkays/MineServerForge/config"
	"github.com/LidorAlmkays/MineServerForge/internal/api"
	_ "github.com/LidorAlmkays/MineServerForge/internal/api/docs"
	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"

	"github.com/LidorAlmkays/MineServerForge/internal/api/REST/handlers"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Server) addRoutes() (http.Handler, error) {
	s.l.Message("Setting up http routes")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})
	h := handlers.NewHandler(s.cfg, s.ctx, s.l, s.s)
	s.mux.Handle("POST /minecraft/create", http.HandlerFunc(h.CreateMinecraftServer))
	s.l.Message("minecraft/create endpoint ready")

	if config.Flags.Mode == enums.DevelopmentMode {
		s.mux.Handle("GET /swagger/", httpSwagger.Handler(
			httpSwagger.URL("http://localhost"+s.httpServer.Addr+"/swagger.yaml"),
			httpSwagger.DeepLinking(true),
			httpSwagger.DocExpansion("list"),
		))
		s.l.Message("Swagger api being served on /swagger/")

		api.ServeSwagger(s.mux)

		s.l.Message("Swagger files are being served on /swagger.yaml and /swagger.json")

	}
	return c.Handler(s.mux), nil
}
