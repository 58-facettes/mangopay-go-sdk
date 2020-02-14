package model

type Hook struct {
	// This is the URL where your receive notification for each EventType
	URL string `json:"Url"`
	// Whether the hook is enabled or not
	Status HookStatus `json:"Status"`
	// Whether the hook is valid or not
	Validity HookValidity `json:"Validity"`
	// The event type
	Event EventType `json:"Event"`
}

// The status of the hook
type HookStatus string

const (
	HookStatusDisable HookStatus = "DISABLED"
	HookStatusEnabled HookStatus = "ENABLED"
)

// Whether the hook is enabled or not
type HookValidity string

const (
	HookValidityUnknown HookValidity = "UNKNOWN"
	HookValidityValid   HookValidity = "VALID"
	HookValidityInvalid HookValidity = "INVALID"
)

// EventType is the type of event for the Hook.
type EventType string

const (
	EventPayInNormalCreated               EventType = "PAYIN_NORMAL_CREATED`"
	EventPayInNormalSucceeded             EventType = "PAYIN_NORMAL_SUCCEEDED`"
	EventPayInNormalFailed                EventType = "PAYIN_NORMAL_FAILED`"
	EventPayOutNormalCreated              EventType = "PAYOUT_NORMAL_CREATED`"
	EventPayOutNormalSucceeded            EventType = "PAYOUT_NORMAL_SUCCEEDED`"
	EventPayOutNormalFailed               EventType = "PAYOUT_NORMAL_FAILED`"
	EventTransferNormalCreated            EventType = "TRANSFER_NORMAL_CREATED`"
	EventTransferNormalSucceeded          EventType = "TRANSFER_NORMAL_SUCCEEDED`"
	EventTransferNormalFailed             EventType = "TRANSFER_NORMAL_FAILED`"
	EventPayInRefundCreated               EventType = "PAYIN_REFUND_CREATED`"
	EventPayInRefundSucceeded             EventType = "PAYIN_REFUND_SUCCEEDED`"
	EventPayInRefundFailed                EventType = "PAYIN_REFUND_FAILED`"
	EventPayOutRefundCreated              EventType = "PAYOUT_REFUND_CREATED`"
	EventPayOutRefundSucceeded            EventType = "PAYOUT_REFUND_SUCCEEDED`"
	EventPayOutRefundFailed               EventType = "PAYOUT_REFUND_FAILED`"
	EventTransferRefundCreated            EventType = "TRANSFER_REFUND_CREATED`"
	EventTransferRefundSucceeded          EventType = "TRANSFER_REFUND_SUCCEEDED`"
	EventTransferRefundFailed             EventType = "TRANSFER_REFUND_FAILED`"
	EventKYCCreated                       EventType = "KYC_CREATED`"
	EventKYCValidationAsked               EventType = "KYC_VALIDATION_ASKED`"
	EventKYCSucceeded                     EventType = "KYC_SUCCEEDED`"
	EventKYCFailed                        EventType = "KYC_FAILED`"
	EventPayInRepudiationCreated          EventType = "PAYIN_REPUDIATION_CREATED`"
	EventPayInRepudiationSucceeded        EventType = "PAYIN_REPUDIATION_SUCCEEDED`"
	EventPayInRepudiationFailed           EventType = "PAYIN_REPUDIATION_FAILED`"
	EventDisputeDocumentCreated           EventType = "DISPUTE_DOCUMENT_CREATED`"
	EventDisputeDocumentValidationAsked   EventType = "DISPUTE_DOCUMENT_VALIDATION_ASKED`"
	EventDisputeDocumentSucceeded         EventType = "DISPUTE_DOCUMENT_SUCCEEDED`"
	EventDisputeDocumentFailed            EventType = "DISPUTE_DOCUMENT_FAILED`"
	EventDisputeCreated                   EventType = "DISPUTE_CREATED`"
	EventDisputeSubmitted                 EventType = "DISPUTE_SUBMITTED`"
	EventDisputeActionRequired            EventType = "DISPUTE_ACTION_REQUIRED`"
	EventDisputeFurtherActionRequired     EventType = "DISPUTE_FURTHER_ACTION_REQUIRED`"
	EventDisputeDisputeClosed             EventType = "DISPUTE_CLOSED`"
	EventDisputeDisputeSentToBank         EventType = "DISPUTE_SENT_TO_BANK`"
	EventTransferSettlementCreated        EventType = "TRANSFER_SETTLEMENT_CREATED`"
	EventTransferSettlementSucceeded      EventType = "TRANSFER_SETTLEMENT_SUCCEEDED`"
	EventTransferSettlementFailed         EventType = "TRANSFER_SETTLEMENT_FAILED`"
	EventMandateCreated                   EventType = "MANDATE_CREATED`"
	EventMandateFailed                    EventType = "MANDATE_FAILED`"
	EventMandateActivated                 EventType = "MANDATE_ACTIVATED`"
	EventMandateSubmitted                 EventType = "MANDATE_SUBMITTED`"
	EventPreAuthorizationPaymentWaiting   EventType = "PREAUTHORIZATION_PAYMENT_WAITING`"
	EventPreAuthorizationPaymentExpired   EventType = "PREAUTHORIZATION_PAYMENT_EXPIRED`"
	EventPreAuthorizationPaymentCanceled  EventType = "PREAUTHORIZATION_PAYMENT_CANCELED`"
	EventPreAuthorizationPaymentValidated EventType = "PREAUTHORIZATION_PAYMENT_VALIDATED`"
	EventUBODeclarationCreated            EventType = "UBO_DECLARATION_CREATED`"
	EventUBODeclarationValidationAsked    EventType = "UBO_DECLARATION_VALIDATION_ASKED`"
	EventUBODeclarationRefused            EventType = "UBO_DECLARATION_REFUSED`"
	EventUBODeclarationValidation         EventType = "UBO_DECLARATION_VALIDATED`"
	EventUBODeclarationIncomplete         EventType = "UBO_DECLARATION_INCOMPLETE"
)

type HookCreate struct {
	// The event type REQUIRED.
	EventType EventType `json:"EventType"`
	// This is the URL where your receive notification for each EventType REQUIRED.
	URL string `json:"Url"`
}
