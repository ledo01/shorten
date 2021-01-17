package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct for the app
type Config struct {
	Server struct {
		Port string `yaml:"port" envconfig:"SERVER_PORT"`
		Host string `yaml:"host" envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Database struct {
		URL string `yaml:"url" envconfig:"DB_URL"`
	} `yaml:"database"`
}

// NewYamlConfig create a Config from a Yaml file
func NewYamlConfig(path string) (*Config, error) {
	config := &Config{}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
