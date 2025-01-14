package config

import (
	"github.com/LidorAlmkays/MineServerForge/pkg/configs"
	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"
)

type Config struct {
	ServiceConfig *ServiceConfig `validate:"required"` // Ensures ServiceConfig is not nil
	StorageConfig *StorageConfig `validate:"required"` // Ensures StorageConfig is not nil
}

type ServiceConfig struct {
	HttpPort    int    `env:"HTTP_PORT" validate:"required,min=1,max=65535"` // Port must be within valid range
	GrpcPort    int    `env:"GRPC_PORT" validate:"required,min=1,max=65535"` // Port must be within valid range
	ProjectName string `env:"PROJECT_NAME" validate:"required"`              // Must not be empty
}

type StorageConfig struct {
	MinecraftModesStorage   *string `env:"MINECRAFT_MODES_STORAGE"`
	MinecraftPluginsStorage *string `env:"MINECRAFT_PLUGINS_STORAGE"`
	MinecraftConfigStorage  *string `env:"MINECRAFT_CONFIG_STORAGE"`
}

func (s *Config) SetUp(configPath string, configType enums.ConfigTypes) error {
	var err error
	s.ServiceConfig, err = configs.GetConfig[ServiceConfig](configPath, configType)
	if err != nil {
		return err
	}
	s.StorageConfig, err = configs.GetConfig[StorageConfig](configPath, configType)
	if err != nil {
		return err
	}
	return nil

}
