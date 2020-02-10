package mangopay

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type clientWalletService struct{}

// ListAll is listing all the ClientWallet with it's associated Currency and Fundtype.
func (cws *clientWalletService) ListAll() ([]model.ClientWallet, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, baseURL+"clients/wallets/", nil)
	if err != nil {
		return nil, err
	}
	return parseClientWalletCollection(data)
}

// View is retriving the ClientWallet from the given fundType and currency.
func (cws *clientWalletService) View(fundType model.FundsType, currency model.Currency) (*model.ClientWallet, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, baseURL+"clients/wallets/"+string(fundType)+"/"+string(currency)+"/", nil)
	if err != nil {
		return nil, err
	}
	return parseClientWallet(data)
}

// ListByFundsType is retriving all the ClientWallet from the given FundType.
func (cws *clientWalletService) ListByFundsType(fundType model.FundsType) ([]model.ClientWallet, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, baseURL+"clients/wallets/"+string(fundType)+"/", nil)
	if err != nil {
		return nil, err
	}
	return parseClientWalletCollection(data)
}

func parseClientWallet(data []byte) (*model.ClientWallet, error) {
	var clientWallet model.ClientWallet
	err := json.Unmarshal(data, &clientWallet)
	if err != nil {
		return nil, err
	}
	return &clientWallet, nil
}

func parseClientWalletCollection(data []byte) ([]model.ClientWallet, error) {
	var clientWalletList []model.ClientWallet
	err := json.Unmarshal(data, &clientWalletList)
	if err != nil {
		return nil, err
	}
	return clientWalletList, nil
}
