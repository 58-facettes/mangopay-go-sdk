package mangopay_test

import "github.com/58-facettes/mangopay-go-sdk"

func ExampleNewWithBasicAuth() {
	api := mangopay.NewWithBasicAuth("client-id", "client-password")
	api.Users.View("1234")
}
