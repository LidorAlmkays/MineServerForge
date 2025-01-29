package db

import (
	"context"
	"database/sql"

	"github.com/LidorAlmkays/MineServerForge/internal/infrastructure"
	"github.com/LidorAlmkays/MineServerForge/internal/model/db/minecraftserverconfig"
)

type minecraftServerConfigRepo struct {
	queries *minecraftserverconfig.Queries
}

func NewMinecraftServerConfigRepo(db *sql.DB) infrastructure.MinecraftServerConfigStorage {
	return &minecraftServerConfigRepo{queries: minecraftserverconfig.New(db)}
}
func (mineRepo *minecraftServerConfigRepo) InitTable() error {
	err := mineRepo.queries.EnsureTableExists(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (mineRepo *minecraftServerConfigRepo) SaveNewServer(ctx context.Context, item minecraftserverconfig.CreateServerConfigParams) (int32, error) {
	return mineRepo.queries.CreateServerConfig(ctx, item)
}

func (mineRepo *minecraftServerConfigRepo) UpdateItem(ctx context.Context, item minecraftserverconfig.UpdateServerConfigParams) error {
	return mineRepo.queries.UpdateServerConfig(ctx, item)
}
