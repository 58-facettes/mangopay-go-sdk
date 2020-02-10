package mangopay

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type userService struct{}

func (cs *userService) CreateNaturalUser(param *model.NaturalUserCreate) (*model.NaturalUser, error) {
	_, data, err := newRequestAndExecute(http.MethodPost, baseURL+"users/natural/", param)
	if err != nil {
		return nil, err
	}
	return parseNaturalUser(data)
}

func (cs *userService) UpdateNaturalUser(userID string, param *model.NaturalUserUpdate) (*model.NaturalUser, error) {
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

func (cs *userService) CreateLegalUser(param *model.LegalUserCreate) (*model.LegalUser, error) {
	_, data, err := newRequestAndExecute(http.MethodPost, baseURL+"users/legal/", param)
	if err != nil {
		return nil, err
	}
	return parseLegalUser(data)
}

func (cs *userService) UpdateLegalUser(userID string, param *model.LegalUserUpdate) (*model.LegalUser, error) {
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
