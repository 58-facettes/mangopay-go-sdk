package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// BankingAliases is responsible of all services for BankingAlias.
type BankingAliases struct{}

// Create is creating a ne BankingAlias connect with the walletID.
func (BankingAliases) Create(walletID string, payload *model.BankingAliasCreate) (res *model.BankingAlias, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "wallets/"+walletID+"/bankingaliases/iban/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Desactivate is desactivating a BankingAlias
func (BankingAliases) Desactivate(walletID, bankingAliasID string) (res *model.BankingAlias, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "wallets/"+walletID+"/bankingaliases/"+bankingAliasID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View reteve a BankingAlias by it's given ID.
func (BankingAliases) View(bankingAliasID string) (res *model.BankingAlias, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "bankingaliases/"+bankingAliasID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// WalletView retreve a BankingAlias from it's wallet's ID.
// To view a wallet's banking aliases - remember that you can only have one banking alias per wallet (even though it is a list format).
func (BankingAliases) WalletView(walletID string) (res *model.BankingAlias, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "wallets/"+walletID+"/bankingaliases/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}
