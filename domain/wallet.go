package domain

import (
	"github.com/fadedreams/gowallet/errors"
)

type Wallet struct {
	WalletId   string
	ClientId   string
	WalletType string
	Amount     float64
}

type IWalletRepository interface {
	Save(Wallet) (*Wallet, error)
	SaveTransaction(transaction Transaction) (*Transaction, *errors.AppError)
	FindBy(string) (*Wallet, *errors.AppError)
}

func (w Wallet) CanWithdraw(amount float64) bool {
	if w.Amount < amount {
		return false
	}
	return true
}
