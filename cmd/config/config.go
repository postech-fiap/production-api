package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func NewConfig() (*Config, error) {
	file, err := os.Open("cmd/resources/config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	config := &Config{}
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	config.Database.Host = os.Getenv(config.Database.Host)
	config.Database.Port = os.Getenv(config.Database.Port)
	config.Database.Username = os.Getenv(config.Database.Username)
	config.Database.Password = os.Getenv(config.Database.Password)

	return config, nil
}
