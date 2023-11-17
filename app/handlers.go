package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
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
	//clients := []Customer{
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
