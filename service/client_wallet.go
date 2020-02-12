package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
	"github.com/58-facettes/mangopay-go-sdk/model/currency"
)

type ServiceClientWallet struct{}

// ListAll is listing all the ClientWallet with it's associated Currency and Fundtype.
func (ServiceClientWallet) ListAll(query ...model.Query) ([]model.ClientWallet, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, addQuery("clients/wallets/", query...), nil)
	if err != nil {
		return nil, err
	}
	return parseClientWalletCollection(data)
}

// View is retriving the ClientWallet from the given fundType and currency.
func (ServiceClientWallet) View(fundType model.Funds, currency currency.ISO3) (*model.ClientWallet, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "clients/wallets/"+string(fundType)+"/"+string(currency)+"/", nil)
	if err != nil {
		return nil, err
	}
	return parseClientWallet(data)
}

// ListByFundsType is retriving all the ClientWallet from the given FundType.
func (ServiceClientWallet) ListByFundsType(fundType model.Funds) ([]model.ClientWallet, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "clients/wallets/"+string(fundType)+"/", nil)
	if err != nil {
		return nil, err
	}
	return parseClientWalletCollection(data)
}

func parseClientWallet(data []byte) (*model.ClientWallet, error) {
	var clientWallet model.ClientWallet
	return &clientWallet, json.Unmarshal(data, &clientWallet)
}

func parseClientWalletCollection(data []byte) ([]model.ClientWallet, error) {
	var clientWalletList []model.ClientWallet
	return clientWalletList, json.Unmarshal(data, &clientWalletList)
}
