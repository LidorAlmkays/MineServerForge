package config

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ServiceConfig struct {
	HttpPort    int    `env:"HTTP_PORT" validate:"required,min=1,max=65535"` // Port must be within valid range
	GrpcPort    int    `env:"GRPC_PORT" validate:"required,min=1,max=65535"` // Port must be within valid range
	ProjectName string `env:"PROJECT_NAME" validate:"required"`              // Must not be empty
}

func (s ServiceConfig) ValidateSelf() error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		err = errors.New("failed to validate FeaturesStorageConfig, error: " + err.Error())
		return err
	}
	return nil
}
