package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// Mandates is responsible of all services for the Mandate.
type Mandates struct{}

// Create is creating a new mandate from the given payload.
func (Mandates) Create(payload *model.MandateCreate) (res *model.Mandate, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "mandates/directdebit/web/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Cancel is cancelling the Mandate from it's given ID.
func (Mandates) Cancel(mandateID string) (res *model.Mandate, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "mandates/"+mandateID+"/cancel/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// List is listing all mandates.
func (Mandates) List(query ...model.Query) (res []model.Mandate, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "mandates/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// UserList is listing all mandates from a given userID.
func (Mandates) UserList(userID string) (res []model.Mandate, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "users/"+userID+"/mandates/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// BankAccountList list all mandates from a given userID and BankAccountID.
func (Mandates) BankAccountList(userID, bankAccountID string) (res []model.Mandate, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "users/"+userID+"/bankaccounts/"+bankAccountID+"/mandates/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
