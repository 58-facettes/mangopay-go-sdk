package mangopay

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type userService struct{}

// CreateNaturalUser creates a NaturalUser.
func (userService) CreateNaturalUser(param *model.NaturalUserCreate) (*model.NaturalUser, error) {
	_, data, err := newRequestAndExecute(http.MethodPost, baseURL+"users/natural/", param)
	if err != nil {
		return nil, err
	}
	return parseNaturalUser(data)
}

// UpdateNaturalUser is updating an exinsting NaturalUser.
func (userService) UpdateNaturalUser(userID string, param *model.NaturalUserUpdate) (*model.NaturalUser, error) {
	_, data, err := newRequestAndExecute(http.MethodPut, baseURL+"users/natural/"+userID+"/", param)
	if err != nil {
		return nil, err
	}
	return parseNaturalUser(data)
}

func parseNaturalUser(data []byte) (*model.NaturalUser, error) {
	var naturalUser model.NaturalUser
	err := json.Unmarshal(data, &naturalUser)
	if err != nil {
		return nil, err
	}
	return &naturalUser, nil
}

// CreateLegalUser is creating a LegalUser.
func (userService) CreateLegalUser(param *model.LegalUserCreate) (*model.LegalUser, error) {
	_, data, err := newRequestAndExecute(http.MethodPost, baseURL+"users/legal/", param)
	if err != nil {
		return nil, err
	}
	return parseLegalUser(data)
}

// UpdateLegalUser is updating a LegalUser.
func (userService) UpdateLegalUser(userID string, param *model.LegalUserUpdate) (*model.LegalUser, error) {
	_, data, err := newRequestAndExecute(http.MethodPut, baseURL+"users/legal/"+userID+"/", param)
	if err != nil {
		return nil, err
	}
	return parseLegalUser(data)
}

func parseLegalUser(data []byte) (*model.LegalUser, error) {
	var legalUser model.LegalUser
	err := json.Unmarshal(data, &legalUser)
	if err != nil {
		return nil, err
	}
	return &legalUser, nil
}

// ViewUser retreve the User fron the given userID.
func (userService) ViewUser(userID string) (*model.User, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, baseURL+"users/"+userID+"/", nil)
	if err != nil {
		return nil, err
	}
	return parseUser(data)
}

// ListAllUsers retreve all Users from the cliendID.
func (userService) ListAllUsers() ([]model.User, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, baseURL+"users/", nil)
	if err != nil {
		return nil, err
	}
	return parseUserCollection(data)
}

func parseUser(data []byte) (*model.User, error) {
	var user model.User
	return &user, json.Unmarshal(data, &user)
}

func parseUserCollection(data []byte) ([]model.User, error) {
	var us []model.User
	return us, json.Unmarshal(data, &us)
}
