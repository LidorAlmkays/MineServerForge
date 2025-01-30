package config

import (
	"strings"

	"github.com/LidorAlmkays/MineServerForge/pkg/configs"
	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"
)

type Config struct {
	FeaturesStorageConfig *FeaturesStorageConfig
	DbConfig              *DbConfig
	ServiceConfig         *ServiceConfig
}

func (s *Config) SetUp(configPath string, configType enums.ConfigTypes, debugMode bool) error {
	var err error

	s.ServiceConfig, err = configs.GetConfig[ServiceConfig](configPath, configType, debugMode)
	if err != nil {
		return err
	}
	s.FeaturesStorageConfig, err = configs.GetConfig[FeaturesStorageConfig](configPath, configType, debugMode)
	if err != nil {
		return err
	}
	s.DbConfig, err = configs.GetConfig[DbConfig](configPath, configType, debugMode)
	if err != nil {
		return err
	}
	s.DbConfig.DbName = strings.ToLower(s.DbConfig.DbName)

	return nil
}
