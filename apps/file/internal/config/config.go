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

	Storage struct {
		MinIO struct {
			AccessKey string `yaml:"accessKey"`
			SecretKey string `yaml:"secretKey"`
			Endpoint  string `yaml:"endpoint"`
			Bucket    string `yaml:"bucket"`
			Region    string `yaml:"region"`
			UseSSL    bool   `yaml:"useSSL"`
		} `yaml:"minio"`
	} `yaml:"storage"`
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
