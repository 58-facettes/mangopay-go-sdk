package mangopay

const (
	APIVersion = "v2.01"
)

const (
	ModeTest       = "test"
	ModeProduction = "production"
)

// Mode is the mode of the SDK calls by defaults this is set to test.
var Mode = ModeTest

var baseURL string

func initBaseURL(ClientID string) {
	switch Mode {
	case ModeProduction:
		baseURL = "https://api.mangopay.com/" + APIVersion + "/" + ClientID + "/"
	default:
		baseURL = "https://api.sandbox.mangopay.com/" + APIVersion + "/" + ClientID + "/"
	}
}
