package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type ServicePayin struct{}

func (ServicePayin) CardWebCreate(param *model.CardWebCreate) (res *model.CardWeb, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "payins/card/web/", param)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

func (ServicePayin) CardDirectCreate(param *model.CardDirectCreate) (res *model.CardDirect, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "payins/card/direct/", param)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

func (ServicePayin) CardPreAuthorizedCreate(param *model.CardPreAuthorizedCreate) (res *model.CardPreAuthorized, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "payins/preauthorized/direct/", param)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

func (ServicePayin) BankwireDirect(param *model.BankwireDirectCreate) (res *model.BankwireDirect, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "payins/bankwire/direct/", param)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// ?? model.DirectDebitWebCreate don't exists YET ??
// func (ServicePayin) DirectDebitWeb(param *model.DirectDebitWebCreate) (res *model.DirectDebitWeb, err error) {
// 	_, data, err := newRequestAndExecute(http.MethodPost, "payins/directdebit/web/", param)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, json.Unmarshal(data, res)
// }

func (ServicePayin) DirectDebitDirect(param *model.DirectDebitDirectCreate) (res *model.DirectDebitDirect, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "payins/directdebit/direct/", param)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// ?? not sure this will retrun a real model.DirectDebitDirect ??
func (ServicePayin) BankwireToClientCreditWalletCreate(param *model.BankwireToClientCreditWalletCreate) (res *model.DirectDebitDirect, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "clients/payins/bankwire/direct/", param)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// ?? this is returning for the moment an interface - what type of data is it returning exactly ??
func (ServicePayin) PayInWebExtended(param *model.WebExtended) (res interface{}, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "?", param)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// ?? We don't have yet the type of struct out ??
// Note that this endpoint is only available for PayIn:
// through Web interface
// when customer has provided some Card infos
func (ServicePayin) ViewCardDetails(payInID string) (res interface{}, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "payins/card/web/"+payInID+"/extended/", nil)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

func (ServicePayin) View(payInID string) (res *model.Payin, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "payins/"+payInID, nil)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}
