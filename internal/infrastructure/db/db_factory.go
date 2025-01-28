package db

import (
	"fmt"

	"github.com/LidorAlmkays/MineServerForge/config"
	"github.com/LidorAlmkays/MineServerForge/internal/infrastructure"
	"github.com/LidorAlmkays/MineServerForge/pkg/logger"
	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"
	"github.com/jmoiron/sqlx"
)

var dbF *dbFactory

type dbFactory struct {
	dbConnections map[enums.DbType]*sqlx.DB
	l             logger.Logger
	lockManager   *dbConnectionLockManager
}

// TODO make standalone thread safe.
func GetDBFactory(l logger.Logger) *dbFactory {
	if dbF == nil {
		dbF = &dbFactory{l: l, dbConnections: make(map[enums.DbType]*sqlx.DB), lockManager: newDbConnectionLockManager()}
	}
	return dbF
}

func (f *dbFactory) GetMinecraftServer(l logger.Logger, t enums.DbType, cfg *config.DbConfig) (infrastructure.MinecraftServerConfigStorage, error) {
	var m infrastructure.MinecraftServerConfigStorage
	var err error
	switch t {
	case enums.Postgres:
		dbF.lockManager.Lock(enums.Postgres)
		if dbF.dbConnections[enums.Postgres] == nil {
			db, err := f.initializePostgresDB(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName, cfg.DbPort)
			if err != nil {
				return nil, err
			}
			dbF.dbConnections[enums.Postgres] = db
			f.l.Message("Connected to postgres database.")
		}
		dbF.lockManager.Unlock(enums.Postgres)

		m = NewMineServerConfRepoPostgres(dbF.dbConnections[t], l)

	default:
		m = nil
		err = fmt.Errorf("not supported database type selected for getting minecraft server.")
	}
	err = m.InitRepo()
	return m, err
}
