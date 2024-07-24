package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type HostConfig struct {
	Host    string `yaml:"host"`
	User    string `yaml:"user"`
	KeyPath string `yaml:"keyPath"`
}

type TaskConfig struct {
	Category       string `yaml:"category"`
	Type           string `yaml:"type"`
	TaskParameters `yaml:"parameters"`
}

type TaskParameters map[string]interface{}

type Config struct {
	Host  HostConfig   `yaml:"host"`
	Facts FactConfig   `yaml:"facts"`
	Tasks []TaskConfig `yaml:"tasks"`
}

type FactConfig struct {
	Commands []string `yaml:"commands"`
}

func LoadConfig(configFile string) (*Config, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
