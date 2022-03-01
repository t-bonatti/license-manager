package config

import (
	"github.com/apex/log"
	"github.com/caarlos0/env"
)

// Config struct of environments
type Config struct {
	Port        string `env:"PORT" envDefault:"3000"`
	DatabaseDSN string `env:"DATABASE_DSN" envDefault:"host=localhost port=5432 user=postgres dbname=license sslmode=disable password="`
}

// Get config
func Get() (cfg Config) {
	if err := env.Parse(&cfg); err != nil {
		log.WithError(err).Fatal("failed to load config")
	}
	return
}
