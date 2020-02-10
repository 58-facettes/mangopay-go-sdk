package model

import "github.com/58-facettes/mangopay-go-sdk/model/currency"

type ClientWallet struct {
	//Balance is the current balance of the wallet
	Balance Money `json:"Balance"`
	// FundsType is the type of funds in the wallet
	FundsType Funds `json:"FundsType"`
}

type Money struct {
	// Currency is the currency - should be ISO_4217 format.
	Currency currency.ISO3 `json:"CurrencyIso"`
	// Amount of money in the smallest sub-division of the currency,
	// e.g. 12.60 EUR would be represented as 1260 whereas 12 JPY would be represented as just 12)
	Amount int `json:"Amount"`
}

type Funds string

const (
	FundsDefault Funds = "DEFAULT"
	FundsFees    Funds = "FEES"
	FundsCredit  Funds = "CREDIT"
)
