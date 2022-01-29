package calendar

import (
	"github.com/6tail/lunar-go/SolarUtil"
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

func GetDaysBetween(ay int, am int, ad int, by int, bm int, bd int) int {
	n := 0
	if ay == by {
		n = SolarUtil.GetDaysInYear(by, bm, bd) - SolarUtil.GetDaysInYear(ay, am, ad)
	} else if ay > by {
		days := SolarUtil.GetDaysOfYear(by) - SolarUtil.GetDaysInYear(by, bm, bd)
		for i := by + 1; i < ay; i++ {
			days += SolarUtil.GetDaysOfYear(i)
		}
		days += SolarUtil.GetDaysInYear(ay, am, ad)
		n = -days
	} else {
		days := SolarUtil.GetDaysOfYear(ay) - SolarUtil.GetDaysInYear(ay, am, ad)
		for i := ay + 1; i < by; i++ {
			days += SolarUtil.GetDaysOfYear(i)
		}
		days += SolarUtil.GetDaysInYear(by, bm, bd)
		n = days
	}
	return n
}

func GetDaysBetweenDate(date0 time.Time, date1 time.Time) int {
	return GetDaysBetween(date0.Year(), int(date0.Month()), date0.Day(), date1.Year(), int(date1.Month()), date1.Day())
}
