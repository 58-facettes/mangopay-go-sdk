package service_test

import (
	"os"
	"testing"
	"time"

	mangopay "github.com/58-facettes/mangopay-go-sdk"
	"github.com/58-facettes/mangopay-go-sdk/model"
	"github.com/davecgh/go-spew/spew"
)

var (
	clientID string
	tokenAPI string
)

func init() {
	clientID = os.Getenv("MANGOPAY_CLIENT_ID")
	tokenAPI = os.Getenv("MANGOPAY_API_TOKEN")
	if len(clientID) == 0 {
		clientID = "test_asd"
	}
	if len(tokenAPI) == 0 {
		tokenAPI = "00000"
	}
}
func TestNewWithBasicAuth(t *testing.T) {
	api := mangopay.NewWithBasicAuth(clientID, tokenAPI)
	param := &model.NaturalUserCreate{
		FirstName:          "Martin",
		LastName:           "Gallager",
		Birthday:           model.Time(time.Now()),
		Nationality:        "FR",
		CountryOfResidence: "FR",
	}
	u, err := api.Users.CreateNaturalUser(param)
	t.Log("ner user is created", spew.Sdump(u))
	t.Log("error is", err)
}
