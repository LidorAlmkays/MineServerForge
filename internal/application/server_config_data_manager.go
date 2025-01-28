package application

import "github.com/LidorAlmkays/MineServerForge/dtos"

type ServerConfigDataManager interface {
	CreateServer(dtos.CreateMinecraftServerDTO) (id int64, err error)
}
