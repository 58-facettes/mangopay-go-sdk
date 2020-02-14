package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// PayIns is responsible of all services for the Payin.
type PayIns struct{}

// CardWebCreate is creating a CardWeb.
func (PayIns) CardWebCreate(payload *model.CardWebCreate) (res *model.CardWeb, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "payins/card/web/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// CardDirectCreate creates a CardDirect.
func (PayIns) CardDirectCreate(payload *model.CardDirectCreate) (res *model.CardDirect, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "payins/card/direct/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// CardPreAuthorizedCreate is used for pre-authorization.
func (PayIns) CardPreAuthorizedCreate(payload *model.CardPreAuthorizedCreate) (res *model.CardPreAuthorized, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "payins/preauthorized/direct/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// BankwireDirect is use to create a direct bankwire.
func (PayIns) BankwireDirect(payload *model.BankwireDirectCreate) (res *model.BankwireDirect, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "payins/bankwire/direct/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// ?? model.DirectDebitWebCreate don't exists YET ??
// func (PayIns) DirectDebitWeb(payload *model.DirectDebitWebCreate) (res *model.DirectDebitWeb, err error) {
// 	_, data, err := newRequestAndExecute(http.MethodPost, "payins/directdebit/web/", payload)
// 	if err != nil {
// 		return
// 	}
// 	return res, json.Unmarshal(data, res)
// }

// DirectDebitDirect is for creating a new direct debit.
func (PayIns) DirectDebitDirect(payload *model.DirectDebitDirectCreate) (res *model.DirectDebitDirect, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "payins/directdebit/direct/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// BankwireToClientCreditWalletCreate is for creating a bankwire to a client wallet.
// ?? not sure this will retrun a real model.DirectDebitDirect ??
func (PayIns) BankwireToClientCreditWalletCreate(payload *model.BankwireToClientCreditWalletCreate) (res *model.DirectDebitDirect, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "clients/payins/bankwire/direct/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// PayInWebExtended is for somthing I have no clue.
// ?? this is returning for the moment an interface - what type of data is it returning exactly ??
func (PayIns) PayInWebExtended(payload *model.WebExtended) (res interface{}, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "?", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// ViewCardDetails is used to get a Card detail in a PayIn from it's ID.
// ?? We don't have yet the type of struct out ??
// Note that this endpoint is only available for PayIn:
// through Web interface
// when customer has provided some Card infos
func (PayIns) ViewCardDetails(payInID string) (res interface{}, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "payins/card/web/"+payInID+"/extended/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View is used to view a Payin from it's ID.
func (PayIns) View(payInID string) (res *model.Payin, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "payins/"+payInID, nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}
