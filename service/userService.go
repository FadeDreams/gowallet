package service

import (
	"github.com/fadedreams/gowallet/domain"
	"net/http"
)

type IUserService interface {
	SignUp(domain.User) error
	SignIn(string, string) (string, error)
	IsAuthorized(http.HandlerFunc) http.HandlerFunc
}

type DefaultUserService struct {
	repo domain.IUserRepository
}

//func (s DefaultClientService) Register() ([]domain.Client, error) {
//return s.repo.FindAll()
//}

func (s DefaultUserService) SignUp(newUser domain.User) error {
	return s.repo.SignUp(newUser)
}

func (s DefaultUserService) SignIn(email string, password string) (string, error) {
	return s.repo.SignIn(email, password)
}

func (s DefaultUserService) IsAuthorized(hf http.HandlerFunc) http.HandlerFunc {
	return s.repo.IsAuthorized(hf)
}

func NewUserService(repository domain.IUserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
