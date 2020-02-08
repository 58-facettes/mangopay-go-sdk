package mangopay

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type clientService struct{}

func (cs *clientService) CreateNaturalUser(param *model.NaturalUserCreate) (*model.NaturalUser, error) {
	_, data, err := newRequestAndExecute(http.MethodPost, baseURL+"users/natural/", param)
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
