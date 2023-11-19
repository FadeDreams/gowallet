package domain

import (
	"net/http"
)

type User struct {
	Id       string `db:"user_id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type IUserRepository interface {
	SignUp(User) error
	SignIn(string, string) (string, error)
	IsAuthorized(http.HandlerFunc) http.HandlerFunc
}
