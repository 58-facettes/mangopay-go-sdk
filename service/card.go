package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// Cards is responsible of all services for Card.
type Cards struct{}

// Register is registering a CardRegistration.
func (Cards) Register(payload *model.CardRegistrationCreate) (res *model.CardRegistration, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "cardregistrations/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// RegisterUpdate is updating an existing CardRegister.
func (Cards) RegisterUpdate(cardRegistrationID string, payload *model.CardRegistrationUpdate) (res *model.CardRegistration, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "CardRegistrations/"+cardRegistrationID+"/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// RegisterUpdateData is a helper that call the RegisterUpdate function without using a model.CardRegistrationUpdate.
func (s *Cards) RegisterUpdateData(cardRegistrationID, data string) (res *model.CardRegistration, err error) {
	return s.RegisterUpdate(
		cardRegistrationID,
		&model.CardRegistrationUpdate{
			RegistrationData: data,
		},
	)
}

// RegisterView allow to view a CardRegistration from the given CardRegistrationID.
func (Cards) RegisterView(cardRegistrationID string) (res *model.CardRegistration, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "CardRegistrations/"+cardRegistrationID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View allow to view a Card from the given CardID.
func (Cards) View(cardID string) (res *model.Card, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "cards/"+cardID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// ListByUser allow to view the Card list of a fiven user from the given userID.
func (Cards) ListByUser(userID string, query *model.Query) (res []model.Card, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("users/"+userID+"/cards/", query), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// ListByFingerprint allow to display a list of Cards from the given fingerprint.
func (Cards) ListByFingerprint(fingerprint string, query *model.Query) (res []model.Card, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("cards/fingerprints/"+fingerprint+"/", query), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// ListTransactionByFingerprint allow to display a list of Transactions from the given fingerprint.
func (Cards) ListTransactionByFingerprint(fingerprint string, query *model.Query) (res []model.Transaction, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("cards/fingerprints/"+fingerprint+"/transactions/", query), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// Desactivate helps you to desactivate a Card from the given cardID.
// Note that once deactivated, a card can't be reactivated afterwards.
func (Cards) Desactivate(cardID string) (res *model.Card, err error) {
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
func (Cards) ListUserByFingerprint(fingerprint string, query *model.Query) (res []model.User, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("cards/fingerprints/"+fingerprint+"/users/", query), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
