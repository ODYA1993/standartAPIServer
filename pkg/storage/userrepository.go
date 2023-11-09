package storage

import (
	"fmt"
	"github.com/DmitryOdintsov/standartAPI_Server/internal/models"
	"github.com/sirupsen/logrus"
	"log"
)

type UserRepository struct {
	storage *Storage
	logger  *logrus.Logger
}

var (
	tableUser = "users"
)

func (ur *UserRepository) CreateUser(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (firstName, lastName, username) VALUES($1, $2, $3)", tableUser)

	_, err := ur.storage.db.Exec(query, u.Name.FirstName, u.Name.LastName, u.Username)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *UserRepository) GetUsers() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT id, firstName, lastName, username FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name.FirstName, &user.Name.LastName, &user.Username)
		if err != nil {
			ur.logger.Println(err)
			continue
		}
		users = append(users, &user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, bool, error) {
	query := fmt.Sprintf("SELECT id, firstName, lastName, username FROM %s WHERE id = $1", tableUser)

	var u models.User
	if err := ur.storage.db.QueryRow(query, id).Scan(&u.ID, &u.Name.FirstName, &u.Name.LastName, &u.Username); err != nil {
		return nil, false, err
	}
	return &u, true, nil

}

func (ur *UserRepository) DeleteUser(id int) (bool, error) {
	_, ok, err := ur.GetUserByID(id)
	if err != nil {
		return false, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableUser)
		_, err = ur.storage.db.Exec(query, id)
		return false, err
	}
	return true, err
}

func (ur *UserRepository) DeleteAllUsers() (bool, error) {
	query := fmt.Sprintf("TRUNCATE TABLE %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return true, err
	}
	defer rows.Close()
	return false, nil
}

func (ur *UserRepository) UpdateUser(user *models.User, id int) string {
	query := fmt.Sprintf("UPDATE %s SET firstName=$1, lastName=$2, username=$3 WHERE id=%d", tableUser, id)
	_, err := ur.storage.db.Exec(query, user.Name.FirstName, user.Name.LastName, user.Username)

	if err != nil {
		log.Println(err)
	}
	return "пользователь обновлен"
}
