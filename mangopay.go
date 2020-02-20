package mangopay

import (
	"encoding/base64"

	"github.com/58-facettes/mangopay-go-sdk/log"
	"github.com/58-facettes/mangopay-go-sdk/model"
	"github.com/58-facettes/mangopay-go-sdk/service"
)

const (
	// APIVersion is the current API version in the URI calls.
	APIVersion = "v2.01"
	// ModeTest is for using the API sandbox.
	ModeTest = "test"
	// ModeProduction is for the API production.
	ModeProduction = "production"
)

var (
	// Mode is the mode of the SDK calls by defaults this is set to test mode.
	Mode = ModeTest
	// Logger is the default internal logging tool that is used.
	// This can be replaced by another logging tool of your choice like Zap or Logrus.
	Logger log.Logger = log.DefaultLogger
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
	Reports             *service.Reports
	Idempotencies       *service.Idempotencies
}

// NewWithBasicAuth sends a new Mangonpay client with Basic Auth.
func NewWithBasicAuth(clientID, clientPassword string) *API {
	return newConnect(clientID, clientPassword, true)
}

// NewWithOAuth is used for an oauth token connexion.
func NewWithOAuth(clientID, clientPassword string) *API {
	return newConnect(clientID, clientPassword, false)
}

func newConnect(clientID, clientPassword string, isBasicAuth bool) *API {
	api := API{
		isBasicAuth:    isBasicAuth,
		clientID:       clientID,
		clientPassword: clientPassword,
	}
	api.logr = Logger
	service.SetLogger(Logger)
	// init basicAuth.
	service.UseBasicAuth = isBasicAuth
	service.BasicAuth = "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientPassword))
	// init base URL.
	switch Mode {
	case ModeProduction:
		service.BaseURL = "https://api.mangopay.com/" + APIVersion + "/" + clientID + "/"
	default:
		service.BaseURL = "https://api.sandbox.mangopay.com/" + APIVersion + "/" + clientID + "/"
	}
	api.Clients = new(service.Clients)
	api.ClientWallets = new(service.ClientWallets)
	api.Users = new(service.Users)
	api.UserEmoney = new(service.UserEmoneys)
	api.Wallets = new(service.Wallets)
	api.Cards = new(service.Cards)
	api.PayIns = new(service.PayIns)
	api.Transferts = new(service.Transferts)
	api.BankAccounts = new(service.BankAccounts)
	api.PayOuts = new(service.PayOuts)
	api.KYCs = new(service.KYCs)
	api.Stats = new(service.Stats)
	api.UBOs = new(service.UBOs)
	api.Mandates = new(service.Mandates)
	api.Hooks = new(service.Hooks)
	api.Events = new(service.Events)
	api.Transactions = new(service.Transactions)
	api.PreAuthorizations = new(service.PreAuthorizations)
	api.Refunds = new(service.Refunds)
	api.Disputes = new(service.Disputes)
	api.DisputeDocuments = new(service.DisputeDocunents)
	api.Repudiations = new(service.Repudiations)
	api.SettlementTransfers = new(service.SettlementTranfers)
	api.BankingAliases = new(service.BankingAliases)
	api.SSOS = new(service.SSOS)
	api.PermissionGroups = new(service.PermissionGoups)
	api.Reports = new(service.Reports)
	api.Idempotencies = new(service.Idempotencies)
	return &api
}

// RateLimits retreve a rate limit value from the given rate range.
func (api *API) RateLimits(rate model.Rate) string {
	rl, err := api.Stats.GetRateLimit()
	if err != nil {
		api.logr.Warnf("mangopay: ratelimit %v", err.Error())
	}
	return rl.GetData(rate)
}
