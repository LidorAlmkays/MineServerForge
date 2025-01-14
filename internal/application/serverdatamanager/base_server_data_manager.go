package serverdatamanager

import (
	"github.com/LidorAlmkays/MineServerForge/dtos"
	"github.com/LidorAlmkays/MineServerForge/internal/infrastructure/minecraftservers"
	"github.com/LidorAlmkays/MineServerForge/internal/models"
)

type baseServerDataManager struct {
	mine minecraftservers.MinecraftServers
}

func NewBaseServerDataManager(mine minecraftservers.MinecraftServers) ServerDataManager {
	return &baseServerDataManager{mine}
}

func (b *baseServerDataManager) CreateServer(createMinecraftDto dtos.CreateMinecraftServerDTO) error {
	model := models.NewMinecraftServerConfigModel(createMinecraftDto.OwnerEmail, createMinecraftDto.ServerName, createMinecraftDto.AllocatedRamMB, createMinecraftDto.MaxPlayerAmount, createMinecraftDto.ModesNames, createMinecraftDto.PluginsNames)
	err := b.mine.SaveNewServer(model)
	if err != nil {
		return err
	}
	return nil
}
