package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// UBOs is responsible of all services for the UBO.
type UBOs struct{}

// DeclarationCreate creates an UBO Declaration.
func (UBOs) DeclarationCreate(userID string) (res *model.UBO, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "users/"+userID+"/kyc/ubodeclarations/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Create creates an UBO.
func (UBOs) Create(userID, declarationID string, payload *model.UBOCreate) (res *model.UBO, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "users/"+userID+"/kyc/ubodeclarations/"+declarationID+"/ubos/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Update updates an UBO.
func (UBOs) Update(userID, declarationID, uboID string, payload *model.UBOUpdate) (res *model.UBO, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "users/"+userID+"/kyc/ubodeclarations/"+declarationID+"/ubos/"+uboID+"/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// DeclarationSubmit is submitting the declaration.
// ?? not sure this is sending back an UBO but ??
// this can send back otherwise an UBODeclaration
// it is not mentioned in the documentation.
// ?? end ??
func (UBOs) DeclarationSubmit(userID, declarationID string) (res *model.UBO, err error) {
	payload := struct {
		Status model.UBODeclarationStatus `json:"Status"`
	}{
		Status: model.UBODeclarationStatusValidationAsked,
	}
	_, data, err := newRequestAndExecute(http.MethodPut, "users/"+userID+"/kyc/ubodeclarations/"+declarationID+"/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// DeclarationView is for retreving the declaration.
// ?? not sure this is sending back an UBO but ??
// this can send back otherwise an UBODeclaration
// it is not mentioned in the documentation.
// ?? end ??
func (UBOs) DeclarationView(declarationID string) (res *model.UBO, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "/kyc/ubodeclarations/"+declarationID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View is for retreving the UBO.
func (UBOs) View(userID, declarationID, uboID string) (res *model.UBO, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "users/"+userID+"/kyc/ubodeclarations/"+declarationID+"/ubos/"+uboID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// List is listing all the userID this can be combined with query.
// The column to sort against with `Sort` key in `model.Query.Filer`
// this can have ASC, DESC as a value.
func (UBOs) List(userID string, query *model.Query) (res *model.UBO, err error) {
	uri := queryURI("/users/"+userID+"/kyc/ubodeclarations/", query)
	_, data, err := newRequestAndExecute(http.MethodGet, uri, nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}
