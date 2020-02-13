package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type ServicePayOut struct{}

// Create is creating a payout form the given PayOutCreate payload.
func (ServicePayOut) Create(payload *model.PayOutCreate) (res *model.PayOut, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "payouts/bankwire/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View display a payout front the payoutID.
func (ServicePayOut) View(payoutID string) (res *model.PayOut, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "payouts/"+payoutID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}
