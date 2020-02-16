package model

// The PreAuthorization Object ensures the solvency of a registered card for 7 days. The overall process is as follows:
//
// Register a card (CardRegistration)
// Create a PreAuthorization with the CardId. This allows you to charge an amount on a card
// Charge the card through the PreAuthorized PayIn object (Payins/preauthorized/direct)
// How does PreAuthorization work?
//
// Once the PreAuthorization object is created the Status is "CREATED" until 3D secure validation.
// If the authorization is successful the status is "SUCCEEDED" if it failed the status is "FAILED".
// Once Status = "SUCCEEDED" and PaymentStatus = "WAITING" you can charge the card.
// The Pay-In amount has to be less than or equal to the amount authorized.
type PreAuthorization struct {
	// A user's ID.
	AuthorID string `json:"AuthorId"`
	// Information about the funds that are being debited.
	DebitedFunds Money `json:"DebitedFunds"`
	// Status of the PreAuthorization.
	Status PreAuthorizationStatus `json:"Status"`
	// The status of the payment after the PreAuthorization.
	// You can pass the PaymentStatus from "WAITING" to "CANCELED" should you need/want to.
	PaymentStatus PaymentStatus `json:"PaymentStatus"`
	// The result code.
	ResultCode string `json:"ResultCode"`
	// A verbal explanation of the ResultCode.
	ResultMessage string `json:"ResultMessage"`
	// How the PreAuthorization has been executed.
	ExecutionType PreAuthorizationExecutionType `json:"ExecutionType"`
	// The SecureMode corresponds to '3D secure' for CB Visa and MasterCard.
	// This field lets you activate it manually.
	// The field lets you activate it automatically with "DEFAULT"
	// (Secured Mode will be activated from €50 or when MANGOPAY detects there is a higher risk ),
	// "FORCE" (if you wish to specifically force the secured mode).
	SecureMode SecureMode `json:"SecureMode"`
	// The ID of a card.
	CardID string `json:"CardId"`
	// The value is 'true' if the SecureMode was used.
	SecureModeNeeded bool `json:"SecureModeNeeded"`
	// This is the URL where to redirect users to proceed to 3D secure validation.
	SecureModeRedirectURL string `json:"SecureModeRedirectUrl"`
	// This is the URL where users are automatically redirected after 3D secure validation (if activated).
	SecureModeReturnURL string `json:"SecureModeReturnURL"`
	// The date when the payment is to be processed by.
	ExpirationDate int64 `json:"ExpirationDate"`
	// The Id of the associated PayIn.
	PayInID string `json:"PayInId"`
	// Contains every useful informations related to the user billing.
	Billing Billing `json:"Billing"`
	// Contains useful informations related to security and fraud.
	SecurityInfo SecurityInfo `json:"SecurityInfo"`
	// The language to use for the mandate confirmation page - needs to be the ISO code of the language.
	Culture CultureCode `json:"Culture"`
}

// PreAuthorizationStatus is the status of the preauthorization.
type PreAuthorizationStatus string

const (
	// PreAuthorizationStatusCreated is for a pre-authorization created.
	PreAuthorizationStatusCreated PreAuthorizationStatus = "CREATED"
	// PreAuthorizationStatusSucceeded is for a pre-authorization succeded.
	PreAuthorizationStatusSucceeded PreAuthorizationStatus = "SUCCEEDED"
	// PreAuthorizationStatusFailed is for a pre-authorization failed.
	PreAuthorizationStatusFailed PreAuthorizationStatus = "FAILED"
)

// PaymentStatus is the status of the payment for a preauthorization.
type PaymentStatus string

const (
	// PaymentStatusWaiting is for the payment status waiting.
	PaymentStatusWaiting PaymentStatus = "WAITING"
	// PaymentStatusCanceled is for the payment status canceled.
	PaymentStatusCanceled PaymentStatus = "CANCELED"
	// PaymentStatusExpired is for the payment status expired.
	PaymentStatusExpired PaymentStatus = "EXPIRED"
	// PaymentStatusValidated is for the payment status validated.
	PaymentStatusValidated PaymentStatus = "VALIDATED"
)

// PreAuthorizationExecutionType is the execution type for a pre-authorization.
type PreAuthorizationExecutionType string

const (
	// PreAuthorizationExecutionDirect is the execution type direct for a pre-autorization.
	PreAuthorizationExecutionDirect PreAuthorizationExecutionType = "DIRECT"
)
