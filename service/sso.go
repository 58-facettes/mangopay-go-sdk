package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// SSOS is responsible of all services for the SSO.
type SSOS struct{}

// Create is creating a new SSO from a given SSOCreate payload.
// ?? the uri in the documentation is not correct ??
func (SSOS) Create(payload *model.SSOCreate) (res *model.SSO, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "clients/ssos/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Update is updating an existing SSO from a given SSOUpdate payload and it's given existing ID.
func (SSOS) Update(ssoID string, payload *model.SSOUpdate) (res *model.SSO, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "clients/ssos/"+ssoID+"/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View retreve an sso from it's ID.
func (SSOS) View(ssoID string) (res *model.SSO, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "clients/ssos/"+ssoID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// List is listing all SSO.
func (SSOS) List(query ...model.Query) (res []model.SSO, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("clients/ssos/", query...), nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}

// ExtendInvitation is extending invitation front it's SSO ID.
// ?? not well describe yet ??
func (SSOS) ExtendInvitation(ssoID string) (res *model.SSO, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "clients/ssos/"+ssoID+"/extendinvitation/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// PermissinGroupList allow to list all SSO from it's permission group ID.
// ?? PermissionGroups or permissiongroup in the uri ??
func (SSOS) PermissinGroupList(permissionGroupID string) (res []model.SSO, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "clients/permissiongroups/"+permissionGroupID+"/SSOs/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
