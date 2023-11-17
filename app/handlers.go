package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Client struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}

func getAllClients(w http.ResponseWriter, r *http.Request) {
	clients := []Client{
		{"Ashish", "New Delhi", "110075"},
		{"Rob", "New Delhi", "110075"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(clients)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(clients)
	}
}
