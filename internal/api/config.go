package api

import (
	"github.com/DmitryOdintsov/standartAPI_Server/storage"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"log"
	"sync"
)

type Config struct {
	PortAddr    string `toml:"port_addr" env:"port_addr" env-default:":8085"`
	LoggerLevel string `toml:"logger_level" env:"logger_level" env-default:"debug"`
	Storage     storage.Config
}

func NewConfig() *Config {
	cfg := &Config{
		Storage: *storage.NewConfig(),
	}
	cfg = initConfigENV()
	return cfg
}

var instance Config
var once sync.Once

func initConfigENV() *Config {
	once.Do(func() {
		logrus.Println("initializing config")
		instance = Config{}
	})
	if err := cleanenv.ReadEnv(&instance); err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	return &instance
}
