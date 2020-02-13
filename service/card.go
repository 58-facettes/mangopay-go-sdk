package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type ServiceCard struct{}

// Register is registering a CardRegistration.
func (ServiceCard) Register(payload *model.CardRegistrationCreate) (res *model.CardRegistration, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "cardregistrations/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// RegisterUpdate is updating an existing CardRegister.
func (ServiceCard) RegisterUpdate(cardRegistrationID string, payload *model.CardRegistrationUpdate) (res *model.CardRegistration, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "CardRegistrations/"+cardRegistrationID+"/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// RegisterUpdateData is a helper that call the RegisterUpdate function without using a model.CardRegistrationUpdate.
func (s *ServiceCard) RegisterUpdateData(cardRegistrationID, data string) (res *model.CardRegistration, err error) {
	return s.RegisterUpdate(
		cardRegistrationID,
		&model.CardRegistrationUpdate{
			RegistrationData: data,
		},
	)
}

// RegisterView allow to view a CardRegistration from the given CardRegistrationID.
func (ServiceCard) RegisterView(cardRegistrationID string) (res *model.CardRegistration, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "CardRegistrations/"+cardRegistrationID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View allow to view a Card from the given CardID.
func (ServiceCard) View(cardID string) (res *model.Card, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "cards/"+cardID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// ListByUser allow to view the Card list of a fiven user from the given userID.
func (ServiceCard) ListByUser(userID string) (res []model.Card, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "users/"+userID+"/cards/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// ListByFingerprint allow to display a list of Cards from the given fingerprint.
func (ServiceCard) ListByFingerprint(fingerprint string) (res []model.Card, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "cards/fingerprints/"+fingerprint+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// ListTransactionByFingerprint allow to display a list of Transactions from the given fingerprint.
func (ServiceCard) ListTransactionByFingerprint(fingerprint string) (res []model.Transaction, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "cards/fingerprints/"+fingerprint+"/transactions/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Desactivate helps you to desactivate a Card from the given cardID.
// Note that once deactivated, a card can't be reactivated afterwards.
func (ServiceCard) Desactivate(cardID string) (res *model.Card, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "cards/"+cardID+"/",
		&model.CardDeactivate{
			Active: false,
		})
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// ListUserByFingerprint allow to display a list of Users from the given fingerprint.
func (ServiceCard) ListUserByFingerprint(fingerprint string) (res []model.User, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "cards/fingerprints/"+fingerprint+"/users/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}
