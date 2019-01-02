package mysql

import (
	"database/sql"
	"log"

	"github.com/dmitriyomelyusik/debts/backend/domain"
)

// DB contains db needs
type DB struct {
	db *sql.DB
}

// NewDB returns new instance of db
func NewDB() DB {
	db, err := sql.Open("mysql", "root:password@/debts?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return DB{db: db}
}

// AddUser adds user into database
func (db DB) AddUser(user domain.User) (domain.User, error) {
	res, err := db.db.Exec("INSERT INTO users (name) VALUES (?)", user.Name)
	if err != nil {
		return domain.User{}, err
	}
	id, _ := res.LastInsertId()
	user.ID = int(id)
	return user, nil
}

// UpdateUser updates user
func (db DB) UpdateUser(id int, user domain.User) error {
	_, err := db.db.Exec("UPDATE users SET name=? WHERE id=?", user.Name, id)
	return err
}

// DeleteUser deletes user
func (db DB) DeleteUser(id int) error {
	_, err := db.db.Exec("DELETE users WHERE id=?", id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

// GetUser returns user by id
func (db DB) GetUser(id int) (domain.User, error) {
	var user domain.User
	err := db.db.QueryRow("SELECT FROM users WHERE id=?", id).Scan(&user.ID, &user.Name)
	return user, err
}

// GetUsers returns all users
func (db DB) GetUsers() ([]domain.User, error) {
	rows, err := db.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	users := make([]domain.User, 0)
	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

// AddDebt adds debt into database
func (db DB) AddDebt(debt domain.Debt) (domain.Debt, error) {
	res, err := db.db.Exec("INSERT INTO debts (creditor, debtor, sum, date) VALUES (?, ?, ?, ?)", debt.Creditor.ID,
		debt.Debtor.ID, debt.Sum, debt.Date)
	id, _ := res.LastInsertId()
	debt.ID = int(id)
	return debt, err
}

// UpdateDebt updates debt
func (db DB) UpdateDebt(debt domain.Debt) error {
	_, err := db.db.Exec("UPDATE debts SET sum=?, date=?, creditor=?, debtor=? WHERE id=?", debt.Sum, debt.Date, debt.Creditor,
		debt.Debtor, debt.ID)
	return err
}

// DeleteDebt deletes debt
func (db DB) DeleteDebt(id int) error {
	_, err := db.db.Exec("DELETE debts WHERE id=?", id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

// GetDebt returns debt by id
func (db DB) GetDebt(id int) (domain.Debt, error) {
	var debt domain.Debt
	err := db.db.QueryRow("SELECT FROM debts WHERE id=?", id).Scan(&debt.ID, &debt.Creditor, &debt.Debtor, &debt.Sum, &debt.Date)
	return debt, err
}

// GetDebts returns all debts
func (db DB) GetDebts() ([]domain.Debt, error) {
	rows, err := db.db.Query("SELECT * FROM debts")
	if err != nil {
		return nil, err
	}
	debts := make([]domain.Debt, 0)
	for rows.Next() {
		var debt domain.Debt
		err = rows.Scan(&debt.ID, &debt.Creditor, &debt.Debtor, &debt.Sum, &debt.Date)
		if err != nil {
			return debts, err
		}
		debts = append(debts, debt)
	}
	return debts, nil
}
