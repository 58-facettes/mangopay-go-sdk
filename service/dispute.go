package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// Disputes is responsible of all services for Dispute.
type Disputes struct{}

// View is used to view a dispute from it's given ID.
func (Disputes) View(disputeID string) (res *model.Dispute, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "disputes/"+disputeID, nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Update is used to update tag's dispute from it's given ID.
func (Disputes) Update(disputeID, tag string) (res *model.Dispute, err error) {
	payload := struct {
		Tag string `json:"Tag"`
	}{
		Tag: tag,
	}
	_, data, err := newRequestAndExecute(http.MethodPut, "disputes/"+disputeID, payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Close method is used to close a "Dispute", effectively confirming that you do not wish to contest it.
// This action is optional, and once the "ContestDeadlineDate" passes, you will not be able to contest it anyway.
func (Disputes) Close(disputeID string) (res *model.Dispute, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "disputes/"+disputeID+"/close/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Submit a Dispute is used to contest a Dispute.
//
// The "ContestedFunds" must be the same "Currency" as the "DisputedFunds",
// and the amount can be up to and including the "DisputedFunds".
//
// Parameters
// ContestedFunds Money OPTIONAL
// The amount you wish to contest
func (Disputes) Submit(disputeID string, payload *model.Money) (res *model.Dispute, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "disputes/"+disputeID+"/submit/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// ReSubmit a Dispute is used to resubmit a "Dispute" if it is reopened requiring more docs.
// ?? more doc yes indeed ??
func (Disputes) ReSubmit(disputeID string) (res *model.Dispute, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "disputes/"+disputeID+"/submit/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// UserList is listing all User's Disputes.
// ?? the optional params are describe in the body but this is a get method ??
//
// DisputeType DisputeType OPTIONAL
// The type of dispute
//
// Status DisputeStatus OPTIONAL
// The status of this KYC/Dispute document
func (Disputes) UserList(userID string, query ...model.Query) (res []model.Dispute, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("users/"+userID+"/disputes/", query...), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// WalletList is listing all Disputes from a given Wallet.
// ?? the optional params are describe in the body but this is a get method ??
//
// DisputeType DisputeType OPTIONAL
// The type of dispute
//
// Status DisputeStatus OPTIONAL
// The status of this KYC/Dispute document
func (Disputes) WalletList(walletID string, query ...model.Query) (res []model.Dispute, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("wallets/"+walletID+"/disputes/", query...), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// List is listing all the Disputes.
// ?? the optional params are describe in the body but this is a get method ??
// params :
// DisputeType DisputeType OPTIONAL
// The type of dispute
//
// Status DisputeStatus OPTIONAL
// The status of this KYC/Dispute document
func (Disputes) List(query ...model.Query) (res []model.Dispute, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("disputes/", query...), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// PendingSettleList is listing all Disputes that need to be settle.
// In the event you have credit following a Dispute (because you lost, or didn’t contest the full amount),
// you can do a settlement transfer to transfer funds from the original wallet to the credit wallet if you wish
// – this is entirely optional and will depend on your workflow whether you want to impact the original wallet or not.
// The endpoint below provides you with the list of disputes that allow a settlement transfer
// - meaning that you have credit for these disputes and that funds are still available in the original wallet.
func (Disputes) PendingSettleList(query ...model.Query) (res []model.Dispute, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("disputes/pendingsettlement/", query...), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
