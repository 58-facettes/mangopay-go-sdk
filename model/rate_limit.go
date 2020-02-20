package model

import "fmt"

// Rate is the type used for rate limit value.
type Rate uint8

const (
	// RateLimit15Minutes stands for 15 minutes.
	RateLimit15Minutes Rate = iota + 1
	// RateLimit30Minutes stands for 30 minutes.
	RateLimit30Minutes
	// RateLimit60Minutes stands for 60 minutes.
	RateLimit60Minutes
	// RateLimit24Mours stands for 24 hours.
	RateLimit24Mours
)

// RateLimit is in the header response of the Mangopay service ??
// it uses some duplicated datas fiedds that are in the header
// so we have to hijac the response before the map[string][]string get the Header values.
// ?? end ??
type RateLimit struct {
	Limit1          int `json:"X-RateLimit-1"`
	Limit2          int `json:"X-RateLimit-2"`
	Limit3          int `json:"X-RateLimit-3"`
	Limit4          int `json:"X-RateLimit-4"`
	LimitRemaining1 int `json:"X-RateLimit-Remaining-1"`
	LimitRemaining2 int `json:"X-RateLimit-Remaining-2"`
	LimitRemaining3 int `json:"X-RateLimit-Remaining-3"`
	LimitRemaining4 int `json:"X-RateLimit-Remaining-4"`
	LimitReset1     int `json:"X-RateLimit-Reset-1"`
	LimitReset2     int `json:"X-RateLimit-Reset-2"`
	LimitReset3     int `json:"X-RateLimit-Reset-3"`
	LimitReset4     int `json:"X-RateLimit-Reset-4"`
}

const (
	rateFormatData = "rateLimit: %v rateRemaining: %v rateReset: %v"
)

// GetData gives the rates limits from the given limit value.
func (rl *RateLimit) GetData(r Rate) string {
	switch r {
	case RateLimit15Minutes:
		return fmt.Sprintf(rateFormatData, rl.Limit1, rl.LimitRemaining1, rl.LimitReset1)
	case RateLimit30Minutes:
		return fmt.Sprintf(rateFormatData, rl.Limit2, rl.LimitRemaining2, rl.LimitReset2)
	case RateLimit60Minutes:
		return fmt.Sprintf(rateFormatData, rl.Limit3, rl.LimitRemaining3, rl.LimitReset3)
	case RateLimit24Mours:
		return fmt.Sprintf(rateFormatData, rl.Limit4, rl.LimitRemaining4, rl.LimitReset4)
	default:
		return fmt.Sprintf(rateFormatData, rl.Limit1, rl.LimitRemaining1, rl.LimitReset1)
	}
}

func (rl *RateLimit) String() string {
	return fmt.Sprintf("ratelimit-1: %v, rateLimit-2: %v, rateLimit-3: %v, rateLimit-4: %v, "+
		"limitRemaining-1: %v, limitRemaining-2: %v, limitRemaining-3: %v, limitRemaining-4: %v, "+
		"limitReset-1: %v, LimitReset-2: %v, LimitReset-3: %v, LimitReset-4: %v",
		rl.Limit1, rl.Limit2, rl.Limit3, rl.Limit4,
		rl.LimitRemaining1, rl.LimitRemaining2, rl.LimitRemaining3, rl.LimitRemaining4,
		rl.LimitReset1, rl.LimitReset2, rl.LimitReset3, rl.LimitReset4)
}
