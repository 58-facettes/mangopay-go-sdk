package model

type Transaction struct {
	// Information about the funds that are being debited
	DebitedFunds Money `json:"DebitedFunds"`
	// Details about the funds that are being credited (DebitedFunds – Fees = CreditedFunds)
	CreditedFunds Money `json:"CreditedFunds"`
	// Information about the fees that were taken by the client for this transaction (and were hence transferred to the Client's platform wallet)
	Fees Money `json:"Fees"`
	// The ID of the wallet that was debited
	DebitedWalletId string `json:"DebitedWalletId"`
	// The ID of the wallet where money will be credited
	CreditedWalletId string `json:"CreditedWalletId"`
	// A user's ID
	AuthorId string `json:"AuthorId"`
	// The user ID who is credited (defaults to the owner of the wallet)
	CreditedUserId string `json:"CreditedUserId"`
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
}

type TransactionNature string

const (
	TransactionNatureRegular     TransactionNature = "REGULAR"
	TransactionNatureRepudiation TransactionNature = "REPUDIATION"
	TransactionNatureRefund      TransactionNature = "REFUND"
	TransactionNatureSettlement  TransactionNature = "SETTLEMENT"
)

type TransactionStatus string

const (
	TransactionStatusCreated   TransactionStatus = "CREATED"
	TransactionStatusSucceeded TransactionStatus = "SUCCEEDED"
	TransactionStatusFailed    TransactionStatus = "FAILED"
)

type TransactionType string

const (
	TransactionTypePayin    TransactionType = "PAYIN"
	TransactionTypeTransfer TransactionType = "TRANSFER"
	TransactionTypePayout   TransactionType = "PAYOUT"
)
