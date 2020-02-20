package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
	"github.com/58-facettes/mangopay-go-sdk/model/currency"
)

// ClientWallets is responsible of all services for the Client'Wallet(s).
type ClientWallets struct{}

// ListAll is listing all the ClientWallet with it's associated Currency and Fundtype.
func (ClientWallets) ListAll(query *model.Query) (res []model.ClientWallet, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("clients/wallets/", query), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// View is retriving the ClientWallet from the given fundType and currency.
func (ClientWallets) View(fundType model.Funds, currency currency.ISO3) (res *model.ClientWallet, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "clients/wallets/"+string(fundType)+"/"+string(currency)+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// ListByFundsType is retriving all the ClientWallet from the given FundType.
func (ClientWallets) ListByFundsType(fundType model.Funds, query *model.Query) (res []model.ClientWallet, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("clients/wallets/"+string(fundType)+"/", query), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
