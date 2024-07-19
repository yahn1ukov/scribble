package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	GRPC struct {
		Server struct {
			Network string `yaml:"network"`
			Host    string `yaml:"host"`
			Port    int    `yaml:"port"`
		} `yaml:"server"`
	} `yaml:"grpc"`

	DB struct {
		Postgres struct {
			Driver   string `yaml:"driver"`
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Name     string `yaml:"name"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			SSLMode  string `yaml:"sslMode"`
		} `yaml:"postgres"`
	} `yaml:"db"`
}

func New() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig("configs/config.yaml", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
