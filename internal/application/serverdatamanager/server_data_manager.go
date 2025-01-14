package serverdatamanager

import "github.com/LidorAlmkays/MineServerForge/dtos"

type ServerDataManager interface {
	CreateServer(dtos.CreateMinecraftServerDTO) error
}
