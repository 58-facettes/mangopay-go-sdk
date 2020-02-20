package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// DisputeDocunents is responsible of all services for the DisputeDocunent.
type DisputeDocunents struct{}

// Create creates a dispute document from the given dispute ID and the kind of ducument.
func (DisputeDocunents) Create(disputeID string, kind model.DisputeDocumentType) (res *model.DisputeDocument, err error) {
	payload := struct {
		Type model.DisputeDocumentType `json:"Type"`
	}{
		Type: kind,
	}
	_, data, err := newRequestAndExecute(http.MethodPost, "disputes/"+disputeID+"/documents/", payload)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// PageCreate is creating a Page from the given disputeID, disputeDocumentID and the file in base64 in DisputeDocumentCreate payload.
func (DisputeDocunents) PageCreate(disputeID, disputeDocumentID string, payload model.DisputeDocumentCreate) (res *model.DisputeDocument, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "disputes/"+disputeID+"/documents/"+disputeDocumentID+"/pages/", payload)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// Submit is sending for validation to the service de referring document.
func (DisputeDocunents) Submit(disputeID, disputeDocumentID string) (res *model.DisputeDocument, err error) {
	payload := struct {
		Status model.DocumentStatus `json:"Status"`
	}{
		Status: model.DocumentStatusCreated,
	}
	_, data, err := newRequestAndExecute(http.MethodPut, "disputes/"+disputeID+"/documents/"+disputeDocumentID+"/pages/", payload)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// View is retriveving the DisputeDocument by it's ID.
func (DisputeDocunents) View(disputeDocumentID string) (res *model.DisputeDocument, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "dispute-documents/"+disputeDocumentID+"/", nil)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// List is listing all the Dispute Documents.
// ?? body param ??
// You can paginate these HTTP GET requests and also add filters
// such as BeforeDate, AfterDate, Type (e.g. REFUND_PROOF) and Status (e.g. VALIDATED).
//
// Get parameters
//
// Status DocumentStatus OPTIONAL
// The status of this KYC/Dispute document
//
// Type DisputeDocumentType OPTIONAL
// The type of the dispute document
func (DisputeDocunents) List(query *model.Query) (res []model.DisputeDocument, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("dispute-documents/", query), nil)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, &res)
}

// ListByDispute is listing all the Document Dispute from a given disputeID.
// ?? body param ??
// Get parameters
// Status DocumentStatus OPTIONAL
// The status of this KYC/Dispute document
//
// Type DisputeDocumentType OPTIONAL
// The type of the dispute document
func (DisputeDocunents) ListByDispute(disputeID string, query *model.Query) (res []model.DisputeDocument, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("dispute/"+disputeID+"/documents/", query), nil)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, &res)
}

// Consult is retriving the DisputeDocument from it's ID.
// ?? Whaaaasaaaat ??
// Consult a dispute document's pages
//
func (DisputeDocunents) Consult(disputeDocumentID string) (res *model.DisputeDocument, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "disputes-documents/"+disputeDocumentID+"/consult/", nil)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}
