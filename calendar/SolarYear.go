package calendar

import (
	"container/list"
	"fmt"
	"time"
)

const MONTH_IN_YEAR = 12

// SolarYear 阳历年
type SolarYear struct {
	year int
}

func NewSolarYear() *SolarYear {
	return NewSolarYearFromDate(time.Now().Local())
}

func NewSolarYearFromYear(year int) *SolarYear {
	solarYear := new(SolarYear)
	solarYear.year = year
	return solarYear
}

func NewSolarYearFromDate(date time.Time) *SolarYear {
	return NewSolarYearFromYear(date.Year())
}

func (solarYear *SolarYear) GetYear() int {
	return solarYear.year
}

func (solarYear *SolarYear) GetMonths() *list.List {
	firstMonth := NewSolarMonthFromYm(solarYear.year, 1)
	l := list.New()
	l.PushBack(firstMonth)
	for i := 1; i < MONTH_IN_YEAR; i++ {
		l.PushBack(firstMonth.Next(i))
	}
	return l
}

func (solarYear *SolarYear) String() string {
	return fmt.Sprintf("%d", solarYear.year)
}

func (solarYear *SolarYear) ToFullString() string {
	return fmt.Sprintf("%d年", solarYear.year)
}

func (solarYear *SolarYear) Next(years int) *SolarYear {
	return NewSolarYearFromYear(solarYear.year + years)
}
