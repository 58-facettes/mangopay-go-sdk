package model

// SettlementTransfer is a transfer that can be used to settle the credit from a repudiation
// following a lost dispute (to impact the balance of the original wallet and settle the credit in your client credit wallet.
type SettlementTransfer struct {
	// Information about the funds that are being debited
	DebitedFunds Money `json:"DebitedFunds"`
	// Details about the funds that are being credited (DebitedFunds – Fees = CreditedFunds)
	CreditedFunds Money `json:"CreditedFunds"`
	// Information about the fees that were taken by the client
	// for this transaction (and were hence transferred to the Client's platform wallet)
	Fees Money `json:"Fees"`
	// The ID of the wallet that was debited
	DebitedWalletID string `json:"DebitedWalletId"`
	// The ID of the wallet where money will be credited
	CreditedWalletID string `json:"CreditedWalletId"`
	// A user's ID
	AuthorID string `json:"AuthorId"`
	// The user ID who is credited (defaults to the owner of the wallet)
	CreditedUserID string `json:"CreditedUserId"`
	// The nature of the transaction
	Nature TransactionNature `json:"Nature"`
	// The status of the transaction
	Status TransactionStatus `json:"Status"`
	// When the transaction happened
	ExecutionDate int64 `json:"ExecutionDate"`
	// The result code
	ResultCode string `json:"ResultCode"`
	// A verbal explanation of the ResultCode
	ResultMessage string `json:"ResultMessage"`
	// The type of the transaction
	Type TransactionType `json:"Type"`
	// The ID of the associated repudiation transaction
	RepudiationID string `json:"RepudiationId"`
}

// SettlementTransferCreate is for creating a SettlementTransfer payload.
// The AuthorId must match the AuthorId of original transaction that was repudiated
type SettlementTransferCreate struct {
	// A user's ID REQUIRED.
	AuthorID string `json:"AuthorId"`
	// Information about the funds that are being debited REQUIRED.
	DebitedFunds Money `json:"DebitedFunds"`
	// Information about the fees that were taken by the client REQUIRED
	// for this transaction (and were hence transferred to the Client's platform wallet).
	Fees Money `json:"Fees"`
}
