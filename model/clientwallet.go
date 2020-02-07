package model

type ClientWallet struct {
	//Balance is the current balance of the wallet
	Balance Money `json:"Balance"`
	// FundsType is the type of funds in the wallet
	FundsType FundsType `json:"FundsType"`
}

type Money struct {
	// CurrencyIso is the currency - should be ISO_4217 format
	CurrencyIso string `json:"CurrencyIso"`
	// Amount of money in the smallest sub-division of the currency, e.g. 12.60 EUR would be represented as 1260 whereas 12 JPY would be represented as just 12)
	Amount int `json:"Amount"`
}

type FundsType string

const (
	FundsDefault FundsType = "DEFAULT"
	FundsFees    FundsType = "FEES"
	FundsCredit  FundsType = "CREDIT"
)
