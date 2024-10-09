package config

import "github.com/ilyakaznacheev/cleanenv"

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

func New(path string) (*Config, error) {
	var cfg Config

	if path != "" {
		if err := cleanenv.ReadConfig(path, &cfg); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
