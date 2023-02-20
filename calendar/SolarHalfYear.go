package calendar

import (
	"container/list"
	"fmt"
	"math"
	"time"
)

const MONTH_IN_HALF_YEAR = 6

// SolarHalfYear 阳历半年
type SolarHalfYear struct {
	year  int
	month int
}

func NewSolarHalfYear() *SolarHalfYear {
	return NewSolarHalfYearFromDate(time.Now().Local())
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
	return fmt.Sprintf("%d.%d", solarHalfYear.year, solarHalfYear.GetIndex())
}

func (solarHalfYear *SolarHalfYear) ToFullString() string {
	index := "下"
	if solarHalfYear.GetIndex() == 1 {
		index = "上"
	}
	return fmt.Sprintf("%d年%s半年", solarHalfYear.year, index)
}

func (solarHalfYear *SolarHalfYear) Next(halfYears int) *SolarHalfYear {
	m := NewSolarMonthFromYm(solarHalfYear.year, solarHalfYear.month).Next(MONTH_IN_HALF_YEAR * halfYears)
	return NewSolarHalfYearFromYm(m.GetYear(), m.GetMonth())
}
