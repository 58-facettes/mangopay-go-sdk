package model

// BankAccount is the old version of BankAccountV2
// ?? can we name it BankAccountAccess ??
type BankAccount struct {
	ID string `json:"Id"`
	// The object owner's UserId
	UserID string `json:"UserId"`
	// The name of the owner of the bank account
	OwnerName string `json:"OwnerName"`
	// The address of the owner of the bank account
	OwnerAddress Address `json:"OwnerAddress"`
	// The type of bank account
	Type BankAccountType `json:"Type"`
	// Whether the bank account is active or not
	Active bool `json:"Active"`
}

// BankAccountIBAN is the IBAN informations to create bankwire.
type BankAccountIBAN struct {
	ID string `json:"Id"`
	// The type of bank account
	Type BankAccountType `json:"Type"`
	// The address of the owner of the bank account
	OwnerAddress Address `json:"OwnerAddress"`
	// The name of the owner of the bank account
	OwnerName string `json:"OwnerName"`
	// The object owner's UserId
	UserID string `json:"UserId"`
	// Whether the bank account is active or not
	Active bool `json:"Active"`
	// The IBAN of the bank account
	IBAN string `json:"IBAN"`
	// The BIC of the bank account
	BIC string `json:"BIC"`
}

// BankAccountIBANCreate contain all data in order to save som account IBAN informations.
type BankAccountIBANCreate struct {
	// The address of the owner of the bank account REQUIRED.
	OwnerAddress Address `json:"OwnerAddress"`
	// The name of the owner of the bank account REQUIRED.
	OwnerName string `json:"OwnerName"`
	// The IBAN of the bank account REQUIRED.
	IBAN string `json:"IBAN"`
	// The BIC of the bank account OPTIONAL.
	BIC string `json:"BIC,omitempty"`
}

// ?? We miss the regular BankAccountUS struct ??
// it is not yet in the documentation.
// ?? end ??

// BankAccountUSCreate is the payload used to create an account in US.
type BankAccountUSCreate struct {
	// The address of the owner of the bank account REQUIRED.
	OwnerAddress Address `json:"OwnerAddress"`
	// The name of the owner of the bank account REQUIRED.
	OwnerName string `json:"OwnerName"`
	// The account number of the bank account. US account numbers must be digits only REQUIRED.
	AccountNumber string `json:"AccountNumber"`
	// The ABA of the bank account. Must be numbers only, and 9 digits long REQUIRED
	ABA string `json:"ABA"`
	// The type of account OPTIONAL
	DepositAccount DepositAccount `json:"DepositAccount,omitempty"`
}

// DepositAccount is the type of account
type DepositAccount string

const (
	// DepositAccountChecking is for checking bank account.
	DepositAccountChecking DepositAccount = "CHECKING"
	// DepositAccountSavings is for saving bank account.
	DepositAccountSavings DepositAccount = "SAVINGS"
)

// BankAccountCA represents the bank account informations for Canada.
type BankAccountCA struct {
	ID string `json:"Id"`
	// The type of bank account
	Type BankAccountType `json:"Type"`
	// The address of the owner of the bank account
	OwnerAddress Address `json:"OwnerAddress"`
	// The name of the owner of the bank account
	OwnerName string `json:"OwnerName"`
	// The object owner's UserId
	UserID string `json:"UserId"`
	// Whether the bank account is active or not
	Active bool `json:"Active"`
	// The institution number of the bank account.
	// Must be numbers only, and 3 or 4 digits long
	InstitutionNumber string `json:"InstitutionNumber"`
	// The account number of the bank account.
	// Must be numbers only. Canadian account numbers must be a maximum of 20 digits.
	AccountNumber string `json:"AccountNumber"`
	// The branch code of the bank where the bank account.
	// Must be numbers only, and 5 digits long
	BranchCode string `json:"BranchCode"`
	// The name of the bank where the account is held.
	// Must be letters or numbers only and maximum 50 characters long.
	BankName string `json:"BankName"`
}

// BankAccountCACreate is the payload used to create an new bank account in Canada.
type BankAccountCACreate struct {
	// The address of the owner of the bank account REQUIRED.
	OwnerAddress Address `json:"OwnerAddress"`
	// The name of the owner of the bank account REQUIRED.
	OwnerName string `json:"OwnerName"`
	// The branch code of the bank where the bank account REQUIRED.
	// Must be numbers only, and 5 digits long.
	BranchCode string `json:"BranchCode"`
	// The institution number of the bank account REQUIRED.
	// Must be numbers only, and 3 or 4 digits long.
	InstitutionNumber string `json:"InstitutionNumber"`
	// The account number of the bank account REQUIRED.
	// Must be numbers only. Canadian account numbers must be a maximum of 20 digits.
	AccountNumber string `json:"AccountNumber"`
	// The name of the bank where the account is held REQUIRED.
	// Must be letters or numbers only and maximum 50 characters long.
	BankName string `json:"BankName"`
}

// BankAccountGB represents the bank account in Great Britain.
type BankAccountGB struct {
	ID string `json:"Id"`
	// The type of bank account
	Type BankAccountType `json:"Type"`
	// The address of the owner of the bank account
	OwnerAddress Address `json:"OwnerAddress"`
	// The name of the owner of the bank account
	OwnerName string `json:"OwnerName"`
	// The object owner's UserId
	UserID string `json:"UserId"`
	// Whether the bank account is active or not
	Active bool `json:"Active"`
	// The sort code of the bank account. Must be numbers only, and 6 digits long
	SortCode string `json:"SortCode"`
	// The account number of the bank account. Must be numbers only. GB account numbers must be 8 digits long.
	AccountNumber string `json:"AccountNumber"`
}

// BankAccountGBCreate is the payload used to create a new bank account in Great Britain.
type BankAccountGBCreate struct {
	// The address of the owner of the bank account REQUIRED.
	OwnerAddress Address `json:"OwnerAddress"`
	// The name of the owner of the bank account REQUIRED.
	OwnerName string `json:"OwnerName"`
	// The sort code of the bank account. Must be numbers only, and 6 digits long REQUIRED.
	SortCode string `json:"SortCode"`
	// The account number of the bank account. Must be numbers only. GB account numbers must be 8 digits long. REQUIRED.
	AccountNumber string `json:"AccountNumber"`
}

// BankAccountOther represent a bank account fron another country.
type BankAccountOther struct {
	ID string `json:"ID"`
	// The type of bank account
	Type BankAccountType `json:"Type"`
	// The address of the owner of the bank account
	OwnerAddress Address `json:"OwnerAddress"`
	// The name of the owner of the bank account
	OwnerName string `json:"OwnerName"`
	// The object owner's UserId
	UserID string `json:"UserId"`
	// Whether the bank account is active or not
	Active bool `json:"Active"`
}

// BankAccountOtherCreate is the payload used to create a new bank account in an other country.
type BankAccountOtherCreate struct {
	// The address of the owner of the bank account REQUIRED.
	OwnerAddress Address `json:"OwnerAddress"`
	// The name of the owner of the bank account REQUIRED.
	OwnerName string `json:"OwnerName"`
	// The Country of the Address REQUIRED.
	Country string `json:"Country"`
	// The BIC of the bank account REQUIRED.
	BIC string `json:"BIC"`
	// The account number of the bank account REQUIRED.
	// Must be numbers only. Canadian account numbers must be a maximum of 20 digits.
	AccountNumber string `json:"AccountNumber"`
}
