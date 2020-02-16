package model

// Mandate is a Direct Debit Mandate is an instruction between a user and a bank account
// which allows you to process payments directly from his bank to a wallet for a dedicated user.
type Mandate struct {
	// An ID of a Bank Account.
	BankAccountID string `json:"BankAccountId"`
	// The object owner's UserId.
	UserID string `json:"UserId"`
	// The URL to redirect to after payment (whether successful or not).
	ReturnURL string `json:"ReturnURL"`
	// The URL to redirect to user to for them to proceed with the payment.
	RedirectURL string `json:"RedirectURL"`
	// The URL to download the mandate.
	DocumentURL string `json:"DocumentURL"`
	// The language to use for the mandate confirmation page - needs to be the ISO code of the language.
	Culture MandateCultureCode `json:"Culture"`
	// The type of mandate, but will only be completed once the mandate has been submitted.
	Scheme MandateScheme `json:"Scheme"`
	// The status of the mandate:
	//
	// "CREATED" - the mandate has been created
	// "SUBMITTED" - the mandate has been submitted to the banks and you can now do payments with this mandate.
	// "ACTIVE" - the mandate is active and has been accepted by the banks and/or successfully used in a payment.
	// "FAILED" - the mandate has failed for a variety of reasons and is no longer available for payments.
	Status MandateStatus `json:"Status"`
	// The result code
	ResultCode string `json:"ResultCode"`
	// A verbal explanation of the ResultCode.
	ResultMessage string `json:"ResultMessage"`
	// The execution type for creating the mandate.
	ExecutionType MandateExecutionType `json:"ExecutionType"`
	// The type of Mandate.
	MandateType MandateType `json:"MandateType"`
	// The banking reference for this mandate.
	BankReference string `json:"BankReference"`
}

// MandateCultureCode is the language to use for the mandate confirmation webpage.
type MandateCultureCode string

const (
	// MandateCultureCodeEN if for an iso2 country name EN.
	MandateCultureCodeEN MandateCultureCode = "EN"
	// MandateCultureCodeFR if for an iso2 country name FR.
	MandateCultureCodeFR MandateCultureCode = "FR"
	// MandateCultureCodeNL if for an iso2 country name NL.
	MandateCultureCodeNL MandateCultureCode = "NL"
	// MandateCultureCodeDE if for an iso2 country name DE.
	MandateCultureCodeDE MandateCultureCode = "DE"
	// MandateCultureCodeES if for an iso2 country name ES.
	MandateCultureCodeES MandateCultureCode = "ES"
	// MandateCultureCodeIT if for an iso2 country name IT.
	MandateCultureCodeIT MandateCultureCode = "IT"
	// MandateCultureCodePL if for an iso2 country name PL.
	MandateCultureCodePL MandateCultureCode = "PL"
)

// MandateScheme is the mandate scheme use.
type MandateScheme string

const (
	// MandateSchemeSEPA is for a kind fo mandate SEPA.
	MandateSchemeSEPA MandateScheme = "SEPA"
	// MandateSchemeBACS is for a kind fo mandate BACS.
	MandateSchemeBACS MandateScheme = "BACS"
)

// MandateStatus is the status of a mandate.
type MandateStatus string

const (
	// MandateStatusCreated is used for the status created.
	MandateStatusCreated MandateStatus = "CREATED"
	// MandateStatusSubmitted is used for the status submitted.
	MandateStatusSubmitted MandateStatus = "SUBMITTED"
	// MandateStatusActive is used for the status active.
	MandateStatusActive MandateStatus = "ACTIVE"
	// MandateStatusFailed is used for the status failed.
	MandateStatusFailed MandateStatus = "FAILED"
)

// MandateExecutionType is the execution type of the mandate.
type MandateExecutionType string

const (
	// MandateExecutionWEB is the execution type web.
	MandateExecutionWEB MandateExecutionType = "WEB"
)

// MandateType is the type of mandate.
type MandateType string

const (
	// MandateTypeDirectDebit is for direct debit.
	MandateTypeDirectDebit MandateType = "DIRECT_DEBIT"
)

// MandateCreate is the payload used for creating a new Mandate.
type MandateCreate struct {
	// An ID of a Bank Account REQUIRED.
	BankAccountID string `json:"BankAccountId"`
	// The language to use for the mandate confirmation page - needs to be the ISO code of the language REQUIRED.
	Culture MandateCultureCode `json:"Culture"`
	// The URL to redirect to after payment (whether successful or not) REQUIRED.
	ReturnURL string `json:"ReturnURL"`
}
