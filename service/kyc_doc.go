package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// KYCs is responsible of all services for the KYC.
type KYCs struct{}

// DocumentCreate is creating a document for the given userID in it given type.
func (KYCs) DocumentCreate(userID string, docType model.DocumentType) (res *model.KYCDocument, err error) {
	payload := struct {
		Type model.DocumentType `json:"Type"`
	}{
		Type: docType,
	}
	_, data, err := newRequestAndExecute(http.MethodPost, "users/"+userID+"/kyc/documents/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// PageCreate is creating a page for the given userID, kycDocumentID (get with the service DocumentCreate) and the given file.
func (KYCs) PageCreate(userID, kycDocumentID string, payload *model.KYCPage) (res *model.KYCDocument, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "users/"+userID+"/kyc/documents/"+kycDocumentID+"/pages/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Submit is submitting for validation the document for verification.
func (KYCs) Submit(userID, kycDocumentID string) (res *model.KYCDocument, err error) {
	payload := struct {
		Status string `json:"Status"`
	}{
		Status: "VALIDATION_ASKED",
	}
	_, data, err := newRequestAndExecute(http.MethodPut, "users/"+userID+"/kyc/documents/"+kycDocumentID+"/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View is retriving the KYC Document from it's ID.
func (KYCs) View(kycDocumentID string) (res *model.KYCDocument, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "kyc/documents/"+kycDocumentID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// ListByUser is retriving all the KYC Document from a userID.
func (KYCs) ListByUser(userID string) (res []model.KYCDocument, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "users/"+userID+"/kyc/documents/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// ListAll is retriving all the KYC Document from the given Query filters params.
// Query parameters :
// The status of this KYC/Dispute document
// Status DocumentStatus OPTIONAL
// The Type of a given KYC Document
// Type KYCDocumentType OPTIONAL
//
// example of params in the filter Query:
// 		{
// 			Filter: map[string]string {
//				"Status": "CREATED, SUCCEEDED",
// 				"Type": "IDENTITY_PROOF,ADDRESS_PROOF",
//			},
// 		}
// ?? the documentation show a body request but with the GET method this is not possible to push a body ??
// so this may be a service with a POST request of a query one.
// ?? end ??
func (KYCs) ListAll(query *model.Query) (res []model.KYCDocument, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("kyc/documents/", query), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
