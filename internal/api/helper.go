package api

import (
	"github.com/DmitryOdintsov/standartAPI_Server/pkg/storage"
	"github.com/sirupsen/logrus"
	"os"
)

func (a *API) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	a.logger.SetLevel(logLevel)
	return nil
}

func (a *API) configureStorageField() error {
	store := storage.NewStorage(&a.config.Storage)
	if err := store.Open(); err != nil {
		return err
	}
	a.storage = *store
	return nil
}
