package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Engine struct {
		Type string `yaml:"type"`
	} `yaml:"engine"`
	Network struct {
		Address        string `yaml:"address"`
		MaxConnections int    `yaml:"max_connections"`
	} `yaml:"network"`
	Logging struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
		Output string `yaml:"output"`
	} `yaml:"logging"`
}

func NewConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
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
