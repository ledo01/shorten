package main

import (
	"github.com/kelseyhightower/envconfig"
)

// Config struct for the app
type Config struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT"`
		Host string `envconfig:"SERVER_HOST"`
	}
	Database struct {
		URL string `envconfig:"DB_URL"`
	}
}

// NewConfig create a Config from env variables
func NewConfig() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
