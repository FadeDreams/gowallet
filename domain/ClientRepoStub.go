package domain

// ///////////stub/////////////////////
type ClientRepositoryStub struct {
	clients []Client
}

func (s ClientRepositoryStub) FindAll() ([]Client, error) {
	return s.clients, nil
}

func NewClientRepositoryStub() ClientRepositoryStub {
	clients := []Client{
		{"1", "m", "m2", "1", "1"},
		{"2", "n", "n3", "2", "2"},
	}
	return ClientRepositoryStub{clients}
}
