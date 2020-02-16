package model

import (
	"github.com/58-facettes/mangopay-go-sdk/model/currency"
)

// CardRegistration you need to register a card in order to process a Direct PayIn.
// Card registration enables you to tokenize a Card.
// These are the steps to follow:
// - Create a CardRegistration Object (1. 2. & 3. in the diagram).
// - Get the PreRegistrationData , CardRegistrationURL and AccessKey (4. in the diagram).
// - The user posts PreRegistrationData, AccessKey and card details through a form to the CardRegistrationURL (5. in the diagram).
// - Get a RegistrationData back (6. in the diagram).
// - Edit the CardRegistration Object with the RegistrationData just received (7.8. in the diagram).
// - Get the CardId ready to use into the Direct PayIn Object (9. in the diagram).
type CardRegistration struct {
	// ID of the current card's registration.
	ID string `json:"Id"` // ?? new ??
	// The object owner's UserId.
	UserID string `json:"UserId"`
	// The currency - should be ISO_4217 format.
	Currency currency.ISO3 `json:"Currency"`
	// A special key to use when registering a card.
	AccessKey string `json:"AccessKey"`
	// A specific value to pass to the CardRegistrationURL.
	PreregistrationData string `json:"PreregistrationData"`
	// The URL to submit the card details form to.
	CardRegistrationURL string `json:"CardRegistrationURL"`
	// Having registered a card, this confirmation hash needs to be updated to the card item.
	RegistrationData string `json:"RegistrationData"`
	// The type of card . The card type is optional, but the default parameter is "CB_VISA_MASTERCARD".
	CardType CardType `json:"CardType"`
	// The ID of a card.
	CardID string `json:"CardId"`
	// The result code.
	ResultCode string `json:"ResultCode"`
	// A verbal explanation of the ResultCode.
	ResultMessage string `json:"ResultMessage"`
	// Status of the card registration.
	Status CardStatus `json:"Status"`
	// Tag of the card.
	Tag string `json:"Tag"` // ?? new ??
	// CreationDate
	CreationDate int64 `json:"CreationDate"` // ?? new ??
}

// CardRegistrationCreate is the payload used for registering a new Card.
type CardRegistrationCreate struct {
	// The object owner's UserId (REQUIRED).
	UserID string `json:"UserId"`
	// The currency - should be ISO_4217 format (REQUIRED).
	Currency currency.ISO3 `json:"Currency"`
	// The type of card. The card type is optional, but the default parameter is "CB_VISA_MASTERCARD" (OPTIONAL).
	CardType CardType `json:"CardType,omitempty"`
}

// CardType is the kind of card.
type CardType string

const (
	// CardCBvisaMastercard is used for type CB VISA MASTERCARD
	CardCBvisaMastercard CardType = "CB_VISA_MASTERCARD"
	// CardDiners is used for Diners card.
	CardDiners CardType = "DINERS"
	// CardMasterpass is used for Masterpass card.
	CardMasterpass CardType = "MASTERPASS"
	// CardAmex is used for Amex card.
	CardAmex CardType = "AMEX "
	// CardMaestro is used for Maestro card.
	CardMaestro CardType = "MAESTRO"
	// CardP24 is used for P24 card.
	CardP24 CardType = "P24"
	// CardIdeal is used for IDEAL card.
	CardIdeal CardType = "IDEAL"
	// CardBcmc is used for BCMC card.
	CardBcmc CardType = "BCMC"
	// CardPaylib is used for Paylib card.
	CardPaylib CardType = "PAYLIB"
)

// CardStatus is the kind of status of a card.
type CardStatus string

const (
	// CardStatusCreated is used for a created status.
	CardStatusCreated CardStatus = "CREATED"
	// CardStatusValidated is used for a validated status.
	CardStatusValidated CardStatus = "VALIDATED"
	// CardStatusError is used for a error status.
	CardStatusError CardStatus = "ERROR"
)

// CardRegistrationUpdate is a payload used for updating a card's registration.
type CardRegistrationUpdate struct {
	// Having registered a card, this confirmation hash needs to be updated to the card item (OPTIONAL).
	RegistrationData string `json:"RegistrationData"`
}

// Card represents the card.
type Card struct {
	// The expiry date of the card - must be in format MMYY
	ExpirationDate int64 `json:"ExpirationDate"`
	// A partially obfuscated version of the credit card number
	Alias string `json:"Alias"`
	// The provider of the card
	CardProvider string `json:"CardProvider"`
	// The type of card . The card type is optional, but the default parameter is "CB_VISA_MASTERCARD" .
	CardType CardType `json:"CardType"`
	// The Country of the Address
	Country string `json:"Country"`
	// The card product type - more info
	Product string `json:"Product"`
	// The BankCode of the card.
	BankCode string `json:"BankCode"`
	// Whether the card is active or not
	Active bool `json:"Active"`
	// The currency - should be ISO_4217 format
	Currency currency.ISO3 `json:"Currency"`
	// Whether the card is valid or not. Once they process (or attempt to process) a payment with the card we are able to indicate if it is "valid" or "invalid". If they didn’t process a payment yet the "Validity" stay at "unknown".
	Validity CardValidity `json:"Validity"`
	// A unique representation of a 16-digits card number
	Fingerprint string `json:"Fingerprint"`
}

// CardValidity is for the validity of the Card.
type CardValidity string

const (
	// CardValidityUnknown is used for a validity status unknown.
	CardValidityUnknown CardValidity = "UNKNOWN"
	// CardValidityValid is used for a validity status valid.
	CardValidityValid CardValidity = "VALID"
	// CardValidityInvalid is used for a validity status invalid.
	CardValidityInvalid CardValidity = "INVALID"
)

// CardDeactivate let you desactivate or activate a Card.
type CardDeactivate struct {
	// Whether the card is active or not (OPTIONAL).
	Active bool `json:"Active"`
}
