package service

import (
	"github.com/fadedreams/gowallet/domain"
	"github.com/fadedreams/gowallet/dto"
	"github.com/fadedreams/gowallet/errors"
	"time"
)

const dbTSLayout = "2006-01-02 15:04:05"

type IWalletService interface {
	CreateWallet(newWallet domain.Wallet) (*domain.Wallet, error)
	MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError)
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

func (s DefaultWalletService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errors.AppError) {
	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// server side validation for checking the available balance in the wallet
	if req.IsTransactionTypeWithdrawal() {
		wallet, err := s.repo.FindBy(req.WalletId)
		if err != nil {
			return nil, err
		}
		if !wallet.CanWithdraw(req.Amount) {
			return nil, errors.NewValidationError("Insufficient balance in the wallet")
		}
	}
	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		WalletId:        req.WalletId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}
