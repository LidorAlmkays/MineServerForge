package serverdatamanager

import (
	"context"

	"github.com/LidorAlmkays/MineServerForge/dtos"
	"github.com/LidorAlmkays/MineServerForge/internal/application"
	"github.com/LidorAlmkays/MineServerForge/internal/infrastructure"
	"github.com/LidorAlmkays/MineServerForge/internal/model/db/minecraftserverconfig"
)

type baseServerConfigDataManager struct {
	mine infrastructure.MinecraftServerConfigStorage
}

func NewBaseServerConfigDataManager(mine infrastructure.MinecraftServerConfigStorage) application.ServerConfigDataManager {
	return &baseServerConfigDataManager{mine}
}

func (b *baseServerConfigDataManager) CreateServer(ctx context.Context, dto dtos.CreateMinecraftServerDTO) (int32, error) {
	model := minecraftserverconfig.CreateServerConfigParams{
		OwnerEmail:      dto.OwnerEmail,
		ServerName:      dto.ServerName,
		AllocatedRamMb:  int32(dto.AllocatedRamMB),
		MaxPlayerAmount: int32(dto.MaxPlayerAmount),
		ModesNames:      make([]string, 0),
		PluginsNames:    make([]string, 0),
	}
	// createMinecraftDto.OwnerEmail, createMinecraftDto.ServerName, createMinecraftDto.AllocatedRamMB, createMinecraftDto.MaxPlayerAmount, createMinecraftDto.ModesNames, createMinecraftDto.PluginsNames
	id, err := b.mine.SaveNewServer(ctx, model)
	if err != nil {
		return -1, err
	}
	return id, nil
}
