package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/LidorAlmkays/MineServerForge/config"
	"github.com/LidorAlmkays/MineServerForge/utils/logger"
)

type Handler struct {
	ctx    context.Context
	l 		logger.Logger
	cfg 	*config.Config
}

func NewHandler(cfg *config.Config, ctx context.Context, l logger.Logger) *Handler {
	return &Handler{ctx: ctx, l: l, cfg: cfg}
}

func (h *Handler) printReceivedMessageBody(r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Error(err)
	}
	bodyString := string(bodyBytes)
	h.l.Debug(bodyString)
}