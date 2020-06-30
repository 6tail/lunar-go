package calendar

import (
	"container/list"
	"math"
	"strconv"
	"time"
)

const MONTH_IN_SEASON = 3

type SolarSeason struct {
	year  int
	month int
}

func NewSolarSeason() *SolarSeason {
	return NewSolarSeasonFromDate(time.Now())
}

func NewSolarSeasonFromYm(year int, month int) *SolarSeason {
	solarSeason := new(SolarSeason)
	solarSeason.year = year
	solarSeason.month = month
	return solarSeason
}

func NewSolarSeasonFromDate(date time.Time) *SolarSeason {
	return NewSolarSeasonFromYm(date.Year(), int(date.Month()))
}

func (solarSeason *SolarSeason) GetYear() int {
	return solarSeason.year
}

func (solarSeason *SolarSeason) GetMonth() int {
	return solarSeason.month
}

func (solarSeason *SolarSeason) GetIndex() int {
	return int(math.Ceil(float64(solarSeason.month) / MONTH_IN_SEASON))
}

func (solarSeason *SolarSeason) GetMonths() *list.List {
	l := list.New()
	index := solarSeason.GetIndex() - 1
	for i := 0; i < MONTH_IN_SEASON; i++ {
		l.PushBack(NewSolarMonthFromYm(solarSeason.year, MONTH_IN_SEASON*index+i+1))
	}
	return l
}

func (solarSeason *SolarSeason) String() string {
	return strconv.Itoa(solarSeason.year) + "." + strconv.Itoa(solarSeason.GetIndex())
}

func (solarSeason *SolarSeason) ToFullString() string {
	return strconv.Itoa(solarSeason.year) + "年" + strconv.Itoa(solarSeason.GetIndex()) + "季度"
}

func (solarSeason *SolarSeason) Next(seassons int) *SolarSeason {
	if 0 == seassons {
		return NewSolarSeasonFromYm(solarSeason.year, solarSeason.month)
	}
	c := time.Date(solarSeason.year, time.Month(solarSeason.month), 1, 0, 0, 0, 0, time.Local)
	c.AddDate(0, MONTH_IN_SEASON*seassons, 0)
	return NewSolarSeasonFromDate(c)
}
