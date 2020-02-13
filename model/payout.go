package model

type PayOut struct {
	// ?? not describe in the documentation yet ??
	ID           string `json:"Id"`
	Tag          string `json:"Tag"`
	CreationDate int64  `json:"CreationDate"`
	// ?? end ??
	// An ID of a Bank Account
	BankAccountID string `json:"BankAccountId"`
	// A custom reference you wish to appear on the user’s bank statement (your Client Name is already shown). We advise you not to add more than 12 characters.
	BankWireRef string `json:"BankWireRef"`
	// The type of Pay out
	Payment PayOutPayment `json:"PayOutPaymentType"`
}

// The payment type of the payout
type PayOutPayment string

const (
	PayOutPaymentBankWire PayOutPayment = "BANK_WIRE"
)

type PayOutCreate struct {
	// A user's ID REQUIRED.
	AuthorId string `json:"AuthorId"`
	// Information about the funds that are being debited REQUIRED.
	DebitedFunds Money `json:"DebitedFunds"`
	// Information about the fees that were taken by the client REQUIRED
	// for this transaction (and were hence transferred to the Client's platform wallet).
	Fees Money `json:"Fees"`
	// An ID of a Bank Account REQUIRED.
	BankAccountId string `json:"BankAccountId"`
	// The ID of the wallet that was debited REQUIRED.
	DebitedWalletId string `json:"DebitedWalletId"`
	// A custom reference you wish to appear on the user’s bank statement OPTIONAL
	// (your Client Name is already shown). We advise you not to add more than 12 characters.
	BankWireRef string `json:"BankWireRef,omitempty"`
}
