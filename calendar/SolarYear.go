package calendar

import (
	"container/list"
	"strconv"
	"time"
)

const MONTH_IN_YEAR = 12

type SolarYear struct {
	year int
}

func NewSolarYear() *SolarYear {
	return NewSolarYearFromDate(time.Now())
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
	return strconv.Itoa(solarYear.year)
}

func (solarYear *SolarYear) ToFullString() string {
	return strconv.Itoa(solarYear.year) + "年"
}

func (solarYear *SolarYear) Next(years int) *SolarYear {
	c := time.Date(solarYear.year, 1, 1, 0, 0, 0, 0, time.Local)
	c.AddDate(years, 0, 0)
	return NewSolarYearFromDate(c)
}
