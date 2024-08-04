package config

import "github.com/ilyakaznacheev/cleanenv"

const PATH = "configs/config.yaml"

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

func New() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(PATH, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
