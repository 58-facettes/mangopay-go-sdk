package model

// Repudiation is created when a user has a requested a chargeback
// for a PayIn and the bank has withdrawn the funds from us automatically – they are always linked to a dispute.
type Repudiation struct {
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
	ResultMessage string
	// The type of the transaction
	Type TransactionType `json:"Type"`
	// The initial transaction ID
	InitialTransactionID string `json:"InitialTransactionId"`
	// The initial transaction type
	InitialTransactionType TransactionType `json:"InitialTransactionType"`
}
