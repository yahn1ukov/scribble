package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HTTP struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"http"`

	GRPC struct {
		Client struct {
			Notebook struct {
				Host string `yaml:"host"`
				Port int    `yaml:"port"`
			} `yaml:"notebook"`

			Note struct {
				Host string `yaml:"host"`
				Port int    `yaml:"port"`
			} `yaml:"note"`

			File struct {
				Host string `yaml:"host"`
				Port int    `yaml:"port"`
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
