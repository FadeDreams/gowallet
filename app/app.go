package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/clients", getAllClients).Methods(http.MethodGet)
	router.HandleFunc("/clients", createClient).Methods(http.MethodPost)

	router.HandleFunc("/clients/{client_id:[0-9]+}", getClient).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

func createClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}

func getClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["client_id"])
}
