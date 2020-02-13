package model

import "time"

func Time(t time.Time) int64 {
	return t.UTC().Unix()
}
