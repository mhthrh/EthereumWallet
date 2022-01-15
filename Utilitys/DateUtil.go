package Utilitys

import "time"

func GetDate(format string) time.Time {
	a, _ := time.Parse(format, time.Now().String())
	return a
}
