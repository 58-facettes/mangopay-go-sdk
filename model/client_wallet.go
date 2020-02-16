package model

import "github.com/58-facettes/mangopay-go-sdk/model/currency"

// ClientWallet is very similar to a normal wallet except the parameters "Description" and "Owners" are removed.
// Currently, there are two types of client wallet (specified by the "FundsType" parameter)
// - "FEES" where your collected turnover is stored and "CREDIT" where repudiations are taken from.
// A normal wallet for a user has the "FundsType" of "DEFAULT".
type ClientWallet struct {
	//Balance is the current balance of the wallet.
	Balance Money `json:"Balance"`
	// FundsType is the type of funds in the wallet.
	FundsType Funds `json:"FundsType"`
}

// Money represent the amout in the currency of a wallet.
type Money struct {
	// Currency is the currency - should be ISO_4217 format.
	Currency currency.ISO3 `json:"CurrencyIso"`
	// Amount of money in the smallest sub-division of the currency,
	// e.g. 12.60 EUR would be represented as 1260 whereas 12 JPY would be represented as just 12)
	Amount int `json:"Amount"`
}

// Funds is the type of fund.
type Funds string

const (
	// FundsDefault is used for type of default fund.
	FundsDefault Funds = "DEFAULT"
	// FundsFees is used for type of fee fund.
	FundsFees Funds = "FEES"
	// FundsCredit is used for type of credit fund.
	FundsCredit Funds = "CREDIT"
)
