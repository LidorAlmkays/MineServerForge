package minecraftservers

import "github.com/LidorAlmkays/MineServerForge/internal/models"

type MinecraftServers interface {
	SaveNewServer(*models.MinecraftServerConfigModel) error
}
