package Utilitys

import (
	"time"
)

func GetDate(format string) string {

	if format == "date" {
		return time.Now().Format("02-01-2006")
	}
	if format == "time" {
		time.Now().Format("15:04:05")
	}
	return time.Now().Format("02-01-2006 15:04:05.000")
}
func DateDiff(first, second time.Time) time.Duration {

	return second.Sub(first)
}
