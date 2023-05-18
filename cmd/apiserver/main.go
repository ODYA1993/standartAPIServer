package main

import (
	"flag"
	"github.com/DmitryOdintsov/standartAPI_Server/internal/app/api"
	"log"
)

var (
	ConfigPath string
	Port       *string
)

func init() {
	flag.StringVar(&ConfigPath, "path", "configs/conf.toml", "путь к конфигурационному файлу в .toml формате.")
	Port = flag.String("port", "8085", "port number")
}

func main() {
	flag.Parse()
	config := api.NewConfig(ConfigPath, *Port)
	server := api.NewAPI(config)
	err := server.Start()
	if err != nil {
		log.Println(err)
	}
}
