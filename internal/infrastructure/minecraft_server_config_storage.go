package infrastructure

import (
	"context"

	"github.com/LidorAlmkays/MineServerForge/internal/model/db/minecraftserverconfig"
)

type MinecraftServerConfigStorage interface {
	SaveNewServer(context.Context, minecraftserverconfig.CreateServerConfigParams) (int32, error)
	UpdateItem(context.Context, minecraftserverconfig.UpdateServerConfigParams) error
	InitTable() error
}
