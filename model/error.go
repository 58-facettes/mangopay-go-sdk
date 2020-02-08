package model

import "fmt"

type ErrorType string

type Error struct {
	ID      string    `json:"Id"`
	Message string    `json:"Message"`
	Type    ErrorType `json:"Type"`
	Date    float64   `json:"Date"`
	Err     error     `json:"error"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%v %v %v %v %v", e.ID, e.Message, e.Type, e.Date, e.Err)
}
