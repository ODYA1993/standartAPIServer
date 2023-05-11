package api

import "github.com/DmitryOdintsov/standartAPI_Server/storage"

type Config struct {
	Port        string `toml:"host"`
	LoggerLevel string `toml:"logger_level"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	return &Config{
		Port:        "8100",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
