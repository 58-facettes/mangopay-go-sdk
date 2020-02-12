package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type ServiceTransfert struct{}

// Create is creating a new Tranfert.
func (ServiceTransfert) Create(param *model.Transfert) (*model.Transfert, error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "transfers/", param)
	if err != nil {
		return nil, err
	}
	return parseTransfert(data)
}

// View is retriving a Transfert form the given transfertID.
func (ServiceTransfert) View(transfertID string) (*model.Transfert, error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "transfers/"+transfertID+"/", nil)
	if err != nil {
		return nil, err
	}
	return parseTransfert(data)
}

func parseTransfert(data []byte) (*model.Transfert, error) {
	var res model.Transfert
	return &res, json.Unmarshal(data, &res)
}
