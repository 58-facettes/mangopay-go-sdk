package mangopay

import "encoding/base64"

var basicAuth string

// initBasicAuth set the basicAuth variable that will be used with "Authorization" in the header.
func initBasicAuth(user, password string) {
	basicAuth = "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password))
}
