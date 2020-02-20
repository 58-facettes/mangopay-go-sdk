package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// Events is responsible of all services for the Event.
type Events struct{}

// List is listing all events.
// it has an EventType optional parameter
// ?? in the documentation this is shown as a body param ??
// ?? as a get method we can't use a body ??
// ?? is it a query param ??
// ?? end ??
func (Events) List(event model.EventType, query *model.Query) (res []model.Event, err error) {
	uri := queryURI("events/", query)
	if len(event) != 0 {
		uri += "?event_type=" + string(event)
	}
	_, data, err := newRequestAndExecute(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, &res)
}
