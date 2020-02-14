package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

// UserEmoneys is responsible of all services for the User Emoney.
type UserEmoneys struct{}

// View retreve the given User's Emoney from the params UserEmoneyParam the UserID is mandatory.
// If you add a month without the Year this will pick the current year.
func (UserEmoneys) View(userID string, param *model.UserEmoneyParam) (res []model.UserEmoney, err error) {
	url := userID + "/emoney/"
	if param != nil {
		if param.Year != 0 {
			url += fmt.Sprint(param.Year) + "/"
		}
		if param.Month != 0 {
			if param.Year == 0 {
				url += fmt.Sprint(time.Now().Year()) + "/"
			}
			url += fmt.Sprint(param.Month) + "/"
		}
		if len(param.Currency) == 3 {
			url += "?currency=" + param.Currency
		}
	}
	_, data, err := newRequestAndExecute(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	return res, json.Unmarshal(data, &res)
}
