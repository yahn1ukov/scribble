package config

type Config struct {
	GRPC struct {
		Network string `yaml:"network"`
		Host    string `yaml:"host"`
		Port    int    `yaml:"port"`
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

	Storage struct {
		MinIO struct {
			Bucket string `yaml:"bucket"`
		} `yaml:"minio"`
	} `yaml:"storage"`
}
