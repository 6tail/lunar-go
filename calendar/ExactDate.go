package calendar

import (
	"time"
)

func NewExactDateFromYmdHms(year int, month int, day int, hour int, minute int, second int) time.Time {
	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
}

func NewExactDateFromYmd(year int, month int, day int) time.Time {
	return NewExactDateFromYmdHms(year, month, day, 0, 0, 0)
}

func NewExactDateFromDate(date time.Time) time.Time {
	return NewExactDateFromYmdHms(date.Year(), int(date.Month()), date.Day(), date.Hour(), date.Minute(), date.Second())
}
