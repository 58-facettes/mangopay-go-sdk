package mangopay

import (
	"encoding/base64"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/data"
	"github.com/58-facettes/mangopay-go-sdk/log"
	"github.com/58-facettes/mangopay-go-sdk/service"
)

const (
	// APIVersion is the current API version in the URI calls.
	APIVersion = "v2.01"
	// Test is for using the API sandbox.
	Test = "test"
	// Production is for the API production.
	Production = "production"
)

type config struct {
	// env mode.
	Mode string
	// logger used.
	Logger log.Logger
	// DB used for idempotency.
	DB data.Manager
	// Use idempotency.
	UseIdempotency bool
	// UseBasicAuth for knowing if it use the basicAuth.
	UseBasicAuth bool
	// HTTPClient for http calls you can bring your own by default this is the http.Client{}.
	HTTPClient *http.Client
}

// Config allow the way you desire to use the API.
// by default this is using an internal logger with a basicAuth in tesing Mode.
var Config = &config{
	Logger:         log.DefaultLogger,
	UseBasicAuth:   true,
	Mode:           Test,
	UseIdempotency: false,
	DB:             nil, // we don't use idempotency to we don't need to store the keys.
	HTTPClient:     http.DefaultClient,
}

// API holds all the services for calling Mongopay API.
type API struct {
	logr log.Logger
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
	api := new(API)
	service.DefaultClient = Config.HTTPClient
	// this way we oblige the usage of a db when the idempotency is used.
	if Config.DB != nil && Config.UseIdempotency {
		service.DB = Config.DB
		service.UseIdempotency = Config.UseIdempotency
	} else if service.DB == nil && Config.UseIdempotency {
		Config.Logger.Fatal("mangopay: no storage tool found in the mangopay.Config.DB.\n" +
			"If you use an idempotency key you should provide your db in mangopay.Config.DB.\n" +
			"Please refer to the data.Manager contract interface.")
	}
	if Config.Logger != nil {
		api.logr = Config.Logger
	}
	service.SetLogger(Config.Logger)
	// init basicAuth.
	Config.UseBasicAuth = isBasicAuth
	service.UseBasicAuth = Config.UseBasicAuth
	service.BasicAuth = "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientPassword))
	// init base URL.
	switch Config.Mode {
	case Production:
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
	return api
}
