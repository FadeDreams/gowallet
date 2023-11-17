package domain

import (
	"github.com/jmoiron/sqlx"
	"log"
	//"strconv"
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

func NewWalletRepositoryDb(dbClient *sqlx.DB) WalletRepositoryDb {
	return WalletRepositoryDb{dbClient}
}
