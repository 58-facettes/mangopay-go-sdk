package service_test

import (
	"log"

	"github.com/58-facettes/mangopay-go-sdk"
	"github.com/58-facettes/mangopay-go-sdk/model"
	"github.com/58-facettes/mangopay-go-sdk/model/currency"
	_ "github.com/58-facettes/mangopay-go-sdk/service"
)

func ExampleWallets_Create() {
	api := mangopay.NewWithBasicAuth("{{clientID}}", "{{clientPassword}}")
	wallet, _ := api.Wallets.Create(
		&model.WalletCreate{
			Owners:      []string{"my-client-id"},
			Description: "Description of the Wallet",
			Currency:    currency.EUR,
			Tag:         "some tag",
		},
	)
	log.Println(wallet)
}
