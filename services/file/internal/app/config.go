package app

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/yahn1ukov/scribble/services/file/internal/config"
)

func NewConfig() (*config.Config, error) {
	var cfg config.Config

	if err := cleanenv.ReadConfig("configs/config.yaml", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
