package model

type Transfert struct {
	ID           string `json:"Id"`
	Tag          string `json:"Tag"`
	CreationDate int64  "CreationDate"
	TansfertCreate
}

type TansfertCreate struct {
	// A user's ID REQUIRED.
	AuthorId string `json:"AuthorId"`
	// The user ID who is credited (defaults to the owner of the wallet) OPTIONAL.
	CreditedUserId string `json:"CreditedUserId,omitempty"`
	// Information about the funds that are being debited REQUIRED.
	DebitedFunds Money `json:"DebitedFunds"`
	// Information about the fees that were taken by the client
	// for this transaction (and were hence transferred to the Client's platform wallet) REQUIRED.
	Fees Money `json:"Fees"`
	// The ID of the wallet that was debited REQUIRED.
	DebitedWalletId string `json:"DebitedWalletId"`
	// The ID of the wallet where money will be credited REQUIRED.
	CreditedWalletId string `json:"CreditedWalletId"`
}
