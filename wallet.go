package mangopay

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type walletService struct{}

// Create is creating a new Wallet.
func (walletService) Create(param *model.WalletCreate) (*model.Wallet, error) {
	_, data, err := newRequestAndExecute(http.MethodPost, baseURL+"wallets/", param)
	if err != nil {
		return nil, err
	}
	return parseWallet(data)
}

// Update is updating the description of an existing Wallet with the given params.
// This method allow to have a nil description into it.
func (walletService) Update(walletID string, param *model.WalletUpdate) (*model.Wallet, error) {
	_, data, err := newRequestAndExecute(http.MethodPut, baseURL+"wallets/"+walletID, param)
	if err != nil {
		return nil, err
	}
	return parseWallet(data)
}

// UpdateDes is a helper that simplify the way to update the description of an existing Wallet with the given params.
func (walletService) UpdateDesc(walletID, desc string) (*model.Wallet, error) {
	_, data, err := newRequestAndExecute(http.MethodPut, baseURL+"wallets/"+walletID, &model.WalletUpdate{
		Description: String(desc),
	})
	if err != nil {
		return nil, err
	}
	return parseWallet(data)
}

// View retrieve the Wallet front the given walletID.
func (walletService) View(walletID string) (*model.Wallet, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, baseURL+"wallets/"+walletID, nil)
	if err != nil {
		return nil, err
	}
	return parseWallet(data)
}

// View retrieve all the Wallets front a given userID.
func (walletService) ListFromUser(userID string) ([]model.Wallet, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, baseURL+"users/"+userID+"/wallets/", nil)
	if err != nil {
		return nil, err
	}
	return parseWalletCollection(data)
}

func parseWallet(data []byte) (*model.Wallet, error) {
	var res model.Wallet
	return &res, json.Unmarshal(data, &res)
}

func parseWalletCollection(data []byte) ([]model.Wallet, error) {
	var res []model.Wallet
	return res, json.Unmarshal(data, &res)
}
