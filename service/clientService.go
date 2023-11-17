package service

import "github.com/fadedreams/gowallet/domain"

type IClientService interface {
	GetAllClient() ([]domain.Client, error)
}

type DefaultClientService struct {
	repo domain.IClientRepository
}

func (s DefaultClientService) GetAllClient() ([]domain.Client, error) {
	return s.repo.FindAll()
}

func NewClientService(repository domain.IClientRepository) DefaultClientService {
	return DefaultClientService{repository}
}
