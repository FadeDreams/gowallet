package domain

import ()

type Client struct {
	Id      string `db:"client_id"`
	Name    string
	City    string
	Zipcode string
	Status  string
}

func (c Client) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

type IClientRepository interface {
	FindAll() ([]Client, error)
	CreateClient(Client) error
}
