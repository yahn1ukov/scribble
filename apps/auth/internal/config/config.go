package config

import "github.com/ilyakaznacheev/cleanenv"

const PATH = "configs/config.yaml"

type Config struct {
	GRPC struct {
		Server struct {
			Network string `yaml:"network"`
			Address string `yaml:"address"`
		} `yaml:"server"`

		Client struct {
			User struct {
				Address string `yaml:"address"`
			} `yaml:"user"`
		} `yaml:"client"`
	} `yaml:"grpc"`

	JWT struct {
		Secret string `yaml:"secret"`
		Expiry int    `yaml:"expiry"`
	} `yaml:"jwt"`
}

func New() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(PATH, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
