// app_test.go
package app

import (
	"bytes"
	"encoding/json"
	// "encoding/xml"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fadedreams/gowallet/domain"
	// "github.com/fadedreams/gowallet/service"
)

// MockClientService is a mock implementation of service.IClientService for testing.
type MockClientService struct{}

func (mcs *MockClientService) GetAllClient() ([]domain.Client, error) {
	// Implement mock behavior for GetAllClient if needed.
	return nil, nil
}

func (mcs *MockClientService) CreateClient(client domain.Client) error {
	// Implement mock behavior for CreateClient if needed.
	return nil
}

func TestGetAllClients(t *testing.T) {
	// Create an instance of ClientHandlers with the mock service.
	clientHandlers := ClientHandlers{service: &MockClientService{}}

	// Create a request to the endpoint.
	req, err := http.NewRequest("GET", "/clients", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response.
	rr := httptest.NewRecorder()

	// Create an http.HandlerFunc from the method and serve the request.
	handler := http.HandlerFunc(clientHandlers.getAllClients)
	handler.ServeHTTP(rr, req)

	// Check the status code is 200 OK.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body or other expectations.
	// You may want to check the format based on the Content-Type header.
}

func TestCreateClient(t *testing.T) {
	// Create an instance of ClientHandlers with the mock service.
	clientHandlers := ClientHandlers{service: &MockClientService{}}

	// Create a sample client for testing.
	newClient := Client{
		Name:    "John Doe",
		City:    "Example City",
		Zipcode: "12345",
	}

	// Convert the client to JSON.
	jsonClient, err := json.Marshal(newClient)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request to the endpoint with the JSON payload.
	req, err := http.NewRequest("POST", "/clients", bytes.NewBuffer(jsonClient))
	if err != nil {
		t.Fatal(err)
	}

	// Set the Content-Type header to application/json.
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response.
	rr := httptest.NewRecorder()

	// Create an http.HandlerFunc from the method and serve the request.
	handler := http.HandlerFunc(clientHandlers.createClient)
	handler.ServeHTTP(rr, req)

	// Check the status code is 201 Created.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check the response body or other expectations.
}
