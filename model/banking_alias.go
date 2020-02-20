package model

import (
	"github.com/58-facettes/mangopay-go-sdk/model/country"
)

// BankingAlias allow you to create a way to pay funds directly into a wallet,
// without having to declare the payin beforehand (unlike a traditional payin bankwire).
// For example, if you create an IBAN banking alias for a wallet, you'll be given a unique IBAN and BIC for this wallet. Any funds that we receive for this IBAN and BIC will be automatically credited to the wallet. You should be aware that you are unable to add Fees to a payin created via a banking alias.
type BankingAlias struct {
	// The user ID who is credited (defaults to the owner of the wallet)
	CreditedUserID string `json:"CreditedUserId"`
	// The ID of a wallet
	WalletID string `json:"WalletId"`
	// The type of banking alias (note that only IBAN is available at present)
	Type BankingAliasType `json:"Type"`
	// The Country of the Address
	Country country.ISO2 `json:"Country"`
	// The owner of the wallet/banking alias
	OwnerName string `json:"OwnerName"`
	// Whether the banking alias is active or not
	Active bool `json:"Active"`
}

// BankingAliasType is the type of banking alias. Currently only IBAN is supported.
type BankingAliasType string

const (
	// BankingAliasIBAN stands for BankingAliasType IBAN.
	BankingAliasIBAN BankingAliasType = "IBAN"
)

// BankingAliasing is for Banking IBAN details.
type BankingAliasing struct {
	// The user ID who is credited (defaults to the owner of the wallet)
	CreditedUserID string `json:"CreditedUserId"`
	// The ID of a wallet
	WalletID string `json:"WalletId"`
	// The type of banking alias (note that only IBAN is available at present)
	Type BankingAliasType `json:"Type"`
	// The Country of the Address
	Country country.ISO2 `json:"Country"`
	// The owner of the wallet/banking alias
	OwnerName string `json:"OwnerName"`
	// Whether the banking alias is active or not
	Active bool `json:"Active"`
	// The IBAN of the banking alias
	IBAN string `json:"IBAN"`
	// The BIC of the banking alias
	BIC string `json:"BIC"`
}

// BankingAliasCreate is creating a BankingAlias.
type BankingAliasCreate struct {
	// The user ID who is credited (defaults to the owner of the wallet) OPTIONAL.
	CreditedUserID string `json:"CreditedUserId"`
	// The owner of the wallet/banking alias REQUIRED.
	OwnerName string `json:"OwnerName"`
	// The country format for the banking alias. Note that we only provide FR and LU IBAN. REQUIRED
	Country BankingAliasCountry `json:"Country"`
}

// BankingAliasCountry is the country format for the banking alias.
// Note that FR is not currently available, but will be shortly.
type BankingAliasCountry string

const (
	// BankingAliasCountryLU stands for Luxembourg.
	BankingAliasCountryLU BankingAliasCountry = "LU"
	// BankingAliasCountryFR stands for France.
	BankingAliasCountryFR BankingAliasCountry = "FR"
)
