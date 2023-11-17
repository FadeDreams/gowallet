package domain

import ()

type Client struct {
	Id          string `db:"client_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Client) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

type IClientRepository interface {
	FindAll(status string) ([]Client, error)
}
