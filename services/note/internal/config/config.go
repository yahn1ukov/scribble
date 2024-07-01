package config

type Config struct {
	HTTP struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"http"`

	GRPC struct {
		Notebook struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"notebook"`

		Storage struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"storage"`

		File struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"file"`
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
}
