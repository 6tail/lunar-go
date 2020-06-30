package calendar

import (
	"container/list"
	"lunar-go/SolarUtil"
	"strconv"
	"time"
)

type SolarMonth struct {
	year  int
	month int
}

func NewSolarMonth() *SolarMonth {
	return NewSolarMonthFromDate(time.Now())
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
		l.PushBack(firstDay.Next(i))
	}
	return l
}

func (solarMonth *SolarMonth) String() string {
	return strconv.Itoa(solarMonth.year) + "-" + strconv.Itoa(solarMonth.month)
}

func (solarMonth *SolarMonth) ToFullString() string {
	return strconv.Itoa(solarMonth.year) + "年" + strconv.Itoa(solarMonth.month) + "月"
}

func (solarMonth *SolarMonth) Next(months int) *SolarMonth {
	c := time.Date(solarMonth.year, time.Month(solarMonth.month), 1, 0, 0, 0, 0, time.Local)
	c.AddDate(0, months, 0)
	return NewSolarMonthFromDate(c)
}
