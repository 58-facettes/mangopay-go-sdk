package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type ServiceWallet struct{}

// Create is creating a new Wallet.
func (ServiceWallet) Create(payload *model.WalletCreate) (res *model.Wallet, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "wallets/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Update is updating the description of an existing Wallet with the given payloads.
// This method allow to have a nil description into it.
func (ServiceWallet) Update(walletID string, payload *model.WalletUpdate) (res *model.Wallet, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "wallets/"+walletID, payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// UpdateDesc is a helper that simplify the way to update the description of an existing Wallet with the given payloads.
func (ServiceWallet) UpdateDesc(walletID, desc string) (res *model.Wallet, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "wallets/"+walletID, &model.WalletUpdate{
		Description: desc,
	})
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View retrieve the Wallet front the given walletID.
func (ServiceWallet) View(walletID string) (res *model.Wallet, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "wallets/"+walletID, nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// ListFromUser retrieve all the Wallets front a given userID.
func (ServiceWallet) ListFromUser(userID string, query ...model.Query) (res []model.Wallet, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, addQuery("users/"+userID+"/wallets/", query...), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
