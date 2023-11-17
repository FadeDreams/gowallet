package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/fadedreams/gowallet/domain"
	"github.com/fadedreams/gowallet/service"
)

func Start() {

	router := mux.NewRouter()
	ch := ClientHandlers{service.NewClientService(domain.NewClientRepositoryStub())}

	// define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/clients", ch.getAllClients).Methods(http.MethodGet)

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
