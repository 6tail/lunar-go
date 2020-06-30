package calendar

import (
	"container/list"
	"math"
	"strconv"
	"time"
)

const MONTH_IN_HALF_YEAR = 6

type SolarHalfYear struct {
	year  int
	month int
}

func NewSolarHalfYear() *SolarHalfYear {
	return NewSolarHalfYearFromDate(time.Now())
}

func NewSolarHalfYearFromYm(year int, month int) *SolarHalfYear {
	solarHalfYear := new(SolarHalfYear)
	solarHalfYear.year = year
	solarHalfYear.month = month
	return solarHalfYear
}

func NewSolarHalfYearFromDate(date time.Time) *SolarHalfYear {
	return NewSolarHalfYearFromYm(date.Year(), int(date.Month()))
}

func (solarHalfYear *SolarHalfYear) GetYear() int {
	return solarHalfYear.year
}

func (solarHalfYear *SolarHalfYear) GetMonth() int {
	return solarHalfYear.month
}

func (solarHalfYear *SolarHalfYear) GetIndex() int {
	return int(math.Ceil(float64(solarHalfYear.month) / MONTH_IN_HALF_YEAR))
}

func (solarHalfYear *SolarHalfYear) GetMonths() *list.List {
	l := list.New()
	index := solarHalfYear.GetIndex() - 1
	for i := 0; i < MONTH_IN_HALF_YEAR; i++ {
		l.PushBack(NewSolarMonthFromYm(solarHalfYear.year, MONTH_IN_HALF_YEAR*index+i+1))
	}
	return l
}

func (solarHalfYear *SolarHalfYear) String() string {
	return strconv.Itoa(solarHalfYear.year) + "." + strconv.Itoa(solarHalfYear.GetIndex())
}

func (solarHalfYear *SolarHalfYear) ToFullString() string {
	index := "下"
	if solarHalfYear.GetIndex() == 1 {
		index = "上"
	}
	return strconv.Itoa(solarHalfYear.year) + "年" + index + "半年"
}

func (solarHalfYear *SolarHalfYear) Next(halfYears int) *SolarHalfYear {
	if 0 == halfYears {
		return NewSolarHalfYearFromYm(solarHalfYear.year, solarHalfYear.month)
	}
	c := time.Date(solarHalfYear.year, time.Month(solarHalfYear.month), 1, 0, 0, 0, 0, time.Local)
	c.AddDate(0, MONTH_IN_HALF_YEAR, 0)
	return NewSolarHalfYearFromDate(c)
}
