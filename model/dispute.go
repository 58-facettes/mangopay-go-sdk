package model

type Dispute struct {
	// The initial transaction ID
	InitialTransactionId string
	// The initial transaction type
	InitialTransactionType TransactionType
	// The result code
	ResultCode string
	// A verbal explanation of the ResultCode
	ResultMessage string
	// Info about the reason for the dispute
	DisputeReason DisputeReason
	// The status of the dispute
	Status DisputeStatus
	// Used to communicate information about the dispute status to you
	StatusMessage string
	// The amount of funds that were disputed
	DisputedFunds Money
	// The amount you wish to contest
	ContestedFunds Money
	// The type of dispute
	DisputeType DisputeType
	// The deadline by which you must contest the dispute (if you wish to contest it)
	ContestDeadlineDate int64
	// The ID of the associated repudiation transaction
	RepudiationId string
}

type DisputeReason struct {
	// The type of reason for the dispute
	DisputeReasonType DisputeReasonType
	// More information about the reason for the dispute
	DisputeReasonMessage string
}

// DisputeReasonType is the reason type for the dispute.
type DisputeReasonType string

const (
	// DisputeReasonDuplicated stands for a DUPLICATE reason.
	DisputeReasonDuplicated DisputeReasonType = "DUPLICATE"
	// DisputeReasonFraud stands for a FRAUD reason.
	DisputeReasonFraud DisputeReasonType = "FRAUD"
	// DisputeReasonProductUnacceptable stands for a PRODUCT_UNACCEPTABLE reason.
	DisputeReasonProductUnacceptable DisputeReasonType = "PRODUCT_UNACCEPTABLE"
	// DisputeReasonUnknown stands for a UNKNOWN reason.
	DisputeReasonUnknown DisputeReasonType = "UNKNOWN"
	// DisputeReasonOther stands for a OTHER reason.
	DisputeReasonOther DisputeReasonType = "OTHER"
	// DisputeReasonRefundConverstionRate stands for a REFUND_CONVERSION_RATE reason.
	DisputeReasonRefundConverstionRate DisputeReasonType = "REFUND_CONVERSION_RATE"
	// DisputeReasonLateFailureInsufficientFunds stands for a LATE_FAILURE_INSUFFICIENT_FUNDS reason.
	DisputeReasonLateFailureInsufficientFunds DisputeReasonType = "LATE_FAILURE_INSUFFICIENT_FUNDS"
	// DisputeReasonLateFailureContactUser stands for a LATE_FAILURE_CONTACT_USER reason.
	DisputeReasonLateFailureContactUser DisputeReasonType = "LATE_FAILURE_CONTACT_USER"
	// DisputeReasonLateFailureBankAccountClosed stands for a LATE_FAILURE_BANKACCOUNT_CLOSED reason.
	DisputeReasonLateFailureBankAccountClosed DisputeReasonType = "LATE_FAILURE_BANKACCOUNT_CLOSED"
	// DisputeReasonLateFailureBancAccountIncompatible stands for a LATE_FAILURE_BANKACCOUNT_INCOMPATIBLE reason.
	DisputeReasonLateFailureBancAccountIncompatible DisputeReasonType = "LATE_FAILURE_BANKACCOUNT_INCOMPATIBLE"
	// DisputeReasonLateFailureBankAccountIncorrect stands for a LATE_FAILURE_BANKACCOUNT_INCORRECT reason.
	DisputeReasonLateFailureBankAccountIncorrect DisputeReasonType = "LATE_FAILURE_BANKACCOUNT_INCORRECT"
	// DisputeReasonAutorizationDisputed stands for a AUTHORISATION_DISPUTED reason.
	// ?? AUTHORISATION or AUTHORIZATION??
	DisputeReasonAutorizationDisputed DisputeReasonType = "AUTHORISATION_DISPUTED"
	// DisputeReasonTransactionNotReconized stands for a TRANSACTION_NOT_RECOGNIZED reason.
	DisputeReasonTransactionNotReconized DisputeReasonType = "TRANSACTION_NOT_RECOGNIZED"
	// DisputeReasonProductNotProvided stands for a PRODUCT_NOT_PROVIDED reason.
	DisputeReasonProductNotProvided DisputeReasonType = "PRODUCT_NOT_PROVIDED"
	// DisputeReasonCanceledReccuringTransaction stands for a CANCELED_REOCCURING_TRANSACTION reason.
	// ?? CANCELED_REOCCURING_TRANSACTION or CANCELED_RECCURING_TRANSACTION ??
	DisputeReasonCanceledReccuringTransaction DisputeReasonType = "CANCELED_REOCCURING_TRANSACTION"
	// DisputeReasonRefundNotProcessed stands for a REFUND_NOT_PROCESSED reason.
	DisputeReasonRefundNotProcessed DisputeReasonType = "REFUND_NOT_PROCESSED"
)

// DisputeType is the type of Dispute.
type DisputeType string

const (
	// DisputeContestable stands for CONTESTABLE.
	DisputeContestable DisputeType = "CONTESTABLE"
	// DisputeNotContestable stands for NOT_CONTESTABLE.
	DisputeNotContestable DisputeType = "NOT_CONTESTABLE"
	// DisputeRetrieval stands for RETRIEVAL.
	DisputeRetrieval DisputeType = "RETRIEVAL"
)

// DisputeStatus is the status of the Dispute.
type DisputeStatus string

const (
	// DisputeStatusCreated stands for the status CREATED.
	DisputeStatusCreated DisputeStatus = "CREATED"
	// DisputeStatusPendingClientAction stands for the status PENDING_CLIENT_ACTION.
	DisputeStatusPendingClientAction DisputeStatus = "PENDING_CLIENT_ACTION"
	// DisputeStatusSubmitted stands for the status SUBMITTED.
	DisputeStatusSubmitted DisputeStatus = "SUBMITTED"
	// DisputeStatusPendingBankAction stands for the status PENDING_BANK_ACTION.
	DisputeStatusPendingBankAction DisputeStatus = "PENDING_BANK_ACTION"
	// DisputeStatusReopenedPendingClientAction stands for the status REOPENED_PENDING_CLIENT_ACTION.
	DisputeStatusReopenedPendingClientAction DisputeStatus = "REOPENED_PENDING_CLIENT_ACTION"
	// DisputeStatusClosed stands for the status CLOSED.
	DisputeStatusClosed DisputeStatus = "CLOSED"
)
