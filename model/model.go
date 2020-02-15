package model

import "time"

// Time simplify the usage of the timestamp in UTC used in the Mangopay services.
func Time(t time.Time) int64 {
	return t.UTC().Unix()
}
