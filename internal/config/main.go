package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type FileEntry struct {
	Content string            `yaml:"content"`
	Headers map[string]string `yaml:"headers"`
}

type Config struct {
	Files map[string]FileEntry `yaml:"files"`
}

func ReadConfig() (error, Config) {
	yamlData, err := os.ReadFile("/etc/well-known/config.yaml")

	if err != nil {
		return err, Config{}
	}

	var config Config
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return err, Config{}
	}

	return nil, config
}
