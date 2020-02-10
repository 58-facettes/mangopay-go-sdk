package model

import "time"

func String(str string) *string {
	return &str
}

func Time(t time.Time) *int64 {
	res := t.UTC().Unix()
	return &res
}

func Int(i int) *int {
	return &i
}
