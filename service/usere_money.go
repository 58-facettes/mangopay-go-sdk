package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

type ServiceUserEmoney struct{}

// View retreve the given User's Emoney from the params UserEmoneyParam the UserID is mandatory.
// If you add a month without the Year this will pick the current year.
func (ServiceUserEmoney) View(userID string, param *model.UserEmoneyParam) (res []model.UserEmoney, err error) {
	url := userID + "/emoney/"
	if param != nil {
		if param.Year != nil {
			url += fmt.Sprint(*param.Year) + "/"
		}
		if param.Month != nil {
			if param.Year == nil {
				url += fmt.Sprint(time.Now().Year()) + "/"
			}
			url += fmt.Sprint(*param.Month) + "/"
		}
		if param.Currency != nil {
			url += "?currency=" + *param.Currency
		}
	}
	_, data, err := newRequestAndExecute(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return res, json.Unmarshal(data, res)
}
