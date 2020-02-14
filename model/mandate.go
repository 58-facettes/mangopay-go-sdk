package model

// Mandate is a Direct Debit Mandate is an instruction between a user and a bank account
// which allows you to process payments directly from his bank to a wallet for a dedicated user.
type Mandate struct {
	// An ID of a Bank Account
	BankAccountID string `json:"BankAccountId"`
	// The object owner's UserId
	UserID string `json:"UserId"`
	// The URL to redirect to after payment (whether successful or not)
	ReturnURL string `json:"ReturnURL"`
	// The URL to redirect to user to for them to proceed with the payment
	RedirectURL string `json:"RedirectURL"`
	// The URL to download the mandate
	DocumentURL string `json:"DocumentURL"`
	// The language to use for the mandate confirmation page - needs to be the ISO code of the language
	Culture MandateCultureCode `json:"Culture"`
	// The type of mandate, but will only be completed once the mandate has been submitted
	Scheme MandateScheme `json:"Scheme"`
	// The status of the mandate:
	//
	// "CREATED" - the mandate has been created
	// "SUBMITTED" - the mandate has been submitted to the banks and you can now do payments with this mandate
	// "ACTIVE" - the mandate is active and has been accepted by the banks and/or successfully used in a payment
	// "FAILED" - the mandate has failed for a variety of reasons and is no longer available for payments
	Status MandateStatus `json:"Status"`
	// The result code
	ResultCode string `json:"ResultCode"`
	// A verbal explanation of the ResultCode
	ResultMessage string `json:"ResultMessage"`
	// The execution type for creating the mandate
	ExecutionType MandateExecutionType `json:"ExecutionType"`
	// The type of Mandate
	MandateType MandateType `json:"MandateType"`
	// The banking reference for this mandate
	BankReference string `json:"BankReference"`
}

// The language to use for the mandate confirmation webpage
type MandateCultureCode string

const (
	MandateCultureCodeEN MandateCultureCode = "EN"
	MandateCultureCodeFR MandateCultureCode = "FR"
	MandateCultureCodeNL MandateCultureCode = "NL"
	MandateCultureCodeDE MandateCultureCode = "DE"
	MandateCultureCodeES MandateCultureCode = "ES"
	MandateCultureCodeIT MandateCultureCode = "IT"
	MandateCultureCodePL MandateCultureCode = "PL"
)

// The mandate scheme
type MandateScheme string

const (
	MandateSchemeSEPA MandateScheme = "SEPA"
	MandateSchemeBACS MandateScheme = "BACS"
)

// The status of a mandate
type MandateStatus string

const (
	MandateStatusCreated   MandateStatus = "CREATED"
	MandateStatusSubmitted MandateStatus = "SUBMITTED"
	MandateStatusActive    MandateStatus = "ACTIVE"
	MandateStatusFailed    MandateStatus = "FAILED"
)

// The execution type of the mandate
type MandateExecutionType string

const (
	MandateExecutionWEB MandateExecutionType = "WEB"
)

// The type of mandate
type MandateType string

const (
	MandateTypeDirectDebit MandateType = "DIRECT_DEBIT"
)

type MandateCreate struct {
	// An ID of a Bank Account REQUIRED.
	BankAccountID string `json:"BankAccountId"`
	// The language to use for the mandate confirmation page - needs to be the ISO code of the language REQUIRED.
	Culture MandateCultureCode `json:"Culture"`
	// The URL to redirect to after payment (whether successful or not) REQUIRED.
	ReturnURL string `json:"ReturnURL"`
}
