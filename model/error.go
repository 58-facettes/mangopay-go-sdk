package model

import (
	"encoding/json"
	"fmt"
)

type ErrorType string

type Error struct {
	ID      string    `json:"Id"`
	Message string    `json:"Message"`
	Type    ErrorType `json:"Type"`
	Date    int64     `json:"Date"`
	Err     error     `json:"error"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%v %v %v %v %v", e.ID, e.Message, e.Type, e.Date, e.Err)
}

// MarshalJSON statify the Marshaler interface.
// It is used to make sure the Date is easy to use in front of the current API impelentation.
func (e Error) MarshalJSON() ([]byte, error) {
	aux := struct {
		ID      string    `json:"Id"`
		Message string    `json:"Message"`
		Type    ErrorType `json:"Type"`
		Date    float64   `json:"Date"`
		Err     error     `json:"error"`
	}{
		ID:      e.ID,
		Message: e.Message,
		Type:    e.Type,
		Date:    float64(e.Date),
		Err:     e.Err,
	}
	return json.Marshal(aux)
}

// UnmarshalJSON satify the Unmarshaler interface,
// this is fixing the date that should be an int64 and not in float64 as it is the currently implemented.
func (e *Error) UnmarshalJSON(b []byte) error {
	aux := struct {
		ID      string    `json:"Id"`
		Message string    `json:"Message"`
		Type    ErrorType `json:"Type"`
		Date    float64   `json:"Date"`
		Err     error     `json:"error"`
	}{}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	e.ID = aux.ID
	e.Message = aux.Message
	e.Type = aux.Type
	e.Date = int64(aux.Date)
	e.Err = aux.Err
	return nil
}
