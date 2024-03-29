package main

import (
	"github.com/DmitryOdintsov/standartAPI_Server/internal/api"
	"github.com/sirupsen/logrus"
)

func main() {
	config := api.NewConfig()
	server := api.NewAPI(config)
	err := server.Start()
	if err != nil {
		logrus.Println(err)
	}

}
