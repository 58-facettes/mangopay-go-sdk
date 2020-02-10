package service_test

import (
	"log"

	"github.com/58-facettes/mangopay-go-sdk"
	"github.com/58-facettes/mangopay-go-sdk/currency"
	"github.com/58-facettes/mangopay-go-sdk/model"
)

func ExampleB_WalletCreate() {
	api := mangopay.NewWithBasicAuth("{{clientID}}", "{{clientPassword}}")
	wallet, _ := api.Wallets.Create(
		&model.WalletCreate{
			Owners:      []string{"my_client_id"},
			Description: "Description of the Wallet",
			Currency:    currency.EUR,
			Tag:         model.String("some tag"),
		},
	)
	log.Println(wallet)
}
