package model

// PayOut it the response payload.
type PayOut struct {
	// ?? not describe in the documentation yet ??
	ID           string `json:"Id"`
	Tag          string `json:"Tag"`
	CreationDate int64  `json:"CreationDate"`
	// ?? end ??
	// An ID of a Bank Account
	BankAccountID string `json:"BankAccountId"`
	// A custom reference you wish to appear on the user’s bank statement (your Client Name is already shown).
	// We advise you not to add more than 12 characters.
	BankWireRef string `json:"BankWireRef"`
	// The type of Pay out.
	Payment PayOutPayment `json:"PayOutPaymentType"`
}

// PayOutPayment is the payment type of the payout.
type PayOutPayment string

const (
	// PayOutPaymentBankWire is used for bank wire.
	PayOutPaymentBankWire PayOutPayment = "BANK_WIRE"
)

// PayOutCreate is used to create a Payout.
type PayOutCreate struct {
	// A user's ID REQUIRED.
	AuthorID string `json:"AuthorId"`
	// Information about the funds that are being debited REQUIRED.
	DebitedFunds Money `json:"DebitedFunds"`
	// Information about the fees that were taken by the client REQUIRED
	// for this transaction (and were hence transferred to the Client's platform wallet).
	Fees Money `json:"Fees"`
	// An ID of a Bank Account REQUIRED.
	BankAccountID string `json:"BankAccountId"`
	// The ID of the wallet that was debited REQUIRED.
	DebitedWalletID string `json:"DebitedWalletId"`
	// A custom reference you wish to appear on the user’s bank statement OPTIONAL
	// (your Client Name is already shown). We advise you not to add more than 12 characters.
	BankWireRef string `json:"BankWireRef,omitempty"`
}
