package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// Clients is responsible of all services for the Client.
type Clients struct{}

// Update is updating the Client from the given ClientUpdate payload.
func (Clients) Update(payload *model.ClientUpdate) (res *model.Client, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "clients/", payload)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// UploadLogo is updating the logo from a given payload.
func (Clients) UploadLogo(payload *model.ClientLogo) (res *model.Client, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "clients/logo/", payload)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// View is retriving a the Client.
func (Clients) View(query *model.Query) (res *model.Client, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("clients/", query), nil)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}
