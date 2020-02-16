package model

import (
	"github.com/58-facettes/mangopay-go-sdk/model/country"
)

// Payin is used for the payin instrcution.
type Payin struct {
	// The type of payin.
	Payment Payment `json:"PaymentType"`
	// The type of execution for the payin.
	Execution Execution `json:"ExecutionType"`
}

// Payment is the payment type of the payin.
type Payment string

const (
	// PaymentCard is for a payment with card.
	PaymentCard Payment = "CARD"
	// PaymentDirectDebit is for a payment with direct debit.
	PaymentDirectDebit Payment = "DIRECT_DEBIT"
	// PaymentPreauthorized is for a payment with preauthoorized.
	PaymentPreauthorized Payment = "PREAUTHORIZED"
	// PaymentBankWire is for a payment with bank wire.
	PaymentBankWire Payment = "BANK_WIRE"
)

// Execution is the execution type of the payin.
type Execution string

const (
	// ExecutionWeb is for an execution web.
	ExecutionWeb Execution = "WEB"
	// ExecutionDirect is for an execution direct.
	ExecutionDirect Execution = "DIRECT"
	// ExecutionExternalInstruction is for an execution external instruction.
	ExecutionExternalInstruction Execution = "EXTERNAL_INSTRUCTION"
)

type CardWeb struct {
	// The URL to redirect to after payment (whether successful or not).
	ReturnURL string
	// The type of card . The card type is optional, but the default parameter is "CB_VISA_MASTERCARD".
	CardType CardType
	// The SecureMode corresponds to '3D secure' for CB Visa and MasterCard.
	// This field lets you activate it manually.
	// The field lets you activate it automatically
	// with "DEFAULT" (Secured Mode will be activated from €50 or when MANGOPAY detects there is a higher risk ),
	// "FORCE" (if you wish to specifically force the secured mode).
	SecureMode SecureMode
	// The language to use for the payment page - needs to be the ISO code of the language
	Culture CultureCode
	// The URL to use for the payment page template
	TemplateURL string
	// A custom description to appear on the user's bank statement.
	// It can be up to 10 characters long, and can only include alphanumeric characters or spaces.
	// See here for important info. Note that each bank handles this information differently, some show less or no information.
	StatementDescriptor string
	// The URL to redirect to user to for them to proceed with the payment.
	RedirectURL string
}

// SecureMode is the type of mode.
type SecureMode string

const (
	// SecureModeDefault is the mode of payin by DEFAULT.
	SecureModeDefault SecureMode = "DEFAULT"
	// SecureModeForce is the mode of payin by FORCE.
	SecureModeForce SecureMode = "FORCE"
)

// CultureCode is the language to use for the payment webpage.
type CultureCode string

const (
	// CultureCodeDE is the culture code for the country iso2 DE.
	CultureCodeDE CultureCode = "DE"
	// CultureCodeEN is the culture code for the country iso2 EN.
	CultureCodeEN CultureCode = "EN"
	// CultureCodeDA is the culture code for the country iso2 DA.
	CultureCodeDA CultureCode = "DA"
	// CultureCodeES is the culture code for the country iso2 ES.
	CultureCodeES CultureCode = "ES"
	// CultureCodeET is the culture code for the country iso2 ET.
	CultureCodeET CultureCode = "ET"
	// CultureCodeFI is the culture code for the country iso2 FI.
	CultureCodeFI CultureCode = "FI"
	// CultureCodeFR is the culture code for the country iso2 FR.
	CultureCodeFR CultureCode = "FR"
	// CultureCodeEL is the culture code for the country iso2 EL.
	CultureCodeEL CultureCode = "EL"
	// CultureCodeHU is the culture code for the country iso2 HU.
	CultureCodeHU CultureCode = "HU"
	// CultureCodeIT is the culture code for the country iso2 IT.
	CultureCodeIT CultureCode = "IT"
	// CultureCodeNL is the culture code for the country iso2 NL.
	CultureCodeNL CultureCode = "NL"
	// CultureCodeNO is the culture code for the country iso2 NO.
	CultureCodeNO CultureCode = "NO"
	// CultureCodePL is the culture code for the country iso2 PL.
	CultureCodePL CultureCode = "PL"
	// CultureCodePT is the culture code for the country iso2 PT.
	CultureCodePT CultureCode = "PT"
	// CultureCodeSK is the culture code for the country iso2 SK.
	CultureCodeSK CultureCode = "SK"
	// CultureCodeSV is the culture code for the country iso2 SV.
	CultureCodeSV CultureCode = "SV"
	// CultureCodeCS is the culture code for the country iso2 CS.
	CultureCodeCS CultureCode = "CS"
)

type CardWebCreate struct {
	// Custom data that you can add to this item OPTIONAL.
	Tag string `json:"Tag,omitempty"`
	// A user's ID REQUIRED.
	AuthorID string `json:"AuthorId"`
	// The user ID who is credited (defaults to the owner of the wallet) OPTIONAL.
	CreditedUserID string `json:"CreditedUserId,omitempty"`
	// Information about the funds that are being debited REQUIRED.
	DebitedFunds Money `json:"DebitedFunds"`
	// Information about the fees that were taken by the client REQUIRED.
	// for this transaction (and were hence transferred to the Client's platform wallet).
	Fees Money `json:"Fees"`
	// The URL to redirect to after payment (whether successful or not) REQUIRED.
	ReturnURL string `json:"ReturnURL"`
	// The type of card . The card type is optional, but the default parameter is "CB_VISA_MASTERCARD" REQUIRED.
	CardType CardType `json:"CardType"`
	// The ID of the wallet where money will be credited REQUIRED.
	CreditedWalletID string `json:"CreditedWalletId"`
	// The SecureMode corresponds to '3D secure' for CB Visa and MasterCard OPTIONAL.
	// This field lets you activate it manually. The field lets you activate it automatically
	// with "DEFAULT" (Secured Mode will be activated from €50 or when MANGOPAY detects there is a higher risk ),
	// "FORCE" (if you wish to specifically force the secured mode).
	SecureMode SecureMode `json:"SecureMode,omitempty"`
	// The language to use for the payment page - needs to be the ISO code of the language REQUIRED.
	Culture CultureCode `json:"Culture"`
	// A URL to an SSL page to allow you to customise the payment page OPTIONAL.
	// Must be in the format: array("PAYLINE"=>"https://...") and meet all the specifications listed here.
	// Note that only a template for Payline is currently available.
	TemplateURLOptions TemplateURLOptions `json:"TemplateURLOptions,omitempty"`
	// A custom description to appear on the user's bank statement OPTIONAL.
	// It can be up to 10 characters long, and can only include alphanumeric characters or spaces.
	// See here for important info. Note that each bank handles this information differently, some show less or no information.
	StatementDescriptor string `json:"StatementDescriptor,omitempty"`
}

type TemplateURLOptions struct {
	// The corresponding template URL for this PSP REQUIRED.
	Payline string `json:"Payline"`
}

type CardDirect struct {
	// This is the URL where users are automatically redirected after 3D secure validation (if activated).
	SecureModeReturnURL string `json:"SecureModeReturnURL"`
	// The ID of a card
	CardID string `json:"CardId"`
	// The SecureMode corresponds to '3D secure' for CB Visa and MasterCard.
	// This field lets you activate it manually.
	// The field lets you activate it automatically with "DEFAULT"
	// (Secured Mode will be activated from €50 or when MANGOPAY detects there is a higher risk ),
	// "FORCE" (if you wish to specifically force the secured mode).
	SecureMode SecureMode `json:"SecureMode"`
	// A custom description to appear on the user's bank statement.
	// It can be up to 10 characters long, and can only include alphanumeric characters or spaces.
	// See here for important info. Note that each bank handles this information differently, some show less or no information.
	StatementDescriptor string `json:"StatementDescriptor"`
	// Contains every useful informations related to the user billing.
	Billing Billing `json:"Billing"`
	// Contains useful informations related to security and fraud.
	SecurityInfo SecurityInfo `json:"SecurityInfo"`
	// The language to use for the payment page - needs to be the ISO code of the language.
	Culture CultureCode `json:"Culture"`
	// The value is 'true' if the SecureMode was used.
	SecureModeNeeded bool `json:"SecureModeNeeded"`
	// This is the URL where to redirect users to proceed to 3D secure validation.
	SecureModeRedirectURL string `json:"SecureModeRedirectUrl"`
}

type CardDirectCreate struct {
	// A user's ID REQUIRED.
	AuthorID string `json:"AuthorId"`
	// The user ID who is credited (defaults to the owner of the wallet) OPTIONAL.
	CreditedUserID string `json:"CreditedUserId,omitempty"`
	// The ID of the wallet where money will be credited REQUIRED.
	CreditedWalletID string `json:"CreditedWalletId"`
	// Information about the funds that are being debited REQUIRED
	DebitedFunds Money `json:"DebitedFunds"`
	// Information about the fees that were taken by the client REQUIRED.
	// for this transaction (and were hence transferred to the Client's platform wallet).
	Fees Money `json:"Fees"`
	// This is the URL where users are automatically redirected after 3D secure validation (if activated) REQUIRED.
	SecureModeReturnURL string `json:"SecureModeReturnURL"`
	// The ID of a card REQUIRED.
	CardID string `json:"CardId"`
	// The SecureMode corresponds to '3D secure' for CB Visa and MasterCard OPTIONAL.
	// This field lets you activate it manually.
	// The field lets you activate it automatically with "DEFAULT"
	// (Secured Mode will be activated from €50 or when MANGOPAY detects there is a higher risk ),
	// "FORCE" (if you wish to specifically force the secured mode).
	SecureMode SecureMode `json:"SecureMode,omitempty"`
	// A custom description to appear on the user's bank statement OPTIONAL.
	// It can be up to 10 characters long, and can only include alphanumeric characters or spaces.
	// See here for important info. Note that each bank handles this information differently, some show less or no information.
	StatementDescriptor string `json:"StatementDescriptor,omitempty"`
	// Contains every useful informations related to the user billing OPTIONAL.
	Billing Billing `json:"Billing,omitempty"`
	// The language to use for the payment page OPTIONAL.
	// - needs to be the ISO code of the language.
	Culture CultureCode `json:"Culture,omitempty"`
	// Custom data that you can add to this item OPTIONAL.
	Tag string `json:"Tag,omitempty"`
}

type Billing struct {
	//The address.
	Address Address `json:"Address"`
}

type SecurityInfo struct {
	// Result of Address Verification System check (only available for UK, US and Australia).
	AVSResult AVSResult `json:"AVSResult"`
}

// Result of Address Verification System check.
type AVSResult string

const (
	AVSResultNoCheck             AVSResult = "NO_CHECK"
	AVSResultNoMatch             AVSResult = "NO_MATCH"
	AVSResultAddressMatchOnly    AVSResult = "ADDRESS_MATCH_ONLY"
	AVSResultPostalCodeMatchOnly AVSResult = "POSTAL_CODE_MATCH_ONLY"
	AVSResultFullMatch           AVSResult = "FULL_MATCH"
)

type CardPreAuthorized struct {
	// The ID of the Preauthorization object
	PreauthorizationID string `json:"PreauthorizationId"`
}

type CardPreAuthorizedCreate struct {
	// A user's ID OPTIONAL.
	AuthorID string `json:"AuthorId,omitempty"`
	// The user ID who is credited (defaults to the owner of the wallet) OPTIONAL.
	CreditedUserID string `json:"CreditedUserId,omitempty"`
	// The ID of the wallet where money will be credited REQUIRED.
	CreditedWalletID string `json:"CreditedWalletId"`
	// Information about the funds that are being debited REQUIRED.
	DebitedFunds Money `json:"DebitedFunds"`
	// Information about the fees that were taken by the client for this transaction
	// (and were hence transferred to the Client's platform wallet) REQUIRED.
	Fees Money `json:"Fees"`
	// The ID of the Preauthorization object REQUIRED.
	PreauthorizationID string `json:"PreauthorizationId"`
}

type BankwireDirect struct {
	ID string `json:"Id"`
	// Custom data that you can add to this item OPTIONAL.
	Tag          string `json:"Tag"`
	CreationDate int64  `json:"CreationDate"`
	// The type of payin
	Payment Payment `json:"Payment"`
	// The type of execution for the payin
	Execution Execution `json:"Execution"`
	// The declared debited funds
	DeclaredDebitedFunds Money `json:"DeclaredDebitedFunds"`
	// The declared fees
	DeclaredFees Money `json:"DeclaredFees"`
	// Wire reference
	WireReference string `json:"WireReference"`
	// Bank account details
	BankAccount BankAccountV2 `json:"BankAccount"`
}

// BankAccountV2 describes a bank account access.
//?? can we name it BankAccount ??
type BankAccountV2 struct {
	// The BIC of the bank account
	BIC string `json:"BIC"`
	// The IBAN of the bank account
	IBAN string `json:"IBAN"`
	// The name of the owner of the bank account
	OwnerName string `json:"OwnerName"`
	// The address of the owner of the bank account
	OwnerAddress string `json:"OwnerAddress"`
	// The type of bank account
	Type BankAccountType `json:"Type"`
}

// The type of BankAccount
type BankAccountType string

const (
	BankAccountTypeIBAN  BankAccountType = "IBAN"
	BankAccountTypeGB    BankAccountType = "GB"
	BankAccountTypeUS    BankAccountType = "US"
	BankAccountTypeCA    BankAccountType = "CA"
	BankAccountTypeOther BankAccountType = "OTHER"
)

type BankwireDirectCreate struct {
	// A user's ID REQUIRED.
	AuthorID string `json:"AuthorId"`
	// The user ID who is credited (defaults to the owner of the wallet) OPTIONAL.
	CreditedUserID string `json:"CreditedUserId,omitempty"`
	// The ID of the wallet where money will be credited REQUIRED.
	CreditedWalletID string `json:"CreditedWalletId"`
	// The declared debited funds REQUIRED.
	DeclaredDebitedFunds Money `json:"DeclaredDebitedFunds"`
	// The declared fees REQUIRED.
	DeclaredFees Money `json:"DeclaredFees"`
}

// A bankwire by "external instruction" is a slightly particular payin whereby (BETA*).
// it is created when we receive funds for a banking alias
// 	 - this means there is no way of creating a payin of this kind via the API
// 	 (in sandbox, please contact our support team for them to create some examples for you).
type BankwireExternalInstruction struct {
	// The ID of a banking alias
	BankingAliasID string `json:"BankingAliasId"`
	// Wire reference
	WireReference string `json:"WireReference"`
	// Information about the account that was debited
	BankAccountDebited BankAccountDebited `json:"DebitedBankAccount"`
}

type BankAccountDebited struct {
	// The name of the owner of the bank account
	OwnerName string `json:"OwnerName"`
	// The type of bank account
	Type string `json:"Type"`
	// The IBAN of the bank account
	IBAN string `json:"IBAN"`
	// The BIC of the bank account
	BIC string `json:"BIC"`
}

// The Direct-Debit Web PayIn
type DirectDebitWeb struct {
	// The type of payin
	PaymentType Payment `json:"PaymentType"` // ?? the fiedl shoud be PayInPayment instead of PaymentType ??
	// The type of execution for the payin
	ExecutionType Execution `json:"ExecutionType"`
	// The URL to redirect to after payment (whether successful or not)
	ReturnURL string `json:"ReturnURL"`
	// The type of web direct debit
	DirectDebit DirectDebit `json:"DirectDebit"`
	// The SecureMode corresponds to '3D secure' for CB Visa and MasterCard.
	// This field lets you activate it manually. The field lets you activate it automatically
	// with "DEFAULT" (Secured Mode will be activated from €50 or when MANGOPAY detects there is a higher risk ),
	// "FORCE" (if you wish to specifically force the secured mode).
	SecureMode SecureMode `json:"SecureMode"`
	// The language to use for the payment page - needs to be the ISO code of the language
	Culture CultureCode `json:"Culture"`
	// The URL to use for the payment page template
	TemplateURL string `json:"TemplateURL"`
	// The URL to redirect to user to for them to proceed with the payment
	RedirectURL string `json:"RedirectURL"`
}

// The type of web direct debit
type DirectDebit string

const (
	DirectDebitSofort  DirectDebit = "SOFORT"
	DirectDebitGiropay DirectDebit = "GIROPAY"
)

//  Create a Bankwire PayIn to your Client Credit Wallet
// It is possible to add funds to your Credit/Repudiation Wallets via a PayIn Bankwire in order to settle the funds after a dispute.
// This endpoint can be used only for this purpose. For more information check The Settement Transfer object on the doc.
type BankwireCreate struct {
	// The ID of the wallet where money will be credited REQUIRED.
	CreditedWalletId ClientWalletAlias `json:"CreditedWalletId"`
	// The declared debited funds REQUIRED.
	DeclaredDebitedFunds Money `json:"DeclaredDebitedFunds"`
}

// An alias for a client wallet - made up from the FundsType and the Currency
// - "FEES_EUR" would therefore give the client's FEES wallet for EUR
type ClientWalletAlias string

const (
	ClientWalletAliasFeeEUR    ClientWalletAlias = "FEES_EUR"
	ClientWalletAliasCreditEUR ClientWalletAlias = "CREDIT_EUR"
	ClientWalletAliasFeeUSD    ClientWalletAlias = "FEES_USD"
	ClientWalletAliasCreditUSD ClientWalletAlias = "CREDIT_USD"
	ClientWalletAliasFeeGBP    ClientWalletAlias = "FEES_GBP"
	ClientWalletAliasCreditGBP ClientWalletAlias = "CREDIT_GBP"
	ClientWalletAliasOthers    ClientWalletAlias = "..." // ? is it for some other currencies ?
)

// The Pay-in web extended view is a view to get more details about the card used to process a payin web.
type WebExtended struct {
	// The item's ID
	ID string `json:"Id"`
	// The type of payin
	Payment Payment `json:"Payment"`
	// When the transaction happened
	ExecutionDate int64 `json:"ExecutionDate"`
	// The expiry date of the card - must be in format MMYY
	ExpirationDate string `json:"ExpirationDate"`
	// A partially obfuscated version of the credit card number
	Alias string `json:"Alias"`
	// The type of card . The card type is optional, but the default parameter is "CB_VISA_MASTERCARD" .
	CardType CardType `json:"CardType"`
	// The Country of the Address
	Country country.ISO2 `json:"Country"`
	// A unique representation of a 16-digits card number
	Fingerprint string `json:"Fingerprint"`
}

type DirectDebitDirect struct {
	ID           string `json:"Id"`
	CreationDate int64  `json:"CreationDate"`
	Ta           string `json:"Tag"`
	// The ID of a Mandate
	MandateID string `json:"MandateID"`
	// The date the user will be charged. Note that for direct debit payments,
	// it will take one more day more the payment becomes successful.
	ChargeDate int64 `json:"ChargeDate"`
	// A custom description to appear on the user's bank statement.
	// It can be up to 100 characters long, and can only include alphanumeric characters or spaces.
	// See here for important info and note that this functionality is only available for SEPA payments.
	StatementDescriptor string `json:"StatementDescriptor"`
}

type DirectDebitDirectCreate struct {
	// A user's ID REQUIRED.
	AuthorId string `json:"AuthorId"`
	// The user ID who is credited (defaults to the owner of the wallet) OPTIONAL.
	CreditedUserId string `json:"CreditedUserId,omitempty"`
	// The ID of the wallet where money will be credited REQUIRED.
	CreditedWalletId string `json:"CreditedWalletId"`
	// Information about the funds that are being debited REQUIRED.
	DebitedFunds Money `json:"DebitedFunds"`
	// Information about the fees that were taken by the client REQUIRED
	// for this transaction (and were hence transferred to the Client's platform wallet) .
	Fees Money `json:"Fees"`
	// The ID of a Mandate REQUIRED.
	MandateId string `json:"MandateId"`
	// A custom description to appear on the user's bank statement OPTIONAL.
	// It can be up to 100 characters long, and can only include alphanumeric characters or spaces.
	// See here for important info and note that this functionality is only available for SEPA payments.
	StatementDescriptor string `json:"StatementDescriptor,omitempty"`
}

// PayInBankwireToClientCreditWalletCreate make it is possible to add funds to your Credit/Repudiation Wallets via a PayIn Bankwire
// in order to settle the funds after a dispute.
// This endpoint can be used only for this purpose. For more information check The Settement Transfer object on the doc
type BankwireToClientCreditWalletCreate struct {
	// The ID of the wallet where money will be credited REQUIRED.
	CreditedWalletID ClientWalletAlias // ?? this should be a string for ID ??.
	// The declared debited funds REQUIRED.
	DeclaredDebitedFunds Money
}
