package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
	"github.com/58-facettes/mangopay-go-sdk/model/currency"
)

type ServiceClientWallet struct{}

// ListAll is listing all the ClientWallet with it's associated Currency and Fundtype.
func (ServiceClientWallet) ListAll(query ...model.Query) (res []model.ClientWallet, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, addQuery("clients/wallets/", query...), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View is retriving the ClientWallet from the given fundType and currency.
func (ServiceClientWallet) View(fundType model.Funds, currency currency.ISO3) (res *model.ClientWallet, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "clients/wallets/"+string(fundType)+"/"+string(currency)+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// ListByFundsType is retriving all the ClientWallet from the given FundType.
func (ServiceClientWallet) ListByFundsType(fundType model.Funds) (res []model.ClientWallet, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "clients/wallets/"+string(fundType)+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}
