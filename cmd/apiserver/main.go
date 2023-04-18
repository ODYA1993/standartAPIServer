package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/DmitryOdintsov/standartAPI_Server/internal/app/api"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "путь к конфигурационному файлу в .toml формате.")
}

func main() {
	flag.Parse()

	fmt.Println("It works")
	config := api.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("не удалось десериализовать .toml файл", err)
	}

	server := api.NewAPI(config)
	if err = server.Start(); err != nil {
		log.Fatal(err)
	}
}
