package config

import (
	"fmt"

	"github.com/flowshot-io/x/pkg/config"
)

const file = "commander-srv.yaml"

type Config struct{}

// LoadConfig Helper function for loading configuration
func LoadConfig(configDir string) (*Config, error) {
	conf := Config{}
	err := config.Load("", file, &conf)
	if err != nil {
		return nil, fmt.Errorf("config file corrupted: %w", err)
	}

	return &conf, nil
}

// Validate validates this config
func (c *Config) Validate() error {
	return nil
}
