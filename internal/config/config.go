package config

import (
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	postgresDSN        string
	migrationsDir      string
	httpServerHostPort string
}

func (c *Config) GetPostgresDSN() string {
	return c.postgresDSN
}

func (c *Config) GetMigrationsDir() string {
	return c.migrationsDir
}

func (c *Config) GetHTTPServerHostPort() string {
	return c.httpServerHostPort
}

func InitConfig() (*Config, error) {
	postgresDSN := os.Getenv("POSTGRES_DSN")
	if postgresDSN == "" {
		return nil, errors.New("POSTGRES_DSN environment variable not set")
	}
	migrationsDir := os.Getenv("MIGRATIONS_DIR")
	if migrationsDir == "" {
		return nil, errors.New("MIGRATIONS_DIR environment variable not set")
	}
	httpServerHostPort := os.Getenv("HTTP_SERVER_HOST_PORT")
	if httpServerHostPort == "" {
		return nil, errors.New("HTTP_SERVER_HOST_PORT environment variable not set")
	}
	return &Config{
		postgresDSN:        postgresDSN,
		migrationsDir:      migrationsDir,
		httpServerHostPort: httpServerHostPort,
	}, nil
}
