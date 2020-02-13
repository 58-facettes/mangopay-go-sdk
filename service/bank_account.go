package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type ServiceBankAccount struct{}

// View display the BankAccount informtion from the given bankAccountID.
func (ServiceBankAccount) View(bankAccountID string) (res *model.BankAccount, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "bankaccounts/"+bankAccountID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// ViewFromUser display all the BankAccounts informtion from the given userID.
func (ServiceBankAccount) ViewFromUser(userID string) (res []model.BankAccount, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "users/"+userID+"/bankaccounts/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// CreateIBAN is creating a new IBAN for a given userID.
func (ServiceBankAccount) CreateIBAN(userID string, payload *model.BankAccountIBANCreate) (res model.BankAccountIBAN, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "users/"+userID+"/bankaccounts/iban/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// CreateUS is creating a new BankAccount in US for the given userID.
// ?? not 100% positive on the return result in the documentation there is no BankAccountUS describe ??
func (ServiceBankAccount) CreateUS(userID string, payload *model.BankAccountUSCreate) (res model.BankAccount, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "users/"+userID+"/bankaccounts/us/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// CreateCA is creating a new BankAccount in CA for the given userID.
func (ServiceBankAccount) CreateCA(userID string, payload *model.BankAccountCACreate) (res model.BankAccountCA, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "users/"+userID+"/bankaccounts/ca/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// CreateGB is creating a new BankAccount in GB for the given userID.
func (ServiceBankAccount) CreateGB(userID string, payload *model.BankAccountGBCreate) (res model.BankAccountGB, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "users/"+userID+"/bankaccounts/gb/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// CreateOther is creating a new BankAccount in other bank type for the given userID.
func (ServiceBankAccount) CreateOther(userID string, payload *model.BankAccountOtherCreate) (res model.BankAccountOther, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "users/"+userID+"/bankaccounts/other/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Desactivate is desactivating a BankAccount from the given userID and bankAccountID.
// ?? not sure what king of payload this is going to send back ??
func (ServiceBankAccount) Desactivate(userID, bankAccountID string) (res model.BankAccount, err error) {
	param := struct {
		Active bool `json:"Active"`
	}{
		false,
	}
	_, data, err := newRequestAndExecute(http.MethodPut, "users/"+userID+"/bankaccounts/"+bankAccountID+"/", param)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}
