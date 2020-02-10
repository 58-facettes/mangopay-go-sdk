package model

type CurrencyISO3 string

func (c CurrencyISO3) String() string {
	return string(c)
}

const (
	EUR CurrencyISO3 = "EUR"
	GBP CurrencyISO3 = "GBP"
	PLN CurrencyISO3 = "PLN"
	CHF CurrencyISO3 = "CHF"
	NOK CurrencyISO3 = "NOK"
	SEK CurrencyISO3 = "SEK"
	DKK CurrencyISO3 = "DKK"
	USD CurrencyISO3 = "USD"
	CAD CurrencyISO3 = "CAD"
	AUD CurrencyISO3 = "AUD"
	JPY CurrencyISO3 = "JPY"
	HKD CurrencyISO3 = "HKD"
	AED CurrencyISO3 = "AED"
	ZAR CurrencyISO3 = "ZAR"
	CZK CurrencyISO3 = "CZK"
)
