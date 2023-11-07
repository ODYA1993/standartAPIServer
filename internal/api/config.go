package api

import (
	"github.com/BurntSushi/toml"
	"github.com/DmitryOdintsov/standartAPI_Server/storage"
	"log"
)

type Config struct {
	Port        string `toml:"port_addr" env:"port_addr" env-default:":8085"`
	LoggerLevel string `toml:"logger_level" env:"logger_level" env-default:"info"`
	Storage     *storage.Config
}

func NewConfig(path string, port string) *Config {
	cfg := &Config{
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
	cfg.initToml(path)
	//cfg = initConfigENV(path)

	return cfg
}

func (conf *Config) initToml(tomlPath string) {
	//_, err := toml.DecodeFile(tomlPath, &conf.Storage)
	_, err := toml.DecodeFile(tomlPath, &conf)

	if err != nil {
		log.Fatal("не удалось десериализовать .toml файл", err)

	}
}

//var instance *Config
//
//func initConfigENV(path string) *Config {
//	if err := cleanenv.ReadConfig(path, &instance); err != nil {
//		log.Fatal(err)
//	}
//	return instance
//}
