package calendar

import (
	"container/list"
	"fmt"
	"github.com/6tail/lunar-go/SolarUtil"
	"time"
)

// SolarMonth 阳历月
type SolarMonth struct {
	year  int
	month int
}

func NewSolarMonth() *SolarMonth {
	return NewSolarMonthFromDate(time.Now().Local())
}

func NewSolarMonthFromYm(year int, month int) *SolarMonth {
	solarMonth := new(SolarMonth)
	solarMonth.year = year
	solarMonth.month = month
	return solarMonth
}

func NewSolarMonthFromDate(date time.Time) *SolarMonth {
	return NewSolarMonthFromYm(date.Year(), int(date.Month()))
}

func (solarMonth *SolarMonth) GetYear() int {
	return solarMonth.year
}

func (solarMonth *SolarMonth) GetMonth() int {
	return solarMonth.month
}

func (solarMonth *SolarMonth) GetDays() *list.List {
	firstDay := NewSolarFromYmd(solarMonth.year, solarMonth.month, 1)
	l := list.New()
	l.PushBack(firstDay)
	days := SolarUtil.GetDaysOfMonth(solarMonth.year, solarMonth.month)
	for i := 1; i < days; i++ {
		l.PushBack(firstDay.NextDay(i))
	}
	return l
}

func (solarMonth *SolarMonth) GetWeeks(start int) *list.List {
	l := list.New()
	week := NewSolarWeekFromYmd(solarMonth.year, solarMonth.month, 1, start)
	for {
		l.PushBack(week)
		week = week.Next(1, false)
		firstDay := week.GetFirstDay()
		if firstDay.GetYear() > solarMonth.year || firstDay.GetMonth() > solarMonth.month {
			break
		}
	}
	return l
}

func (solarMonth *SolarMonth) String() string {
	return fmt.Sprintf("%d-%d", solarMonth.year, solarMonth.month)
}

func (solarMonth *SolarMonth) ToFullString() string {
	return fmt.Sprintf("%d年%d月", solarMonth.year, solarMonth.month)
}

func (solarMonth *SolarMonth) Next(months int) *SolarMonth {
	n := 1
	m := months
	if months < 0 {
		n = -1
		m = -months
	}
	y := solarMonth.year + m/12*n
	m = solarMonth.month + m%12*n
	if m > 12 {
		m -= 12
		y++
	} else if m < 1 {
		m += 12
		y--
	}
	return NewSolarMonthFromYm(y, m)
}
