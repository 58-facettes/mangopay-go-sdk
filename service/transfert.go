package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// Transferts is responsible of all services for the Transfert.
type Transferts struct{}

// Create is creating a new Tranfert.
func (Transferts) Create(payload *model.Transfert) (res *model.Transfert, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "transfers/", payload)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// View is retriving a Transfert form the given transfertID.
func (Transferts) View(transfertID string) (res *model.Transfert, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "transfers/"+transfertID+"/", nil)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}
