package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	GRPC struct {
		Server struct {
			Network string `yaml:"network"`
			Address string `yaml:"address"`
		} `yaml:"server"`
	} `yaml:"grpc"`

	DB struct {
		Postgres struct {
			URL string `yaml:"url"`
		} `yaml:"postgres"`
	} `yaml:"db"`
}

func New(path string) (*Config, error) {
	var cfg Config

	if path != "" {
		if err := cleanenv.ReadConfig(path, &cfg); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
