package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/LidorAlmkays/MineServerForge/config"
	"github.com/LidorAlmkays/MineServerForge/internal/application"
	"github.com/LidorAlmkays/MineServerForge/pkg/logger"
)

type Handler struct {
	ctx context.Context
	l   logger.Logger
	cfg *config.Config
	s   application.ServerConfigDataManager
}

func NewHandler(cfg *config.Config, ctx context.Context, l logger.Logger, s application.ServerConfigDataManager) *Handler {
	return &Handler{ctx, l, cfg, s}
}

func (h *Handler) printReceivedMessageBody(r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Error(err)
	}
	bodyString := string(bodyBytes)
	h.l.Debug(bodyString)
}
