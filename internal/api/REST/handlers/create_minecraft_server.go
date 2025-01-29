package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LidorAlmkays/MineServerForge/dtos"
)

// CreateMinecraftServer godoc
// @Summary Create a new Minecraft server
// @Description This endpoint allows the creation of a new Minecraft server with specified configurations.
// @Tags Minecraft
// @Accept  json
// @Produce  json
// @Param body body dtos.CreateMinecraftServerDTO true "Minecraft Server Creation Request"
// @Success 200 {object} dtos.CreateMinecraftServerResponseDTO "Successfully created a Minecraft server"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Router /minecraft/create [post]
func (h *Handler) CreateMinecraftServer(w http.ResponseWriter, r *http.Request) {
	h.l.Info("Received a request to create a minecraft server")
	var req dtos.CreateMinecraftServerDTO
	// Decode the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	id, err := h.s.CreateServer(r.Context(), req)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}
	res := &dtos.CreateMinecraftServerResponseDTO{
		Id: id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
