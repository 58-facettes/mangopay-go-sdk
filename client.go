package mangopay

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type clientService struct{}

// Update is updating the Client from the given ClientUpdate param.
func (cs *userService) Update(param *model.ClientUpdate) (*model.Client, error) {
	_, data, err := newRequestAndExecute(http.MethodPut, baseURL+"clients/", param)
	if err != nil {
		return nil, err
	}
	return parseClient(data)
}

// UploadLogo is updating the logo from a given param.
func (cs *userService) UploadLogo(param *model.ClientLogo) (*model.Client, error) {
	_, data, err := newRequestAndExecute(http.MethodPut, baseURL+"clients/logo/", param)
	if err != nil {
		return nil, err
	}
	return parseClient(data)
}

// View is retriving a the Client.
func (cs *userService) View() (*model.Client, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, baseURL+"clients/", nil)
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
