package serverdatamanager

import (
	"github.com/LidorAlmkays/MineServerForge/dtos"
	"github.com/LidorAlmkays/MineServerForge/internal/application"
	"github.com/LidorAlmkays/MineServerForge/internal/infrastructure"
	"github.com/LidorAlmkays/MineServerForge/internal/model"
)

type baseServerConfigDataManager struct {
	mine infrastructure.MinecraftServerConfigStorage
}

func NewBaseServerConfigDataManager(mine infrastructure.MinecraftServerConfigStorage) application.ServerConfigDataManager {
	return &baseServerConfigDataManager{mine}
}

func (b *baseServerConfigDataManager) CreateServer(createMinecraftDto dtos.CreateMinecraftServerDTO) (int64, error) {
	model := model.NewMinecraftServerConfigModel(createMinecraftDto.OwnerEmail, createMinecraftDto.ServerName, createMinecraftDto.AllocatedRamMB, createMinecraftDto.MaxPlayerAmount, createMinecraftDto.ModesNames, createMinecraftDto.PluginsNames)
	id, err := b.mine.SaveNewServer(*model)
	if err != nil {
		return -1, err
	}
	return id, nil
}
