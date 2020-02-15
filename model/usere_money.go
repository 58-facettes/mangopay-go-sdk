package model

// UserEmoney represent the payload for a User Emoney.
type UserEmoney struct {
	// The object owner's UserId
	UserID string `json:"userId"`
	// The amount of money that has been credited to this user
	CreditedEMoney Money `json:"CreditedEMoney"`
	// The amount of money that has been debited from this user
	DebitedEMoney Money `json:"DebitedEMoney"`
}

// UserEmoneyParam is the param that can be used to call a User Emoney.
type UserEmoneyParam struct {
	// Year is the year of the Emoney (optional).
	Year int
	// Month is the month in number (optional).
	Month int
	// Currency in ISO3 displayed (optional).
	Currency string
}
