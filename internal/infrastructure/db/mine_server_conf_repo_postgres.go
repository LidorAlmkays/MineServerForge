package db

import (
	"fmt"

	"github.com/LidorAlmkays/MineServerForge/internal/infrastructure"
	"github.com/LidorAlmkays/MineServerForge/internal/model"
	"github.com/LidorAlmkays/MineServerForge/pkg/logger"
	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"
	"github.com/LidorAlmkays/MineServerForge/queries"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type mineServerConfRepoPostgres struct {
	l              logger.Logger
	db             *sqlx.DB
	queryDirectory string
}

func NewMineServerConfRepoPostgres(db *sqlx.DB, l logger.Logger) infrastructure.MinecraftServerConfigStorage {
	return &mineServerConfRepoPostgres{db: db, l: l,
		queryDirectory: "minecraft_server_config/"}
}

func (repo *mineServerConfRepoPostgres) InitRepo() error {
	// Load the query to check if the table exists
	query, err := queries.GetQuery(repo.queryDirectory + "check_table_exists.sql")
	if err != nil {
		return fmt.Errorf("failed to get postgres check table exists query, error: %w", err)
	}

	// Execute the query to check if the table exists
	var exists bool
	err = dbF.dbConnections[enums.Postgres].Get(&exists, query)
	if err != nil {
		return fmt.Errorf("failed to check if table exists in postgres, error: %w", err)
	}

	// If the table already exists, do nothing
	if exists {
		repo.l.Info(fmt.Sprintf("Table minecraft server already exists in postgres, skipping creation."))
		return nil
	}

	// Load the query to create the table
	query, err = queries.GetQuery(repo.queryDirectory + "create_table.sql")
	if err != nil {
		return fmt.Errorf("failed to get create table query, error: %w", err)
	}

	// Execute the create table query
	repo.l.Info(fmt.Sprintf("Creating table minecraft server in postgres."))
	_, err = dbF.dbConnections[enums.Postgres].Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table minecraft server in postgres database, error: %w", err)
	}

	repo.l.Info(fmt.Sprintf("Table minecraft server successfully created in postgres."))
	return nil

}

func (repo *mineServerConfRepoPostgres) SaveNewServer(item model.MinecraftServerConfigModel) (int64, error) {
	query, err := queries.GetQuery(repo.queryDirectory + "insert_server.sql")
	if err != nil {
		return -1, err
	}

	// Map struct fields manually and use pq.Array for slices
	args := map[string]interface{}{
		"owner_email":       item.OwnerEmail,
		"server_name":       item.ServerName,
		"allocated_ram_mb":  item.AllocatedRamMB,
		"max_player_amount": item.MaxPlayerAmount,
		"modes_names":       pq.Array(item.ModesNames),
		"plugins_names":     pq.Array(item.PluginsNames),
	}

	result, err := repo.db.NamedExec(query, args)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId() // If your DB supports this
	if err != nil {
		return -1, err
	}

	return id, err
}

func (repo *mineServerConfRepoPostgres) UpdateItem(item model.MinecraftServerConfigModel) error {
	query, err := queries.GetQuery(repo.queryDirectory + "update_server.sql")
	if err != nil {
		return err
	}
	_, err = repo.db.NamedExec(query, item)
	return err
}

func (repo *mineServerConfRepoPostgres) Shutdown() error {
	return repo.db.Close()
}
