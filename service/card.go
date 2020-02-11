package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type ServiceCard struct{}

// Register is registering a CardRegistration.
func (ServiceCard) Register(param *model.CardRegistrationCreate) (*model.CardRegistration, error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "cardregistrations/", param)
	if err != nil {
		return nil, err
	}
	return parseCardRegistration(data)
}

// RegisterUpdate is updating an existing CardRegister.
func (ServiceCard) RegisterUpdate(cardRegistrationID string, param *model.CardRegistrationUpdate) (*model.CardRegistration, error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "CardRegistrations/"+cardRegistrationID+"/", param)
	if err != nil {
		return nil, err
	}
	return parseCardRegistration(data)
}

// RegisterUpdateData is a helper that call the RegisterUpdate function without using a model.CardRegistrationUpdate.
func (s *ServiceCard) RegisterUpdateData(cardRegistrationID, data string) (*model.CardRegistration, error) {
	return s.RegisterUpdate(
		cardRegistrationID,
		&model.CardRegistrationUpdate{
			RegistrationData: data,
		},
	)
}

// RegisterView allow to view a CardRegistration from the given CardRegistrationID.
func (ServiceCard) RegisterView(cardRegistrationID string) (*model.CardRegistration, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "CardRegistrations/"+cardRegistrationID+"/", nil)
	if err != nil {
		return nil, err
	}
	return parseCardRegistration(data)
}

func parseCardRegistration(data []byte) (*model.CardRegistration, error) {
	var cr model.CardRegistration
	return &cr, json.Unmarshal(data, &cr)
}

// View allow to view a Card from the given CardID.
func (ServiceCard) View(cardID string) (*model.Card, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "cards/"+cardID+"/", nil)
	if err != nil {
		return nil, err
	}
	return parseCard(data)
}

// ListByUser allow to view the Card list of a fiven user from the given userID.
func (ServiceCard) ListByUser(userID string) ([]model.Card, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "users/"+userID+"/cards/", nil)
	if err != nil {
		return nil, err
	}
	return parseCardCollection(data)
}

// ListByFingerprint allow to display a list of Cards from the given fingerprint.
func (ServiceCard) ListByFingerprint(fingerprint string) ([]model.Card, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "cards/fingerprints/"+fingerprint+"/", nil)
	if err != nil {
		return nil, err
	}
	return parseCardCollection(data)
}

// ListTransactionByFingerprint allow to display a list of Transactions from the given fingerprint.
func (ServiceCard) ListTransactionByFingerprint(fingerprint string) ([]model.Transaction, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "cards/fingerprints/"+fingerprint+"/transactions/", nil)
	if err != nil {
		return nil, err
	}
	return parseTransactionCollection(data)
}

// Desactivate helps you to desactivate a Card from the given cardID.
// Note that once deactivated, a card can't be reactivated afterwards.
func (ServiceCard) Desactivate(cardID string) (*model.Card, error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "cards/"+cardID+"/",
		&model.CardDeactivate{
			Active: false,
		})
	if err != nil {
		return nil, err
	}
	return parseCard(data)
}

func parseCard(data []byte) (*model.Card, error) {
	var c model.Card
	return &c, json.Unmarshal(data, &c)
}

func parseCardCollection(data []byte) ([]model.Card, error) {
	var list []model.Card
	return list, json.Unmarshal(data, &list)
}

// ListUserByFingerprint allow to display a list of Users from the given fingerprint.
func (ServiceCard) ListUserByFingerprint(fingerprint string) ([]model.User, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "cards/fingerprints/"+fingerprint+"/users/", nil)
	if err != nil {
		return nil, err
	}
	return parseUserCollection(data)
}
