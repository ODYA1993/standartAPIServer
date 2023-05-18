package api

import (
	"github.com/DmitryOdintsov/standartAPI_Server/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func NewAPI(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("Запуск APIServer на порту", api.config.Port)
	api.configureRouterField()
	if err := api.configureStorageField(); err != nil {
		return err
	}

	return http.ListenAndServe(api.config.Port, api.router)
}
