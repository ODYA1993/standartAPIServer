package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	configPath = "../../../configs/conf_test.toml"
	port       = "8090"
	id         = 1
)

func initApi() *API {
	config := NewConfig(configPath, port)
	api := NewAPI(config)
	api.configureRouterField()
	_ = api.configureStorageField()
	return api
}

func TestApiPostUser(t *testing.T) {
	respUserID := createUser(t)
	fmt.Println(respUserID)
}

func TestApiGetAllUsers(t *testing.T) {
	api := initApi()
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	api.router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "ошибка!")
}

func TestApiGetUserByID(t *testing.T) {
	api := initApi()
	createdUserID := createUser(t)

	resp := httptest.NewRecorder()
	url := fmt.Sprint("/users/", createdUserID)
	fmt.Println(url)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	api.router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "ошибка!!!")
}

func TestApiDeleteUserByID(t *testing.T) {
	api := initApi()
	createdUserID := createUser(t)
	resp := httptest.NewRecorder()

	url := fmt.Sprint("/users/", createdUserID)
	fmt.Println(url)

	req := httptest.NewRequest(http.MethodDelete, url, nil)
	api.router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusAccepted, resp.Code, "ошибка!")
}

func TestApiUpdateUserAgeByID(t *testing.T) {
	api := initApi()
	createdUserID := createUser(t)

	type user struct {
		Id  int
		Age int
	}

	testData := user{
		Age: 88888,
	}

	testDataByte, err := json.Marshal(testData)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	resp := httptest.NewRecorder()

	url := fmt.Sprint("/users/", createdUserID)
	fmt.Println(url)

	req := httptest.NewRequest(http.MethodPut, url, bytes.NewBuffer(testDataByte))
	api.router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "ошибка!")
}

func TestApiPostFriends(t *testing.T) {
	api := initApi()
	firstCreatedUserID := createUser(t)
	SecondCreatedUserID := createUser(t)

	type friend struct {
		Id       int
		UserID   any
		FriendID any
	}

	testData := friend{
		UserID:   firstCreatedUserID,
		FriendID: SecondCreatedUserID,
	}

	testDataByte, err := json.Marshal(testData)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/friend", bytes.NewBuffer(testDataByte))
	api.router.ServeHTTP(resp, req)

	var respFriend friend
	err = json.NewDecoder(resp.Body).Decode(&respFriend)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	assert.Equal(t, http.StatusCreated, resp.Code, "ошибка!")
	assert.NotEqual(t, testData.Id, respFriend.Id, "ошибка!")
	assert.Equal(t, testData.UserID, respFriend.UserID, "ошибка!")
	assert.Equal(t, testData.FriendID, respFriend.FriendID, "ошибка!")
}

func createUser(t *testing.T) any {
	api := initApi()
	type user struct {
		Id   int
		Name string
		Age  int
	}

	testUser := user{
		Name: "dima",
		Age:  30,
	}
	testDataByte, err := json.Marshal(testUser)
	if err != nil {
		log.Println(err)
	}
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(testDataByte))
	api.router.ServeHTTP(resp, req)

	var respUser user
	err = json.NewDecoder(resp.Body).Decode(&respUser)
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, http.StatusCreated, resp.Code, "ошибка!")
	assert.NotEqual(t, testUser.Id, respUser.Id, "ошибка!")
	assert.Equal(t, testUser.Name, respUser.Name, "ошибка!")
	assert.Equal(t, testUser.Age, respUser.Age, "ошибка!")

	return respUser.Id
}
