package service

import "github.com/fadedreams/gowallet/domain"

type IUserService interface {
	SignUp(domain.User) error
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

func NewUserService(repository domain.IUserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
