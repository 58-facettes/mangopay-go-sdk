package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// Idempotencies is responsible of all services for the Idempotency.
type Idempotencies struct{}

// View retreve the payload from the idempotencyKEY.
func (Idempotencies) View(idempotencyKEY string) (res *model.IempotencyResponse, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "responses/"+idempotencyKEY+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}
