package model

// Transfert is the payload for tranferts.
type Transfert struct {
	ID           string `json:"Id"`
	Tag          string `json:"Tag"`
	CreationDate int64  `json:"CreationDate"`
	TansfertCreate
}

// TansfertCreate is used to create a new tranfert.
type TansfertCreate struct {
	// A user's ID REQUIRED.
	AuthorID string `json:"AuthorId"`
	// The user ID who is credited (defaults to the owner of the wallet) OPTIONAL.
	CreditedUserID string `json:"CreditedUserId,omitempty"`
	// Information about the funds that are being debited REQUIRED.
	DebitedFunds Money `json:"DebitedFunds"`
	// Information about the fees that were taken by the client
	// for this transaction (and were hence transferred to the Client's platform wallet) REQUIRED.
	Fees Money `json:"Fees"`
	// The ID of the wallet that was debited REQUIRED.
	DebitedWalletID string `json:"DebitedWalletId"`
	// The ID of the wallet where money will be credited REQUIRED.
	CreditedWalletID string `json:"CreditedWalletId"`
}
