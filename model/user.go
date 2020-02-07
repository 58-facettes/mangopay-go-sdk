package model

type User struct {
	// ID is the user's unique identification.
	ID int `json:"id"`
	// PersonType is the type of user.
	PersonType PersonType `json:"PersonType"`
	// Email is the person's email address (not more than 12 consecutive numbers) - must be a valid email.
	Email string `json:"Email"`
	// KYCLevel More info here.
	KYCLevel KYCLevelType `json:"KYCLevel"`
	// Tag is a custom data that you can add to this item.
	Tag string `json:"Tag"`
	// CreationDate is when the item was created.
	CreationDate int64 `json:"CreationDate"`
}

type PersonType string

const (
	PersonNatural PersonType = "NATURAL"
	PersonLegal   PersonType = "LEGAL"
)

// KYCLevelType there are two levels of user verification, also called API levels: Light (default) and Regular.
// Light (default) verification requires basic information provided during the user creation process.
// Regular verification requires identity proof which allows users to handle unlimited funds freely.
type KYCLevelType string

const (
	KYCLevelLight   KYCLevelType = "LIGHT"
	KYCLevelRegular KYCLevelType = "REGULAR"
)
