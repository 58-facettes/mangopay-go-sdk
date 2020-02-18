package service

import (
	"encoding/json"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// Repudiations is responsible of all services for the Repudiation.
type Repudiations struct{}

// View retreve a Repudiation from it's ID.
func (Repudiations) View(repudiationID string) (res *model.Repudiation, err error) {
	_, data, err := newRequestAndExecute(http.MethodGet, "repudiations/"+repudiationID+"/", nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, res)
}
