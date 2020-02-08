package mangopay

import (
	"time"
)

type Mangopay struct {
	isBasicAuth bool
	clientID    string
	apiKey      string
	Client      clientService
}

// NewWithBasicAuth sends a new Mangonpay client with Basic Auth.
func NewWithBasicAuth(CliendID, APIkey string) *Mangopay {
	m := Mangopay{
		isBasicAuth: true,
		clientID:    CliendID,
		apiKey:      APIkey,
	}
	initBasicAuth(CliendID, APIkey)
	initBaseURL(CliendID)
	return &m
}

func String(str string) *string {
	return &str
}

func Time(t time.Time) *int64 {
	res := t.UTC().Unix()
	return &res
}

func Int(i int) *int {
	return &i
}
