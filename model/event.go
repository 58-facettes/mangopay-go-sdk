package model

type Event struct {
	// The ID of whatever the event is
	ResourceID string `json:"ResourceId"`
	// When the event happened
	Date int64 `json:"Date"`
	// The event type
	EventType EventType `json:"EventType"`
}
