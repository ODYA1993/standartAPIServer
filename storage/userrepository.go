package storage

import (
	"fmt"
	"github.com/DmitryOdintsov/standartAPI_Server/internal/app/models"
	"log"
)

type UserRepository struct {
	storage *Storage
}

var (
	tableUser = "users"
)

func (ur *UserRepository) CreateUSer(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, age) VALUES($1, $2) RETURNING id", tableUser)
	if err := ur.storage.db.QueryRow(
		query,
		u.Name,
		u.Age,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, bool, error) {
	users, err := ur.GetUsers()
	var faunded bool
	if err != nil {
		return nil, faunded, err
	}
	var userFinded *models.User
	for _, u := range users {
		if u.ID == id {
			userFinded = u
			faunded = true
			break
		}
	}
	return userFinded, faunded, nil
}

func (ur *UserRepository) GetUsers() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT id, name, age FROM %s;", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) DeleteUser(id int) (*models.User, error) {
	user, ok, err := ur.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", tableUser)
		_, err = ur.storage.db.Exec(query, id)
		return nil, err
	}
	return user, err
}

func (ur *UserRepository) UpdateUserAge(age int, id int) string {
	query := fmt.Sprintf("UPDATE %s SET age = %d WHERE id = %d;", tableUser, age, id)
	_, err := ur.storage.db.Query(query)
	if err != nil {
		log.Println(err)
	}
	return "возраст пользователя успешно обновлен"
}
