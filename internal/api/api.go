package api

import (
	"github.com/DmitryOdintsov/standartAPI_Server/pkg/storage"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  chi.Router
	storage storage.Storage
}

func NewAPI(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: chi.NewRouter(),
	}
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	logrus.Printf("Запуск APIServer на порту %s", api.config.PortAddr)
	api.configureRouterField()
	if err := api.configureStorageField(); err != nil {
		return err
	}
	return http.ListenAndServe(api.config.PortAddr, api.router)
}
