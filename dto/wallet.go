package dto

type ClientResponse struct {
	Id          string `json:"client_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

type NewWalletRequest struct {
	ClientId   string  `json:"client_id"`
	WalletType string  `json:"type"`
	Amount     float64 `json:"amount"`
}

type NewWalletResponse struct {
	WalletId string `json:"wallet_id"`
}
