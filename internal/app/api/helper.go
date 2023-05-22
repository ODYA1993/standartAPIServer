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

func (a *API) configureRouterField() {
	a.router.HandleFunc("/users", a.GetAllUsers).Methods("GET")
	a.router.HandleFunc("/users/{id}", a.GetUserByID).Methods("GET")
	a.router.HandleFunc("/users/{id}", a.DeleteUserByID).Methods("DELETE")
	a.router.HandleFunc("/user", a.PostUser).Methods("POST")
	a.router.HandleFunc("/users/{id}", a.UpdateUserAgeByID).Methods("PUT")
	a.router.HandleFunc("/friend", a.PostFriends).Methods("POST")
	a.router.HandleFunc("/user/{id}/friends", a.GetAllFriends).Methods("GET")
}

func (a *API) configureStorageField() error {
	store := storage.NewStorage(a.config.Storage)
	if err := store.Open(); err != nil {
		return err
	}
	a.storage = store
	return nil
}
