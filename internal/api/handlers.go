package api

import (
	"encoding/json"
	"fmt"
	"github.com/DmitryOdintsov/standartAPI_Server/internal/app/models"
	"github.com/go-chi/chi/v5"
	"io"
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
	user, err := api.storage.User().CreateUser(&users)
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
	initHeaders(w)
	api.logger.Info("Get All Users GET /users")
	users, err := api.storage.User().GetUsers()
	if err != nil {
		api.logger.Infof("ошибка при User.GetAll %s", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Возникли некоторые проблемы с доступом к базе данных. Попробуйте позже..",
			IsError:    true,
		}
		Json(w, msg, http.StatusNotImplemented)
		return
	}
	Json(w, users, http.StatusOK)
}

func (api *API) GetUserByID(w http.ResponseWriter, req *http.Request) {
	api.logger.Info("Get user by ID GET/user{id}")

	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	fmt.Println("-----------", id)
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

	Json(w, user, http.StatusOK)
}
func (api *API) DeleteUserByID(w http.ResponseWriter, req *http.Request) {
	api.logger.Info("Delete user by ID DELETE/users{id}")
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
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

	ok, err := api.storage.User().DeleteUser(id)
	if err != nil {
		api.logger.Info("Проблемы при удалении с базы данных (users) с id. err", err)
		msg := Message{
			StatusCode: http.StatusNotImplemented,
			Message:    "Возникли проблемы с удалением из базы данных, попробуйте еще..",
			IsError:    ok,
		}
		Json(w, msg, http.StatusNotImplemented)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	msg := Message{
		StatusCode: http.StatusAccepted,
		Message:    fmt.Sprintf("Пользователь с ID %d успешно удален", id),
		IsError:    ok,
	}
	Json(w, msg, http.StatusAccepted)

}

func (api *API) UpdateUserByID(w http.ResponseWriter, req *http.Request) {
	api.logger.Info("Update user by ID PUT/user{id}")
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
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
	fmt.Println("===================", string(content))
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

	str := api.storage.User().UpdateUserAge(user, id)
	Json(w, str, http.StatusOK)
}
