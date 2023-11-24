package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/fadedreams/gowallet/domain"
	"github.com/fadedreams/gowallet/service"
	"net/http"
)

type ClientHandlers struct {
	service service.IClientService
}

type Client struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}

func (ch *ClientHandlers) getAllClients(w http.ResponseWriter, r *http.Request) {
	//clients := []Client{
	//	{"m", "m2", "1"},
	//	{"n", "n3", "2"},
	//}

	clients, _ := ch.service.GetAllClient()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(clients)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(clients)
	}
}

func (ch *ClientHandlers) createClient(w http.ResponseWriter, r *http.Request) {
	var newClient Client
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newClient)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = ch.service.CreateClient(domain.Client{
		Name:    newClient.Name,
		City:    newClient.City,
		Zipcode: newClient.Zipcode,
	})
	if err != nil {
		http.Error(w, "Error creating client", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Client created successfully")
}
