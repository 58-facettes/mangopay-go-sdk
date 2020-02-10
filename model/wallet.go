package model

// Wallet is an object in which PayIns and Transfers from users are stored in order to collect money.
// You can pay into a Wallet, withdraw funds from a wallet or transfer funds from a Wallet to another Wallet.
type Wallet struct {
	// An array of userIDs of who own's the wallet. For now, you only can set up a unique owner.
	Owners []string `json:"Owners"`
	// The current balance of the wallet
	Balance Money `json:"Balance"`
	// The type of funds in the wallet
	FundsType FundsType `json:"FundsType"`
	// A desciption of the wallet
	Description string `json:"Description"`
	// The currency - should be ISO_4217 format
	Currency CurrencyISO3 `json:"Currency"`
}

type WalletCreate struct {
	// An array of userIDs of who own's the wallet. For now, you only can set up a unique owner. (REQUIRED)
	Owners []string `json:"Owners"`
	// A desciption of the wallet (REQUIRED)
	Description string `json:"Description"`
	// The currency - should be ISO_4217 format (REQUIRED).
	Currency CurrencyISO3 `json:"Currency"`
	// Custom data that you can add to this item (OPTIONAL).
	Tag *string `json:"Tag"`
}

type WalletUpdate struct {
	// A desciption of the wallet (OPTIONAL)
	Description *string `json:"Description"`
}
