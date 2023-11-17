package domain

import (
	"github.com/fadedreams/gowallet/errors"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type WalletRepositoryDb struct {
	client *sqlx.DB
}

func (d WalletRepositoryDb) Save(a Wallet) (*Wallet, error) {
	sqlInsert := "INSERT INTO wallets (client_id, wallet_type, amount) VALUES ($1, $2, $3) RETURNING wallet_id"

	result := d.client.QueryRowx(sqlInsert, a.ClientId, a.WalletType, a.Amount)

	err := result.Scan(&a.WalletId)
	if err != nil {
		log.Println("Error while creating new wallet: " + err.Error())
		return nil, err
	}

	return &a, nil
}

func (d WalletRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errors.AppError) {
	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		log.Println("Error while starting a new transaction for bank wallet transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank wallet transaction
	result, _ := tx.Exec(`INSERT INTO transactions (wallet_id, amount, transaction_type, transaction_date) 
											values (?, ?, ?, ?)`, t.WalletId, t.Amount, t.TransactionType, t.TransactionDate)

	// updating wallet balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE wallets SET amount = amount - ? where wallet_id = ?`, t.Amount, t.WalletId)
	} else {
		_, err = tx.Exec(`UPDATE wallets SET amount = amount + ? where wallet_id = ?`, t.Amount, t.WalletId)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		log.Println("Error while saving transaction: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println("Error while commiting transaction for bank wallet: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	// getting the last transaction ID from the transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		log.Println("Error while getting the last transaction id: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	// Getting the latest wallet information from the wallets table
	wallet, appErr := d.FindBy(t.WalletId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// updating the transaction struct with the latest balance
	t.Amount = wallet.Amount
	return &t, nil
}

func (d WalletRepositoryDb) FindBy(walletId string) (*Wallet, *errors.AppError) {
	sqlGetWallet := "SELECT wallet_id, customer_id, opening_date, wallet_type, amount from wallets where wallet_id = ?"
	var wallet Wallet
	err := d.client.Get(&wallet, sqlGetWallet, walletId)
	if err != nil {
		log.Println("Error while fetching wallet information: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}
	return &wallet, nil
}

func NewWalletRepositoryDb(dbClient *sqlx.DB) WalletRepositoryDb {
	return WalletRepositoryDb{dbClient}
}
