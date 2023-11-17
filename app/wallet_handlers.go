// handlers.go in the app package

package app

import (
	"encoding/json"
	"github.com/fadedreams/gowallet/dto"

	"github.com/fadedreams/gowallet/errors"
	"net/http"

	"github.com/fadedreams/gowallet/domain"
	"github.com/fadedreams/gowallet/service"
)

type WalletHandlers struct {
	service service.IWalletService
}

func (h WalletHandlers) newWallet(w http.ResponseWriter, r *http.Request) {
	var request dto.NewWalletRequest

	// Decode the incoming JSON request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Extract the client ID from the URL parameters
	//clientID := mux.Vars(r)["clientID"]

	// Convert the NewWalletRequest to a domain.Wallet
	newWallet := domain.Wallet{
		ClientId:   request.ClientId,
		WalletType: request.WalletType,
		Amount:     request.Amount,
	}

	// Call the IWalletService to create a new wallet
	response, err := h.service.CreateWallet(newWallet)
	if err != nil {
		http.Error(w, "Error creating wallet", http.StatusInternalServerError)
		return
	}

	// Create a NewWalletResponse instance
	newWalletResponse := dto.NewWalletResponse{
		WalletId: response.WalletId,
	}

	// Encode the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newWalletResponse)
}

func (h WalletHandlers) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	var request dto.TransactionRequest

	// Decode the incoming JSON request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Call the IWalletService to make a transaction
	response, appError := h.service.MakeTransaction(request)
	if appError != nil {
		// Handle the specific errors as needed
		if appError.Code == errors.ValidationErrorCode {
			http.Error(w, appError.Message, http.StatusBadRequest)
		} else if appError.Code == errors.WithdrawalError {
			http.Error(w, appError.Message, http.StatusUnprocessableEntity)
		} else {
			http.Error(w, "Error making transaction", http.StatusInternalServerError)
		}
		return
	}

	// Encode the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
