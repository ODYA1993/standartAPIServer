package api

import (
	"github.com/DmitryOdintsov/standartAPI_Server/pkg/storage"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"log"
	"sync"
)

type Config struct {
	PortAddr    string `env:"port_addr" env-default:":8085"`
	LoggerLevel string `env:"logger_level" env-default:"debug"`
	Storage     storage.Config
}

func NewConfig(folder, filename string) *Config {
	cfg := &Config{
		Storage: *storage.NewConfig(),
	}
	cfg = initConfigENV(folder, filename)
	return cfg
}

var instance Config
var once sync.Once

func initConfigENV(folder, filename string) *Config {
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
