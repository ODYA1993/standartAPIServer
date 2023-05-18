package api

import (
	"encoding/json"
	"fmt"
	"github.com/DmitryOdintsov/standartAPI_Server/internal/app/models"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func (api *API) PostUser(w http.ResponseWriter, req *http.Request) {
	api.logger.Info("Post User POST /user")
	var users models.User
	err := json.NewDecoder(req.Body).Decode(&users)
	if err != nil {
		api.logger.Info("Недопустимый json, полученный от клиента")
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Предоставленный json недействителен",
			IsError:    true,
		}
		Json(w, msg, http.StatusBadRequest)
		return
	}
	user, err := api.storage.User().CreateUSer(&users)
	if err != nil {
		api.logger.Info("Проблемы при создании нового пользователя")
		msg := Message{
			StatusCode: http.StatusNotImplemented,
			Message:    "Возникли проблемы с доступом к базе данных. Пробовать снова..",
			IsError:    true,
		}
		Json(w, msg, http.StatusNotImplemented)
		return
	}
	Json(w, user, http.StatusCreated)
}

func (api *API) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	api.logger.Info("Get All User GET /users")
	user, err := api.storage.User().GetUsers()
	if err != nil {
		api.logger.Info("ошибка при User.GetAll", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Возникли некоторые проблемы с доступом к базе данных. Попробуйте позже..",
			IsError:    true,
		}
		Json(w, msg, http.StatusNotImplemented)
		return
	}
	Json(w, user, http.StatusOK)
}

func (api *API) GetUserByID(w http.ResponseWriter, req *http.Request) {
	api.logger.Info("Get user by ID GET/users{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Проблемы при разборе параметра id")
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Неподходящее значение идентификатора",
			IsError:    true,
		}
		Json(w, msg, http.StatusBadRequest)
		return
	}
	user, ok, err := api.storage.User().GetUserByID(id)
	if err != nil {
		api.logger.Info("Проблемы при доступе к таблице базы данных (users) с id:", id, err)
		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "Возникли проблемы с доступом к базе данных, попробуйте еще..",
			IsError:    true,
		}
		Json(w, msg, http.StatusInternalServerError)
		return
	}
	if !ok {
		api.logger.Info("Не удается найти user с таким id в базе данных")
		msg := Message{
			StatusCode: http.StatusNotFound,
			Message:    "<user> с таким ID не существует в базе данных",
			IsError:    true,
		}
		Json(w, msg, http.StatusNotFound)
		return
	}
	Json(w, user, http.StatusOK)
}
func (api *API) DeleteUserByID(w http.ResponseWriter, req *http.Request) {
	api.logger.Info("Delete user by ID DELETE/users{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Проблемы при разборе параметра id")
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Неподходящее значение идентификатора",
			IsError:    true,
		}
		Json(w, msg, http.StatusBadRequest)
		return
	}
	_, ok, err := api.storage.User().GetUserByID(id)
	if err != nil {
		api.logger.Info("Проблемы при доступе к таблице базы данных (users) с id. err", err)
		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "Возникли проблемы с доступом к базе данных, попробуйте еще..",
			IsError:    true,
		}
		Json(w, msg, http.StatusInternalServerError)
		return
	}
	if !ok {
		api.logger.Info("Не удается найти user с таким id в базе данных")
		msg := Message{
			StatusCode: http.StatusNotFound,
			Message:    "<user> с таким ID не существует в базе данных",
			IsError:    true,
		}
		Json(w, msg, http.StatusNotFound)
		return
	}
	_, err = api.storage.User().DeleteUser(id)
	if err != nil {
		api.logger.Info("Проблемы при удалении с базы данных (users) с id. err", err)
		msg := Message{
			StatusCode: http.StatusNotImplemented,
			Message:    "Возникли проблемы с удалением из базы данных, попробуйте еще..",
			IsError:    true,
		}
		Json(w, msg, http.StatusNotImplemented)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	msg := Message{
		StatusCode: http.StatusAccepted,
		Message:    fmt.Sprintf("Пользователь с ID %d успешно удален", id),
		IsError:    false,
	}
	//json.NewEncoder(rw).Encode(msg)
	Json(w, msg, http.StatusAccepted)

}

func (api *API) UpdateUserAgeByID(w http.ResponseWriter, req *http.Request) {
	api.logger.Info("Update userAge by ID PUT/users{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Проблемы при разборе параметра id")
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Неподходящее значение идентификатора",
			IsError:    true,
		}
		Json(w, msg, http.StatusBadRequest)
		return
	}
	var user *models.User

	content, err := io.ReadAll(req.Body)
	if err != nil {
		api.logger.Info("Проблемы при считывании из body")
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Проблемы при считывании из body",
			IsError:    true,
		}
		Json(w, msg, http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(content, &user)
	if err != nil {
		api.logger.Info("Недопустимый json, полученный от клиента")
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Предоставленный json недействителен",
			IsError:    true,
		}
		Json(w, msg, http.StatusBadRequest)
		return
	}

	str := api.storage.User().UpdateUserAge(user.Age, id)
	Json(w, str, http.StatusOK)

}

func (api *API) PostFriends(w http.ResponseWriter, req *http.Request) {
	api.logger.Info("Post Friend POST /friends")
	initHeaders(w)
	defer req.Body.Close()

	content, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	var friendInput models.Friends
	err = json.Unmarshal(content, &friendInput)
	if err != nil {
		api.logger.Info("Недопустимый json, полученный от клиента", err)
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Предоставленный json недействителен",
			IsError:    true,
		}
		Json(w, msg, http.StatusBadRequest)
		return
	}
	friend, err := api.storage.Friends().CreateFriends(&friendInput)
	if err != nil {
		api.logger.Info("Проблемы при создании друга", err)
		msg := Message{
			StatusCode: http.StatusNotImplemented,
			Message:    "Возникли проблемы с доступом к базе данных. Попробуйте снова..",
			IsError:    true,
		}
		Json(w, msg, http.StatusNotImplemented)
		return
	}

	us, ok, err := api.storage.User().GetUserByID(friend.UserID)
	if err != nil {
		log.Println(err)
		return
	}
	if !ok {
		log.Println("Нету такого user")
		return
	}
	fr, ok, err := api.storage.User().GetUserByID(friend.FriendID)
	if err != nil {
		log.Println(err)
		return
	}
	if !ok {
		log.Println("Нету такого user")
		return
	}
	Json(w, fmt.Sprintf("%v и %v теперь друзья", us.Name, fr.Name), http.StatusCreated)
}

func (api *API) GetAllFriends(w http.ResponseWriter, req *http.Request) {
	initHeaders(w)
	api.logger.Info("Get All Friends GET /user/id/friends")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Проблемы при разборе параметра id")
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Неподходящее значение идентификатора",
			IsError:    true,
		}
		Json(w, msg, http.StatusBadRequest)
		return
	}

	friendID, err := api.storage.Friends().GetFriends(id)
	if err != nil {
		api.logger.Info("Ошибка при Friends.GetAll", err)
		msg := Message{
			StatusCode: 501,
			Message:    "У нас возникли некоторые проблемы с доступом к базе данных. Попробуйте позже..",
			IsError:    true,
		}
		Json(w, msg, http.StatusNotImplemented)
		return
	}
	users, err := api.storage.User().GetUsers()
	if err != nil {
		log.Println(err)
	}

	for _, v := range friendID {
		for _, k := range users {
			if v.ID == k.ID {
				Json(w, k, http.StatusOK)

			}
		}
	}
}
