package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// Users is responsible of all services for the User.
type Users struct{}

// CreateNaturalUser creates a NaturalUser.
func (Users) CreateNaturalUser(payload *model.NaturalUserCreate) (res *model.NaturalUser, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "users/natural/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// UpdateNaturalUser is updating an exinsting NaturalUser.
func (Users) UpdateNaturalUser(userID string, payload *model.NaturalUserUpdate) (res *model.NaturalUser, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "users/natural/"+userID+"/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// CreateLegalUser is creating a LegalUser.
func (Users) CreateLegalUser(payload *model.LegalUserCreate) (res *model.LegalUser, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "users/legal/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// UpdateLegalUser is updating a LegalUser.
func (Users) UpdateLegalUser(userID string, payload *model.LegalUserUpdate) (res *model.LegalUser, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "users/legal/"+userID+"/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View retreve the User fron the given userID.
func (Users) View(userID string) (res *model.User, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "users/"+userID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// List retreve all Users from the cliendID.
func (Users) List(query *model.Query) (res []model.User, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("users/", query), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
