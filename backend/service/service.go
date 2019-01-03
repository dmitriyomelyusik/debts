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

// DeleteUser deletes user
func (s Service) DeleteUser(id int) error {
	return s.db.DeleteUser(id)
}

// GetUser returns user by id
func (s Service) GetUser(id int) (domain.User, error) {
	return s.db.GetUser(id)
}

// GetUsers returns all users
func (s Service) GetUsers() ([]domain.User, error) {
	return s.db.GetUsers()
}

// AddDebt adds debt
func (s Service) AddDebt(debt domain.Debt) (domain.Debt, error) {
	if debt.Date == nil {
		return domain.Debt{}, errors.New("date can't be nil")
	}
	if _, err := s.GetUser(debt.Creditor.ID); err != nil {
		return domain.Debt{}, err
	}
	if _, err := s.GetUser(debt.Debtor.ID); err != nil {
		return domain.Debt{}, err
	}
	if debt.Sum <= 0 {
		return domain.Debt{}, errors.New("sum can't be not positive")
	}
	return s.db.AddDebt(debt)
}

// GetDebt returns debt by id
func (s Service) GetDebt(id int) (domain.Debt, error) {
	return s.db.GetDebt(id)
}

// GetDebts returns debts
func (s Service) GetDebts() ([]domain.Debt, error) {
	return s.db.GetDebts()
}

// DeleteDebt deletes debt by id
func (s Service) DeleteDebt(id int) error {
	return s.db.DeleteDebt(id)
}

// UpdateDebt updates debt
func (s Service) UpdateDebt(id int, debt domain.Debt) error {
	return s.db.UpdateDebt(id, debt)
}
