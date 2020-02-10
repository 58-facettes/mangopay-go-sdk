package currency

type ISO3 string

func (c ISO3) String() string {
	return string(c)
}

const (
	EUR ISO3 = "EUR"
	GBP ISO3 = "GBP"
	PLN ISO3 = "PLN"
	CHF ISO3 = "CHF"
	NOK ISO3 = "NOK"
	SEK ISO3 = "SEK"
	DKK ISO3 = "DKK"
	USD ISO3 = "USD"
	CAD ISO3 = "CAD"
	AUD ISO3 = "AUD"
	JPY ISO3 = "JPY"
	HKD ISO3 = "HKD"
	AED ISO3 = "AED"
	ZAR ISO3 = "ZAR"
	CZK ISO3 = "CZK"
)
