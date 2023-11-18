package app

import (
	"encoding/json"
	//"encoding/xml"
	"fmt"
	"github.com/fadedreams/gowallet/domain"
	"github.com/fadedreams/gowallet/service"
	"net/http"
)

type UserHandlers struct {
	service service.IUserService
}

type User struct {
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (ch *UserHandlers) createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = ch.service.SignUp(domain.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		Role:     newUser.Role,
	})

	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User created successfully")
}
