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
	FindAll() ([]Client, error)
}

// ///////////stub/////////////////////
type ClientRepositoryStub struct {
	clients []Client
}

func (s ClientRepositoryStub) FindAll() ([]Client, error) {
	return s.clients, nil
}

func NewClientRepositoryStub() ClientRepositoryStub {
	clients := []Client{
		{"1", "m", "m2", "1", "2000-01-01", "1"},
		{"2", "n", "n3", "2", "2000-02-02", "2"},
	}
	return ClientRepositoryStub{clients}
}
