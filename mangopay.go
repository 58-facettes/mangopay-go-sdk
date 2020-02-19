package mangopay

import (
	"encoding/base64"

	"github.com/58-facettes/mangopay-go-sdk/log"
	"github.com/58-facettes/mangopay-go-sdk/model"
	"github.com/58-facettes/mangopay-go-sdk/service"
)

// API holds all the services for calling Mongopay API.
type API struct {
	isBasicAuth    bool
	clientID       string
	clientPassword string
	logr           log.Logger
	// List of all avalable services.
	Clients             *service.Clients
	ClientWallets       *service.ClientWallets
	Users               *service.Users
	UserEmoney          *service.UserEmoneys
	Wallets             *service.Wallets
	Cards               *service.Cards
	PayIns              *service.PayIns
	Transferts          *service.Transferts
	BankAccounts        *service.BankAccounts
	PayOuts             *service.PayOuts
	KYCs                *service.KYCs
	Stats               *service.Stats
	UBOs                *service.UBOs
	Mandates            *service.Mandates
	Hooks               *service.Hooks
	Events              *service.Events
	Transactions        *service.Transactions
	PreAuthorizations   *service.PreAuthorizations
	Refunds             *service.Refunds
	Disputes            *service.Disputes
	DisputeDocuments    *service.DisputeDocunents
	Repudiations        *service.Repudiations
	SettlementTransfers *service.SettlementTranfers
	BankingAliases      *service.BankingAliases
	SSOS                *service.SSOS
	PermissionGroups    *service.PermissionGoups
	// Reporting *service.Reporting
	Idempotencies *service.Idempotcencies
}

// Logger is the default internal logging tool that is used.
// This can be replaced by another logging tool of your choice like Zap or Logrus.
var Logger log.Logger = log.DefaultLogger

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
	api.logr = Logger
	service.SetLogger(Logger)
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
		api.logr.Warnf("mangopay: ratelimit %v", err.Error())
	}
	return rl.GetData(rate)
}

// TemPath is the temporary path that can be used for files.
var TemPath string
