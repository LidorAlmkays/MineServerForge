package application

import (
	"context"

	"github.com/LidorAlmkays/MineServerForge/dtos"
)

type ServerConfigDataManager interface {
	CreateServer(context.Context, dtos.CreateMinecraftServerDTO) (id int32, err error)
}
