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
		Host     string
		Port     string
		Username string
		Password string
	}

	RabbitMQ struct {
		Host     string
		Port     string
		Username string
		Password string
	}
}

func NewConfig() (*Config, error) {
	file, err := os.Open("resources/config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	config := &Config{}
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	config.Database.Host = os.Getenv("MONGO_HOST")
	config.Database.Port = os.Getenv("MONGO_PORT")
	config.Database.Username = os.Getenv("MONGO_USERNAME")
	config.Database.Password = os.Getenv("MONGO_PASSWORD")

	config.RabbitMQ.Host = os.Getenv("RABBITMQ_HOST")
	config.RabbitMQ.Port = os.Getenv("RABBITMQ_PORT")
	config.RabbitMQ.Username = os.Getenv("RABBITMQ_USERNAME")
	config.RabbitMQ.Password = os.Getenv("RABBITMQ_PASSWORD")

	return config, nil
}
