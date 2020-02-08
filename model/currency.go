package model

type Currency string

func (c Currency) String() string {
	return string(c)
}

const (
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	PLN Currency = "PLN"
	CHF Currency = "CHF"
	NOK Currency = "NOK"
	SEK Currency = "SEK"
	DKK Currency = "DKK"
	USD Currency = "USD"
	CAD Currency = "CAD"
	AUD Currency = "AUD"
	JPY Currency = "JPY"
	HKD Currency = "HKD"
	AED Currency = "AED"
	ZAR Currency = "ZAR"
	CZK Currency = "CZK"
)
