/*
Package Mangopay provides services to communicate with the official API.

A basic usage would be to create a Basic Auth connexion.

	api := mangopay.NewWithBasicAuth(cliendID, clientPassword)
	cli, _ := api.Clients.List(nil)
	log.Print(cli) // this will display your current Client information.
*/
package mangopay
