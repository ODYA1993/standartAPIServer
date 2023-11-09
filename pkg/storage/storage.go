package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

func NewStorage(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Open() error {
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s port=%s user=%s password=%s sslmode=%s", s.config.DBname, s.config.Port, s.config.User, s.config.Password, s.config.Sslmode))
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}
	s.db = db
	logrus.Println("Успешное подключение к базе данных")
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
