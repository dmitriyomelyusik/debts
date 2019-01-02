package service

import (
	"errors"

	"github.com/dmitriyomelyusik/debts/backend/database"
	"github.com/dmitriyomelyusik/debts/backend/domain"
)

// Service is service
type Service struct {
	db mysql.DB
}

// NewService returns new instance of service
func NewService(db mysql.DB) Service {
	return Service{db: db}
}

// AddUser adds user
func (s Service) AddUser(user domain.User) (domain.User, error) {
	if user.Name == "" {
		return domain.User{}, errors.New("user name can't be nil")
	}
	user, err := s.db.AddUser(user)
	return user, err
}

// UpdateUser updates user
func (s Service) UpdateUser(id int, user domain.User) error {
	if _, err := s.db.GetUser(id); err != nil {
		return errors.New("user is not found")
	}
	return s.db.UpdateUser(id, user)
}
