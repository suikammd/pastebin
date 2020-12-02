package util

import (
	"time"
)

const (
	code62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length = len(code62)
)

func Encode() string {
	curTime := time.Now().Nanosecond()

	s := ""
	for ; curTime > 0; curTime /= length {
		s = string(code62[curTime %length]) + s
	}
	return s
}
