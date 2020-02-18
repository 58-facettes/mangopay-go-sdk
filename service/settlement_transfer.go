package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// SettlementTranfers is responsible of all services for the SettlementTranfer.
type SettlementTranfers struct{}

// Create is creating a SettlementTransfer on a given repudiationID and SettlementTransferCreate payload.
func (SettlementTranfers) Create(repuditionID string, payload *model.SettlementTransferCreate) (res *model.SettlementTransfer, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "repudiations/"+repuditionID+"/settlementtranfer/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View retreve a SettlementTranfer from it's ID.
// ?? settelmentID or settlementTranfserID we use it likewise the documentation ??
func (SettlementTranfers) View(settlementID string) (res *model.SettlementTransfer, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "settlements/"+settlementID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}
