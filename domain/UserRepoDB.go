package domain

import (
	"github.com/jmoiron/sqlx"
)

type UserRepositoryDb struct {
	db *sqlx.DB
}

func (d UserRepositoryDb) SignUp(user User) error {
	_, err := d.db.Exec(`
		INSERT INTO users (name, email, password, role)
		VALUES ($1, $2, $3, $4)
	`, user.Name, user.Email, user.Password, user.Role)

	if err != nil {
		return err
	}

	return nil
}

func NewUserRepositoryDb(dbClient *sqlx.DB) UserRepositoryDb {
	return UserRepositoryDb{db: dbClient}
}
