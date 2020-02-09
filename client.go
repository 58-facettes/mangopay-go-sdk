package mangopay

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type clientService struct{}

// Update is updating the Client from the given ClientUpdate param.
func (cs *userService) Update(clientID string, param *model.ClientUpdate) (*model.Client, error) {
	_, data, err := newRequestAndExecute(http.MethodPut, baseURL+clientID+"/clients/", param)
	if err != nil {
		return nil, err
	}
	return parseClient(data)
}

// UploadLogo is updating the logo from a given cliendID.
func (cs *userService) UploadLogo(clientID string, param *model.ClientLogo) (*model.Client, error) {
	_, data, err := newRequestAndExecute(http.MethodPut, baseURL+clientID+"/clients/logo/", param)
	if err != nil {
		return nil, err
	}
	return parseClient(data)
}

// View is retriving a Client from the given ClientID.
func (cs *userService) View(clientID string) (*model.Client, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, baseURL+clientID+"/clients/", nil)
	if err != nil {
		return nil, err
	}
	return parseClient(data)
}

func parseClient(data []byte) (*model.Client, error) {
	var client model.Client
	err := json.Unmarshal(data, &client)
	if err != nil {
		return nil, err
	}
	return &client, nil
}
