package mangopay

import (
	"time"
)

type API struct {
	isBasicAuth    bool
	clientID       string
	clientPassword string
	Clients        clientService
	ClientWallets  clientWalletService
	Users          userService
	UserEmoney     userEmoneyService
}

// NewWithBasicAuth sends a new Mangonpay client with Basic Auth.
func NewWithBasicAuth(cliendID, clientPassword string) *API {
	api := API{
		isBasicAuth:    true,
		clientID:       cliendID,
		clientPassword: clientPassword,
	}
	api.init()
	return &api
}

func (api *API) init() {
	if api.isBasicAuth {
		initBasicAuth(api.clientID, api.clientPassword)
	}
	initBaseURL(api.clientID)
}

func String(str string) *string {
	return &str
}

func Time(t time.Time) *int64 {
	res := t.UTC().Unix()
	return &res
}

func Int(i int) *int {
	return &i
}
