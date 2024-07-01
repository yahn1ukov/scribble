package config

type Config struct {
	GRPC struct {
		Network string `yaml:"network"`
		Host    string `yaml:"host"`
		Port    int    `yaml:"port"`
	} `yaml:"grpc"`

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
