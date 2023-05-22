package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	config            *Config
	db                *sql.DB
	userRepository    *UserRepository
	friendsRepository *FriendRepository
}

func NewStorage(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Open() error {
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s port=%s user=%s password=%s sslmode=%s", s.config.DBname, s.config.Port, s.config.User, s.config.Password, s.config.SSLmode))
	if err != nil {
		return err
	}

	//Проверим, что все ок. Реально соединение тут не создается. Соединение только при первом вызове
	//db.Ping() // Пустой SELECT *
	if err = db.Ping(); err != nil {
		return err
	}
	s.db = db
	log.Println("Connection to db successfully")
	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
	}
	return s.userRepository
}

func (s *Storage) Friends() *FriendRepository {
	if s.friendsRepository != nil {
		return s.friendsRepository
	}
	s.friendsRepository = &FriendRepository{
		storage: s,
	}
	return s.friendsRepository
}
