package infrastructure

import "github.com/LidorAlmkays/MineServerForge/internal/model"

type MinecraftServerConfigStorage interface {
	SaveNewServer(model.MinecraftServerConfigModel) (int64, error)
	UpdateItem(item model.MinecraftServerConfigModel) error
	Shutdown() error
	InitRepo() error
}
