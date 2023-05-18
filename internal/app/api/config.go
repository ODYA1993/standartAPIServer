package api

import (
	"github.com/BurntSushi/toml"
	"github.com/DmitryOdintsov/standartAPI_Server/storage"
	"log"
)

type Config struct {
	Port        string `toml:"host"`
	LoggerLevel string `toml:"logger_level"`
	Storage     *storage.Config
}

func NewConfig(tomlPath string, port string) *Config {
	config := &Config{
		Port:        port,
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
	config.initToml(tomlPath)

	return config
}

func (conf *Config) initToml(tomlPath string) {
	//_, err := toml.DecodeFile(tomlPath, &conf.Storage)
	_, err := toml.DecodeFile(tomlPath, &conf)

	if err != nil {
		log.Fatal("не удалось десериализовать .toml файл", err)

	}
}
