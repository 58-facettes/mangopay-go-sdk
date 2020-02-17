package model

// Refund represent the refund payload.
type Refund struct {
	// Information about the funds that are being debited.
	DebitedFunds Money `json:"DebitedFunds"`
	// Details about the funds that are being credited (DebitedFunds – Fees = CreditedFunds).
	CreditedFunds Money `json:"CreditedFunds"`
	// Information about the fees that were taken by the client
	// for this transaction (and were hence transferred to the Client's platform wallet).
	Fees Money `json:"Fees"`
	// The ID of the wallet that was debited.
	DebitedWalletID string `json:"DebitedWalletId"`
	// The ID of the wallet where money will be credited.
	CreditedWalletID string `json:"CreditedWalletId"`
	// A user's ID.
	AuthorID string `json:"AuthorId"`
	// The user ID who is credited (defaults to the owner of the wallet).
	CreditedUserID string `json:"CreditedUserId"`
	// The nature of the transaction.
	Nature TransactionNature `json:"Nature"`
	// The status of the transaction.
	Status TransactionStatus `json:"Status"`
	// When the transaction happened.
	ExecutionDate int64 `json:"ExecutionDate"`
	// The result code.
	ResultCode string `json:"ResultCode"`
	// A verbal explanation of the ResultCode.
	ResultMessage string `json:"ResultMessage"`
	// The type of the transaction.
	Type TransactionType `json:"Type"`
	// The initial transaction ID.
	InitialTransactionID string `json:"InitialTransactionId"`
	// The initial transaction type.
	InitialTransactionType TransactionType `json:"InitialTransactionType"`
	// Contains info about the reason for refund.
	RefundReason RefundReason `json:"RefundReason"`
}

// RefundReason describe the reason of the refund.
type RefundReason struct {
	// The type of reason for refusal
	RefusedReasonType RefundReasonType `json:"RefusedReasonType"`
	// The message accompanying a refusal
	RefusedReasonMessage string `json:"RefusedReasonMessage"`
}

// RefundReasonType is the reason type for the refund
type RefundReasonType string

const (
	// RefundReasonTypeInitializedByClient is for reason type INITIALIZED_BY_CLIENT.
	RefundReasonTypeInitializedByClient RefundReasonType = "INITIALIZED_BY_CLIENT"
	// RefundReasonTypeBankAccountIncorrect is for reason type BANKACCOUNT_INCORRECT.
	RefundReasonTypeBankAccountIncorrect RefundReasonType = "BANKACCOUNT_INCORRECT"
	// RefundReasonTypeOwnerDoNotMatchBankAccount is for reason type OWNER_DO_NOT_MATCH_BANKACCOUNT.
	RefundReasonTypeOwnerDoNotMatchBankAccount RefundReasonType = "OWNER_DO_NOT_MATCH_BANKACCOUNT"
	// RefundReasonTypeBankAccountHasBeenClosed is for reason type BANKACCOUNT_HAS_BEEN_CLOSED.
	RefundReasonTypeBankAccountHasBeenClosed RefundReasonType = "BANKACCOUNT_HAS_BEEN_CLOSED"
	// RefundReasonTypeWithdrawalImpossibleOnSavingsAccounts is for reason typWITHDRAWAL_IMPOSSIBLE_ON_SAVINGS_ACCOUNTS.e
	RefundReasonTypeWithdrawalImpossibleOnSavingsAccounts RefundReasonType = "WITHDRAWAL_IMPOSSIBLE_ON_SAVINGS_ACCOUNTS"
	// RefundReasonTypeOther is for reason type OTHER.
	RefundReasonTypeOther RefundReasonType = "OTHER"
)

// RefundCreate is for creating a refund.
// A PayIn Refund is a request to reimburse a user on their payment card.
// The money which has already been paid will automatically go back to the user’s bank account.
//
// If you're doing a partial Refund, note that you can only refund the same amount
// on the same transaction once per day (this is to prevent unintended duplicate refunds).
// After 24h you can do another refund of the same amount on the same transaction.
// If it is a different amount on the same transaction, there is not this limit.
// You can only refund a Payin in the 11months period following the original transaction.
//
// You can refund Card Direct Payins, Card Web Payins, Preauthorization Payins,
// Direct-Debit Web Payins and Direct-Debit Direct Payins.
// Note that for Direct-Debit Direct Payins, you are limited to 5 refunds for each Payin.
type RefundCreate struct {
	// A user's ID REQUIRED.
	AuthorID string `json:"AuthorId"`
	// Information about the funds that are being debited OPTIONAL.
	DebitedFunds Money `json:"DebitedFunds,omitempty"`
	// Information about the fees that were taken by the client OPTIONAL.
	// for this transaction (and were hence transferred to the Client's platform wallet)
	Fees Money `json:"Fees,omitempty"`
}
