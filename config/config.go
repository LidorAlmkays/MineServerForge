package config

type Config struct {
	ServiceConfig *ServiceConfig
}

type ServiceConfig struct {
	Port        int
	ProjectName string
}
