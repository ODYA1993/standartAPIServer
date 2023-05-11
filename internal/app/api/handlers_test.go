package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func initApi() *API {
	config := NewConfig()
	return NewAPI(config)
}

func TestGetAllUsers(t *testing.T) {
	api := initApi()
	srv := httptest.NewServer(http.HandlerFunc(api.GetAllUsers))
	defer srv.Close()
	t.Error("ошибка")
}
