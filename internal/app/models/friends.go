package models

type Friends struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	FriendID int `json:"friend_id"`
}
