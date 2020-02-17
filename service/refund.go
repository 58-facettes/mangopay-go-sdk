package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// Refunds is responsible of all services for the Refund.
type Refunds struct{}

// PayInCreate is creating a refund of a given payinID and a RefundCreate payload.
func (Refunds) PayInCreate(payinID string, payload *model.RefundCreate) (res *model.Refund, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "payins/"+payinID+"/refunds/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// TransferCreate is creating a refund of a given tranferID and a authorID.
func (Refunds) TransferCreate(tranferID, authorID string) (res *model.Refund, err error) {
	payload := struct {
		AuthorID string `json:"authorId"`
	}{
		AuthorID: authorID,
	}
	_, data, err := newRequestAndExecute(http.MethodPost, "transfers/"+tranferID+"/refunds/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View retrive a Refund fron it's given ID.
func (Refunds) View(refundID string) (res *model.Refund, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "refunds/"+refundID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// PayoutList is listing all refunds from a payoutID.
func (Refunds) PayoutList(payoutID string, query ...model.Query) (res []model.Refund, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("payouts/"+payoutID+"/refunds/", query...), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
