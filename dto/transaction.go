package dto

import "github.com/fadedreams/gowallet/errors"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	WalletId        string  `json:"wallet_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"-"`
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	if r.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}

func (r TransactionRequest) Validate() *errors.AppError {
	if r.TransactionType != WITHDRAWAL && r.TransactionType != DEPOSIT {
		return errors.NewValidationError("Transaction type can only be deposit or withdrawal")
	}
	if r.Amount < 0 {
		return errors.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	WalletId        string  `json:"wallet_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
