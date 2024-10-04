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

func ReadConfig(path string) (error, Config) {
	yamlData, err := os.ReadFile(path)

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
