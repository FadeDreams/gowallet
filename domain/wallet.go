package domain

type Wallet struct {
	WalletId   string
	ClientId   string
	WalletType string
	Amount     float64
}

type IWalletRepository interface {
	Save(Wallet) (*Wallet, error)
}
