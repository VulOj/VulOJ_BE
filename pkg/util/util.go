package util

import (
	"log"
	"time"
)

func CheckError(err error) bool {
	if err != nil {
		log.Fatal(err.Error())
		return true
	}
	return false
}

func ConvertShanghaiTimeZone(t time.Time) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}
