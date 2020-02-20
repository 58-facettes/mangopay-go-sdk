package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// PermissionGoups is responsible of all services for the PermissionGroup.
type PermissionGoups struct{}

// Create is creating a new PermissionGroup.
func (PermissionGoups) Create(payload *model.PermissionGroupCreate) (res *model.PermissionGroup, err error) {
	_, data, err := newRequestAndExecute(http.MethodPost, "clients/permissiongroups/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// Update is updating an existing PermissionGroup.
func (PermissionGoups) Update(permissionGroupID string, payload *model.PermissionGroupUpdate) (res *model.PermissionGroup, err error) {
	_, data, err := newRequestAndExecute(http.MethodPut, "clients/permissiongroups/"+permissionGroupID+"/", payload)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// View retreve a PermissionGroup from it's ID.
func (PermissionGoups) View(permissionGroupID string) (res *model.PermissionGroup, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "clients/permissiongroups/"+permissionGroupID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}

// List retreve all PermissionGroups.
func (PermissionGoups) List(query *model.Query) (res []model.PermissionGroup, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "clients/permissiongroups/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
