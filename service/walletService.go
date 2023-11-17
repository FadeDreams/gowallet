package service

import (
	"github.com/fadedreams/gowallet/domain"
	//"github.com/fadedreams/gowallet/dto"
)

type IWalletService interface {
	CreateWallet(newWallet domain.Wallet) (*domain.Wallet, error)
}

type DefaultWalletService struct {
	repo domain.IWalletRepository
}

func (s DefaultWalletService) CreateWallet(newWallet domain.Wallet) (*domain.Wallet, error) {
	return s.repo.Save(newWallet)
}

func NewWalletService(repository domain.IWalletRepository) DefaultWalletService {
	return DefaultWalletService{repository}
}
