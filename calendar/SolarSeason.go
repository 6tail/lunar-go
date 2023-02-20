package calendar

import (
	"container/list"
	"fmt"
	"math"
	"time"
)

const MONTH_IN_SEASON = 3

// SolarSeason 阳历季度
type SolarSeason struct {
	year  int
	month int
}

func NewSolarSeason() *SolarSeason {
	return NewSolarSeasonFromDate(time.Now().Local())
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
	return fmt.Sprintf("%d.%d", solarSeason.year, solarSeason.GetIndex())
}

func (solarSeason *SolarSeason) ToFullString() string {
	return fmt.Sprintf("%d年%d季度", solarSeason.year, solarSeason.GetIndex())
}

func (solarSeason *SolarSeason) Next(seasons int) *SolarSeason {
	m := NewSolarMonthFromYm(solarSeason.year, solarSeason.month).Next(MONTH_IN_SEASON * seasons)
	return NewSolarSeasonFromYm(m.GetYear(), m.GetMonth())
}
