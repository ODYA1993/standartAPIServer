package storage

import (
	"fmt"
	"github.com/DmitryOdintsov/standartAPI_Server/internal/app/models"
	"log"
)

var (
	tableFriends = "friends"
)

type FriendsRepository struct {
	storage *Storage
}

func (fr *FriendsRepository) CreateFriends(f *models.Friends) (*models.Friends, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, friend_id) VALUES($1, $2) RETURNING id", tableFriends)

	if err := fr.storage.db.QueryRow(
		query,
		f.UserID,
		f.FriendID,
	).Scan(&f.ID); err != nil {
		return nil, err
	}
	return f, nil
}

func (fr *FriendsRepository) GetFriends(id int) ([]*models.User, error) {
	//query := fmt.Sprintf("SELECT id, user_id, friend_id FROM %s", tableFriends)
	query := fmt.Sprintf("with recursive cte as (SELECT us.id, fr.friend_id FROM friends fr JOIN users us ON us.id=fr.user_id WHERE us.id = %d UNION ALL SELECT us.id, fr.friend_id FROM friends fr JOIN cte ON cte.friend_id=fr.user_id JOIN users us ON us.id=fr.user_id WHERE fr.friend_id <>1)SELECT DISTINCT friend_id FROM cte", id)
	rows, err := fr.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID)
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
