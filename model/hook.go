package model

// Hook is used for a hook.
type Hook struct {
	// This is the URL where your receive notification for each EventType.
	URL string `json:"Url"`
	// Whether the hook is enabled or not.
	Status HookStatus `json:"Status"`
	// Whether the hook is valid or not.
	Validity HookValidity `json:"Validity"`
	// The event type.
	Event EventType `json:"Event"`
}

// HookStatus is the status of the hook.
type HookStatus string

const (
	// HookStatusDisable is for a status disabled.
	HookStatusDisable HookStatus = "DISABLED"
	// HookStatusEnabled is for a status enabled.
	HookStatusEnabled HookStatus = "ENABLED"
)

// HookValidity describe Whether the hook is enabled or not.
type HookValidity string

const (
	// HookValidityUnknown is for hook validity unknown.
	HookValidityUnknown HookValidity = "UNKNOWN"
	// HookValidityValid is for hook validity valid.
	HookValidityValid HookValidity = "VALID"
	// HookValidityInvalid is for hook validity invalid.
	HookValidityInvalid HookValidity = "INVALID"
)

// EventType is the type of event for the Hook.
type EventType string

const (
	// EventPayInNormalCreated is used for an event type PAYIN_NORMAL_CREATED.
	EventPayInNormalCreated EventType = "PAYIN_NORMAL_CREATED`"
	// EventPayInNormalSucceeded is used for an event type PAYIN_NORMAL_SUCCEEDED.
	EventPayInNormalSucceeded EventType = "PAYIN_NORMAL_SUCCEEDED`"
	// EventPayInNormalFailed is used for an event type PAYIN_NORMAL_FAILED.
	EventPayInNormalFailed EventType = "PAYIN_NORMAL_FAILED`"
	// EventPayOutNormalCreated is used for an event type PAYOUT_NORMAL_CREATED.
	EventPayOutNormalCreated EventType = "PAYOUT_NORMAL_CREATED`"
	// EventPayOutNormalSucceeded is used for an event type PAYOUT_NORMAL_SUCCEEDED.
	EventPayOutNormalSucceeded EventType = "PAYOUT_NORMAL_SUCCEEDED`"
	// EventPayOutNormalFailed is used for an event type PAYOUT_NORMAL_FAILED.
	EventPayOutNormalFailed EventType = "PAYOUT_NORMAL_FAILED`"
	// EventTransferNormalCreated is used for an event type TRANSFER_NORMAL_CREATED.
	EventTransferNormalCreated EventType = "TRANSFER_NORMAL_CREATED`"
	// EventTransferNormalSucceeded is used for an event type TRANSFER_NORMAL_SUCCEEDED.
	EventTransferNormalSucceeded EventType = "TRANSFER_NORMAL_SUCCEEDED`"
	// EventTransferNormalFailed is used for an event type TRANSFER_NORMAL_FAILED.
	EventTransferNormalFailed EventType = "TRANSFER_NORMAL_FAILED`"
	// EventPayInRefundCreated is used for an event type PAYIN_REFUND_CREATED.
	EventPayInRefundCreated EventType = "PAYIN_REFUND_CREATED`"
	// EventPayInRefundSucceeded is used for an event type PAYIN_REFUND_SUCCEEDED.
	EventPayInRefundSucceeded EventType = "PAYIN_REFUND_SUCCEEDED`"
	// EventPayInRefundFailed is used for an event type PAYIN_REFUND_FAILED.
	EventPayInRefundFailed EventType = "PAYIN_REFUND_FAILED`"
	// EventPayOutRefundCreated is used for an event type PAYOUT_REFUND_CREATED.
	EventPayOutRefundCreated EventType = "PAYOUT_REFUND_CREATED`"
	// EventPayOutRefundSucceeded is used for an event type PAYOUT_REFUND_SUCCEEDED.
	EventPayOutRefundSucceeded EventType = "PAYOUT_REFUND_SUCCEEDED`"
	// EventPayOutRefundFailed is used for an event type PAYOUT_REFUND_FAILED.
	EventPayOutRefundFailed EventType = "PAYOUT_REFUND_FAILED`"
	// EventTransferRefundCreated is used for an event type TRANSFER_REFUND_CREATED.
	EventTransferRefundCreated EventType = "TRANSFER_REFUND_CREATED`"
	// EventTransferRefundSucceeded is used for an event type TRANSFER_REFUND_SUCCEEDED.
	EventTransferRefundSucceeded EventType = "TRANSFER_REFUND_SUCCEEDED`"
	// EventTransferRefundFailed is used for an event type TRANSFER_REFUND_FAILED.
	EventTransferRefundFailed EventType = "TRANSFER_REFUND_FAILED`"
	// EventKYCCreated is used for an event type EventType.
	EventKYCCreated EventType = "KYC_CREATED`"
	// EventKYCValidationAsked is used for an event type KYC_VALIDATION_ASKED.
	EventKYCValidationAsked EventType = "KYC_VALIDATION_ASKED`"
	// EventKYCSucceeded is used for an event type.
	EventKYCSucceeded EventType = "KYC_SUCCEEDED` ="
	// EventKYCFailed is used for an event type EventType.
	EventKYCFailed EventType = "KYC_FAILED`"
	// EventPayInRepudiationCreated is used for an event type PAYIN_REPUDIATION_CREATED.
	EventPayInRepudiationCreated EventType = "PAYIN_REPUDIATION_CREATED`"
	// EventPayInRepudiationSucceeded is used for an event type PAYIN_REPUDIATION_SUCCEEDED.
	EventPayInRepudiationSucceeded EventType = "PAYIN_REPUDIATION_SUCCEEDED`"
	// EventPayInRepudiationFailed is used for an event type PAYIN_REPUDIATION_FAILED.
	EventPayInRepudiationFailed EventType = "PAYIN_REPUDIATION_FAILED`"
	// EventDisputeDocumentCreated is used for an event type DISPUTE_DOCUMENT_CREATED.
	EventDisputeDocumentCreated EventType = "DISPUTE_DOCUMENT_CREATED`"
	// EventDisputeDocumentValidationAsked is used for an event type DISPUTE_DOCUMENT_VALIDATION_ASKED.
	EventDisputeDocumentValidationAsked EventType = "DISPUTE_DOCUMENT_VALIDATION_ASKED`"
	// EventDisputeDocumentSucceeded is used for an event type DISPUTE_DOCUMENT_SUCCEEDED.
	EventDisputeDocumentSucceeded EventType = "DISPUTE_DOCUMENT_SUCCEEDED`"
	// EventDisputeDocumentFailed is used for an event type DISPUTE_DOCUMENT_FAILED.
	EventDisputeDocumentFailed EventType = "DISPUTE_DOCUMENT_FAILED`"
	// EventDisputeCreated is used for an event type.
	EventDisputeCreated EventType = "DISPUTE_CREATED` DISPUTE_CREATED"
	// EventDisputeSubmitted is used for an event type.
	EventDisputeSubmitted EventType = "DISPUTE_SUBMITTED` DISPUTE_SUBMITTED"
	// EventDisputeActionRequired is used for an event type DISPUTE_ACTION_REQUIRED.
	EventDisputeActionRequired EventType = "DISPUTE_ACTION_REQUIRED`"
	// EventDisputeFurtherActionRequired is used for an event type DISPUTE_FURTHER_ACTION_REQUIRED.
	EventDisputeFurtherActionRequired EventType = "DISPUTE_FURTHER_ACTION_REQUIRED`"
	// EventDisputeDisputeClosed is used for an event type DISPUTE_CLOSED.
	EventDisputeDisputeClosed EventType = "DISPUTE_CLOSED`"
	// EventDisputeDisputeSentToBank is used for an event type DISPUTE_SENT_TO_BANK.
	EventDisputeDisputeSentToBank EventType = "DISPUTE_SENT_TO_BANK`"
	// EventTransferSettlementCreated is used for an event type TRANSFER_SETTLEMENT_CREATED.
	EventTransferSettlementCreated EventType = "TRANSFER_SETTLEMENT_CREATED`"
	// EventTransferSettlementSucceeded is used for an event type TRANSFER_SETTLEMENT_SUCCEEDED.
	EventTransferSettlementSucceeded EventType = "TRANSFER_SETTLEMENT_SUCCEEDED`"
	// EventTransferSettlementFailed is used for an event type TRANSFER_SETTLEMENT_FAILED.
	EventTransferSettlementFailed EventType = "TRANSFER_SETTLEMENT_FAILED`"
	// EventMandateCreated is used for an event type.
	EventMandateCreated EventType = "MANDATE_CREATED` MANDATE_CREATED"
	// EventMandateFailed is used for an event type.
	EventMandateFailed EventType = "MANDATE_FAILED` ="
	// EventMandateActivated is used for an event type.
	EventMandateActivated EventType = "MANDATE_ACTIVATED` MANDATE_ACTIVATED"
	// EventMandateSubmitted is used for an event type.
	EventMandateSubmitted EventType = "MANDATE_SUBMITTED` MANDATE_SUBMITTED"
	// EventPreAuthorizationPaymentWaiting is used for an event type PREAUTHORIZATION_PAYMENT_WAITING.
	EventPreAuthorizationPaymentWaiting EventType = "PREAUTHORIZATION_PAYMENT_WAITING`"
	// EventPreAuthorizationPaymentExpired is used for an event type PREAUTHORIZATION_PAYMENT_EXPIRED.
	EventPreAuthorizationPaymentExpired EventType = "PREAUTHORIZATION_PAYMENT_EXPIRED`"
	// EventPreAuthorizationPaymentCanceled is used for an event type PREAUTHORIZATION_PAYMENT_CANCELED.
	EventPreAuthorizationPaymentCanceled EventType = "PREAUTHORIZATION_PAYMENT_CANCELED`"
	// EventPreAuthorizationPaymentValidated is used for an event type PREAUTHORIZATION_PAYMENT_VALIDATED.
	EventPreAuthorizationPaymentValidated EventType = "PREAUTHORIZATION_PAYMENT_VALIDATED`"
	// EventUBODeclarationCreated is used for an event type UBO_DECLARATION_CREATED.
	EventUBODeclarationCreated EventType = "UBO_DECLARATION_CREATED`"
	// EventUBODeclarationValidationAsked is used for an event type UBO_DECLARATION_VALIDATION_ASKED.
	EventUBODeclarationValidationAsked EventType = "UBO_DECLARATION_VALIDATION_ASKED`"
	// EventUBODeclarationRefused is used for an event type UBO_DECLARATION_REFUSED.
	EventUBODeclarationRefused EventType = "UBO_DECLARATION_REFUSED`"
	// EventUBODeclarationValidation is used for an event type UBO_DECLARATION_VALIDATED.
	EventUBODeclarationValidation EventType = "UBO_DECLARATION_VALIDATED`"
	// EventUBODeclarationIncomplete is used for an event type UBO_DECLARATION_INCOMPLETE.
	EventUBODeclarationIncomplete EventType = "UBO_DECLARATION_INCOMPLETE"
)

// HookCreate is the payload used for creating a new hook.
type HookCreate struct {
	// The event type REQUIRED.
	EventType EventType `json:"EventType"`
	// This is the URL where your receive notification for each EventType REQUIRED.
	URL string `json:"Url"`
}
