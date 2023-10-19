package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mdshack/shackelford.dev/internal/config/auth"
	"github.com/mdshack/shackelford.dev/internal/config/database"
	"github.com/mdshack/shackelford.dev/internal/config/services"
)

type Config struct {
	Auth     auth.Auth
	Services services.Services
	Database database.Database
}

func New() *Config {
	c := Config{
		Auth:     auth.Auth{},
		Services: services.Services{},
		Database: database.Database{},
	}

	envconfig.Process("auth", &c.Auth)
	envconfig.Process("services", &c.Services)
	envconfig.Process("database", &c.Database)

	return &c
}
