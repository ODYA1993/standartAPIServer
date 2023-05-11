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
	port       *string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "путь к конфигурационному файлу в .toml формате.")
	port = flag.String("port", "8089", "port number")
}

func main() {

	flag.Parse()
	fmt.Println("port", *port)
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, &config)
	_, err = toml.DecodeFile(configPath, &config.Storage)

	if err != nil {
		log.Println("не удалось десериализовать .toml файл", err)
	}
	config.Port = *port
	server := api.NewAPI(config)

	err = server.Start()
	if err != nil {
		log.Println(err)
	}
}
