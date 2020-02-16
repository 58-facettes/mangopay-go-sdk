package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// PreAuthorizations is responsible of all services for the PreAuthorization.
type PreAuthorizations struct{}

// View retrive a PreAuthorization from it's ID.
func (PreAuthorizations) View(preAuthorizationID string) (res *model.PreAuthorization, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "preauthorizations/"+preAuthorizationID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Cancel is cancelling the pre-authorization.
// The status of the payment after the PreAuthorization.
// You can pass the PaymentStatus from "WAITING" to "CANCELED" should you need/want to.
func (PreAuthorizations) Cancel(preAuthorizationID string, status model.PaymentStatus) (res *model.PreAuthorization, err error) {
	payload := struct {
		PaymentStatus model.PaymentStatus `json:"PaymentStatus"`
	}{
		PaymentStatus: status,
	}
	_, data, err := newRequestAndExecute(http.MethodPut, "preauthorizations/"+preAuthorizationID+"/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// CardList is listing all the pre-authorization of a given cardID.
func (PreAuthorizations) CardList(cardID string) (res []model.PreAuthorization, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "cards/"+cardID+"/preauthorizations/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// UserList is listing all the pre-authorization of a given userID.
//
// ?? this should not be a body payload as describe in the documentation ??
// ?? for now I place it in the query ??
// ?? end ??
//
// Status PreAuthorizationStatus OPTIONAL
// Status of the PreAuthorization
//
// ResultCode string OPTIONAL
// The result code of the transaction (you can filter your transactions list by multiple ResultCode values, each one must be separated by a comma)
//
// PaymentStatus PaymentStatus OPTIONAL
// The status of the payment after the PreAuthorization. You can pass the PaymentStatus from "WAITING" to "CANCELED" should you need/want to
func (PreAuthorizations) UserList(userID string, query ...model.Query) (res []model.PreAuthorization, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("users/"+userID+"/preauthorizations/", query...), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
