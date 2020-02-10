package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type ServiceClient struct{}

// Update is updating the Client from the given ClientUpdate param.
func (ServiceClient) Update(param *model.ClientUpdate) (*model.Client, error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "clients/", param)
	if err != nil {
		return nil, err
	}
	return parseClient(data)
}

// UploadLogo is updating the logo from a given param.
func (ServiceClient) UploadLogo(param *model.ClientLogo) (*model.Client, error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "clients/logo/", param)
	if err != nil {
		return nil, err
	}
	return parseClient(data)
}

// View is retriving a the Client.
func (ServiceClient) View() (*model.Client, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "clients/", nil)
	if err != nil {
		return nil, err
	}
	return parseClient(data)
}

func parseClient(data []byte) (*model.Client, error) {
	var client model.Client
	return &client, json.Unmarshal(data, &client)
}
