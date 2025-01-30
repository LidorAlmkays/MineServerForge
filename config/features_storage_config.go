package config

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type FeaturesStorageConfig struct {
	MinecraftModesStorage   string `env:"MINECRAFT_MODES_STORAGE" validate:"required"`
	MinecraftPluginsStorage string `env:"MINECRAFT_PLUGINS_STORAGE" validate:"required"`
}

func (f FeaturesStorageConfig) ValidateSelf() error {
	validate := validator.New()
	err := validate.Struct(f)
	if err != nil {
		err = errors.New("failed to validate FeaturesStorageConfig, error: " + err.Error())
		return err
	}
	return nil
}
