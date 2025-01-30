package config

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type DbConfig struct {
	DbName     string `env:"DB_NAME" validate:"required"`
	DbUser     string `env:"DB_USERNAME" validate:"required"`
	DbPassword string `env:"DB_PASSWORD" validate:"required"`
	DbHost     string `env:"DB_HOST" validate:"required"`
	DbPort     int    `env:"DB_PORT" validate:"required,min=1,max=65535"`
}

func (m DbConfig) ValidateSelf() error {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		err = errors.New("failed to validate MinecraftServersConfigsDbConfigs, error: " + err.Error())
		return err
	}
	return nil
}
