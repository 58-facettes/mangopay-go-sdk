package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// Hooks is responsible of all services for the Hook.
type Hooks struct{}

// Create is creating a new event hook with the given EventType and URL.
func (Hooks) Create(event model.EventType, url string) (res *model.Hook, err error) {
	payload := model.HookCreate{
		EventType: event,
		URL:       url,
	}
	_, data, err := newRequestAndExecute(http.MethodPost, "hooks/", payload)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// View retrive a Hook from it's given ID.
func (Hooks) View(hookID string) (res *model.Hook, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "hooks/"+hookID+"/", nil)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// Update is updating an event hook with the given ID, EventType and URL.
func (Hooks) Update(hookID string, event model.EventType, url string) (res *model.Hook, err error) {
	payload := model.HookCreate{
		EventType: event,
		URL:       url,
	}
	_, data, err := newRequestAndExecute(http.MethodPut, "hooks/"+hookID+"", payload)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}

// List retrive all existing Hooks.
// ?? can we add query params ??
func (Hooks) List(query *model.Query) (res []model.Hook, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, queryURI("hooks/", query), nil)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, &res)
}
