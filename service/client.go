package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type ServiceClient struct{}

// Update is updating the Client from the given ClientUpdate payload.
func (ServiceClient) Update(payload *model.ClientUpdate) (res *model.Client, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "clients/", payload)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// UploadLogo is updating the logo from a given payload.
func (ServiceClient) UploadLogo(payload *model.ClientLogo) (res *model.Client, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "clients/logo/", payload)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// View is retriving a the Client.
func (ServiceClient) View(query ...model.Query) (res *model.Client, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, addQuery("clients/", query...), nil)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}
