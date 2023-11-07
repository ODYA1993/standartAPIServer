package api

import (
	"github.com/DmitryOdintsov/standartAPI_Server/storage"
	"github.com/sirupsen/logrus"
)

func (a *API) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
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
