package models

type Friends struct {
	ID     int `json:"id"`
	User   int `json:"user_id"`
	Friend int `json:"friend_id"`
}
