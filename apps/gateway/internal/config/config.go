package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HTTP struct {
		Address string `yaml:"address"`
	} `yaml:"http"`

	GRPC struct {
		Client struct {
			Notebook struct {
				Address string `yaml:"address"`
			} `yaml:"notebook"`

			Note struct {
				Address string `yaml:"address"`
			} `yaml:"note"`

			File struct {
				Address string `yaml:"address"`
			} `yaml:"file"`
		} `yaml:"client"`
	} `yaml:"grpc"`
}

func New() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig("configs/config.yaml", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
