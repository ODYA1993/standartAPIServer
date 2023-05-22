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
	t.Log(respUserID)
}

func TestApiGetAllUsers(t *testing.T) {
	api := initApi()
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	t.Log(req.URL)
	api.router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "ошибка!")
}

func TestApiGetUserByID(t *testing.T) {
	api := initApi()
	createdUserID := createUser(t)

	resp := httptest.NewRecorder()
	url := fmt.Sprint("/users/", createdUserID)
	t.Log(url)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	api.router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "ошибка!!!")
}

func TestApiDeleteUserByID(t *testing.T) {
	api := initApi()
	createdUserID := createUser(t)
	resp := httptest.NewRecorder()

	url := fmt.Sprint("/users/", createdUserID)
	t.Log(url)

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
	t.Log(url)

	req := httptest.NewRequest(http.MethodPut, url, bytes.NewBuffer(testDataByte))
	api.router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "ошибка!")
}

func TestApiPostFriends(t *testing.T) {
	userID := createFriend(t)
	t.Log(userID)
}

func TestApiGetAllFriends(t *testing.T) {
	api := initApi()
	userID := createFriend(t)

	resp := httptest.NewRecorder()

	url := fmt.Sprint("/user/", userID, "/friends")
	t.Log(url)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	api.router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "ошибка!!!")
}

// Вспомогательные функции....
func createUser(t *testing.T) int {
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
	assert.Equal(t, http.StatusCreated, resp.Code, "ошибка code!")
	assert.NotEqual(t, testUser.Id, respUser.Id, "ошибка id!")
	assert.Equal(t, testUser.Name, respUser.Name, "ошибка Name!")
	assert.Equal(t, testUser.Age, respUser.Age, "ошибка Age!")

	return respUser.Id
}

func createFriend(t *testing.T) int {
	api := initApi()
	firstCreatedUserID := createUser(t)
	SecondCreatedUserID := createUser(t)

	type friend struct {
		Id       int `json:"id"`
		UserID   int `json:"user_id"`
		FriendID int `json:"friend_id"`
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
	t.Log(req.URL)
	api.router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code, "ошибка!")
	return firstCreatedUserID
}
