package model

// Transaction is used to describe a transaction.
type Transaction struct {
	// Information about the funds that are being debited.
	DebitedFunds Money `json:"DebitedFunds"`
	// Details about the funds that are being credited (DebitedFunds – Fees = CreditedFunds)/
	CreditedFunds Money `json:"CreditedFunds"`
	// Information about the fees that were taken by the client
	// for this transaction (and were hence transferred to the Client's platform wallet).
	Fees Money `json:"Fees"`
	// The ID of the wallet that was debited.
	DebitedWalletID string `json:"DebitedWalletId"`
	// The ID of the wallet where money will be credited,
	CreditedWalletID string `json:"CreditedWalletId"`
	// A user's ID.
	AuthorID string `json:"AuthorId"`
	// The user ID who is credited (defaults to the owner of the wallet),
	CreditedUserID string `json:"CreditedUserId"`
	// The nature of the transaction,
	Nature TransactionNature `json:"Nature"`
	// The status of the transaction,
	Status TransactionStatus `json:"Status"`
	// When the transaction happened,
	ExecutionDate int64 `json:"ExecutionDate"`
	// The result code,
	ResultCode string `json:"ResultCode"`
	// A verbal explanation of the ResultCode,
	ResultMessage string `json:"ResultMessage"`
	// The type of the transaction,
	Type TransactionType `json:"Type"`
}

// TransactionNature is the nature of a transaction.
type TransactionNature string

const (
	// TransactionNatureRegular is for the transaction nature regular.
	TransactionNatureRegular TransactionNature = "REGULAR"
	// TransactionNatureRepudiation is for the transaction nature repudiation.
	TransactionNatureRepudiation TransactionNature = "REPUDIATION"
	// TransactionNatureRefund is for the transaction nature refund.
	TransactionNatureRefund TransactionNature = "REFUND"
	// TransactionNatureSettlement is for the transaction nature settlement.
	TransactionNatureSettlement TransactionNature = "SETTLEMENT"
)

// TransactionStatus stands for the transaction status.
type TransactionStatus string

const (
	// TransactionStatusCreated is for the transaction status created.
	TransactionStatusCreated TransactionStatus = "CREATED"
	// TransactionStatusSucceeded is for the transaction status succeeded.
	TransactionStatusSucceeded TransactionStatus = "SUCCEEDED"
	// TransactionStatusFailed is for the transaction status failed.
	TransactionStatusFailed TransactionStatus = "FAILED"
)

// TransactionType stands for the transaction type.
type TransactionType string

const (
	// TransactionTypePayin is for the transaction type payin.
	TransactionTypePayin TransactionType = "PAYIN"
	// TransactionTypeTransfer is for the transaction type transfert.
	TransactionTypeTransfer TransactionType = "TRANSFER"
	// TransactionTypePayout is for the transaction type payout.
	TransactionTypePayout TransactionType = "PAYOUT"
)
