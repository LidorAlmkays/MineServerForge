package minecraftservers

import (
	"fmt"

	"github.com/LidorAlmkays/MineServerForge/internal/models"
)

type minecraftServersPostgres struct {
	dsn string
}

func NewMinecraftServerPostgres(dbUser, dbPassword, dbHost string, dbPort int) MinecraftServers {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/postgres", dbUser, dbPassword, dbHost, dbPort)
	return &minecraftServersPostgres{dsn}
}

// SaveConfig implements MinecraftServers.
func (m *minecraftServersPostgres) SaveNewServer(*models.MinecraftServerConfigModel) error {

	return nil
}
