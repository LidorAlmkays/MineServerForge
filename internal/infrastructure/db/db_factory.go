package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/LidorAlmkays/MineServerForge/config"
	"github.com/LidorAlmkays/MineServerForge/internal/infrastructure"
	"github.com/LidorAlmkays/MineServerForge/internal/infrastructure/db/postgres"

	"github.com/LidorAlmkays/MineServerForge/pkg/logger"
	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"
)

var dbF *dbFactory

type dbFactory struct {
	l            logger.Logger
	mutexes      map[enums.DbType]*sync.Mutex
	dbConnection *sql.DB    //holds the db connection
	mu           sync.Mutex // protects access to the mutexes map
}

// TODO make standalone thread safe.
func GetDBFactory(l logger.Logger) *dbFactory {
	if dbF == nil {
		dbF = &dbFactory{l: l, mutexes: map[enums.DbType]*sync.Mutex{}}
	}
	return dbF
}

func (f *dbFactory) GetMinecraftServer(l logger.Logger, t enums.DbType, cfg *config.DbConfig) (infrastructure.MinecraftServerConfigStorage, error) {
	var m infrastructure.MinecraftServerConfigStorage
	var err error
	dbF.lock(enums.Postgres)
	defer dbF.unlock(enums.Postgres)

	if dbF.dbConnection == nil {
		switch t {
		case enums.Postgres:
			db, err := postgres.InitializeDB(f.l, cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName, cfg.DbPort)
			if err != nil {
				return nil, err
			}
			dbF.dbConnection = db
			f.l.Message("Connected to postgres database.")
		default:
			return nil, fmt.Errorf("not supported database type selected for getting minecraft server.")
		}
	}
	m = NewMinecraftServerConfigRepo(dbF.dbConnection)
	if err = m.InitTable(); err != nil {
		return nil, err
	}
	return m, err
}

// Lock locks the mutex for the given key.
func (f *dbFactory) lock(key enums.DbType) {
	f.mu.Lock()
	if _, exists := f.mutexes[key]; !exists {
		f.mutexes[key] = &sync.Mutex{}
	}
	mutex := f.mutexes[key]
	f.mu.Unlock()

	mutex.Lock()
}

// Unlock unlocks the mutex for the given key.
func (f *dbFactory) unlock(key enums.DbType) {
	f.mu.Lock()
	mutex, exists := f.mutexes[key]
	f.mu.Unlock()

	if exists {
		mutex.Unlock()
	}
}

func (f *dbFactory) Shutdown() error {
	return f.dbConnection.Close()
}
