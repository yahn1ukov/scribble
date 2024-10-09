package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HTTP struct {
		Address string `yaml:"address"`
	} `yaml:"http"`

	GRPC struct {
		Client struct {
			User struct {
				Address string `yaml:"address"`
			} `yaml:"user"`

			Notebook struct {
				Address string `yaml:"address"`
			} `yaml:"notebook"`

			Note struct {
				Address string `yaml:"address"`
			} `yaml:"note"`

			File struct {
				Address string `yaml:"address"`
			} `yaml:"file"`

			Auth struct {
				Address string `yaml:"address"`
			} `yaml:"auth"`
		} `yaml:"client"`
	} `yaml:"grpc"`

	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
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
