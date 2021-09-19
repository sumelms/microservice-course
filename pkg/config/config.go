package config

import (
	"os"

	"github.com/pkg/errors"
	"github.com/sherifabdlnaby/configuro"
)

// Config struct
type Config struct {
	Server struct {
		HTTP *Server `validate:"required"`
	} `validate:"required"`
	Database *Database `validate:"required"`
}

// Database config struct
type Database struct {
	Driver   string `validate:"required"`
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
}

// Server config struct
type Server struct {
	Host string `validate:"required"`
}

// NewConfig creates a new configurator
func NewConfig(configPath string) (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.Wrapf(err, "Config file does not exists in %s", configPath)
	}

	loader, err := configuro.NewConfig(
		configuro.WithLoadFromConfigFile(configPath, false),
		configuro.WithLoadFromEnvVars("SUMELMS"))
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := loader.Load(config); err != nil {
		return nil, err
	}

	if err := loader.Validate(config); err != nil {
		return nil, err
	}

	return config, nil
}
