package configs

type ConfigObject interface {
	ValidateSelf() error
}
