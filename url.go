package sdk

import (
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

const (
	APIVersion = "v2.01"
)

type MethodURL struct{}

// UpdateClient is the method and url for updating a Client.
func (MethodURL) UpdateClient(ClientID string) (string, string) {
	return http.MethodPut, "/" + APIVersion + "/" + ClientID + "/clients/"
}

// UploadClientLogo is uploading a client logo.
func (MethodURL) UploadClientLogo(ClientID string) (string, string) {
	return http.MethodPut, "/" + APIVersion + "/" + ClientID + "/clients/logo/"
}

// ViewClient is for viewing a client.
func (MethodURL) ViewClient(ClientID string) (string, string) {
	return http.MethodGet, "/" + APIVersion + "/" + ClientID + "/clients/"
}

// ListallClientWallets is listing all Client Wallets.
// It is possible to view your Fees and Credit Wallets thanks to the API calls below.
func (MethodURL) ListallClientWallets(ClientID string) (string, string) {
	return http.MethodGet, "/" + APIVersion + "/" + ClientID + "/clients/wallets/"
}

// ViewClientWallet is for viewing a Client Wallet.
func (MethodURL) ViewClientWallet(ClientID string, funds model.FundsType, currency string) (string, string) {
	return http.MethodGet, "/" + APIVersion + "/" + ClientID + "/clients/wallets/" + string(funds) + "/" + currency + "/"
}

// ListClientWalletsFundsType is listing Client Wallets by FundsType.
func (MethodURL) ListClientWalletsFundsType(ClientID string, funds model.FundsType) (string, string) {
	return http.MethodGet, "/" + APIVersion + "/" + ClientID + "/clients/wallets/" + string(funds)
}
