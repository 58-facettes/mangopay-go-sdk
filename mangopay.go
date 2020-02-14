package mangopay

import (
	"encoding/base64"
	"log"

	"github.com/58-facettes/mangopay-go-sdk/model"
	"github.com/58-facettes/mangopay-go-sdk/service"
)

type API struct {
	isBasicAuth    bool
	clientID       string
	clientPassword string
	// List of all avalable services.
	Clients       *service.ServiceClient
	ClientWallets *service.ServiceClientWallet
	Users         *service.ServiceUser
	UserEmoney    *service.ServiceUserEmoney
	Wallets       *service.ServiceWallet
	Cards         *service.ServiceCard
	PayIns        *service.ServicePayIn
	Transferts    *service.ServiceTransfert
	BankAccounts  *service.ServiceBankAccount
	PayOuts       *service.ServicePayOut
	KYCs          *service.ServiceKYC
	Stats         *service.ServiceStat
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

var basicAuth string

// initBasicAuth set the basicAuth variable that will be used with "Authorization" in the header.
func initBasicAuth(user, password string) {
	basicAuth = "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password))
}

const (
	// APIVersion is the current API version that is used in the URI.
	APIVersion = "v2.01"
)

const (
	// ModeTest is for using the API sandbox for testing purposes.
	ModeTest = "test"
	// ModeProduction is of the API production usage.
	ModeProduction = "production"
)

// Mode is the mode of the SDK calls by defaults this is set to test mode.
var Mode = ModeTest

var baseURL string

func initBaseURL(ClientID string) {
	switch Mode {
	case ModeProduction:
		baseURL = "https://api.mangopay.com/" + APIVersion + "/" + ClientID + "/"
	default:
		baseURL = "https://api.sandbox.mangopay.com/" + APIVersion + "/" + ClientID + "/"
	}
}

func (api *API) RateLimits(rate model.Rate) string {
	rl, err := api.Stats.GetRateLimit()
	if err != nil {
		log.Println("mangopay api: ratelimit %v", err.Error())
	}
	return rl.GetData(rate)
}

// TemPath is the temporary path that can be used for files.
var TemPath string
