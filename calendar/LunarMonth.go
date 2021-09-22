package calendar

import (
	"fmt"
	"github.com/6tail/lunar-go/LunarUtil"
)

// 阴历月
type LunarMonth struct {
	year           int
	month          int
	dayCount       int
	firstJulianDay float64
}

func NewLunarMonth(lunarYear int, lunarMonth int, dayCount int, firstJulianDay float64) *LunarMonth {
	month := new(LunarMonth)
	month.year = lunarYear
	month.month = lunarMonth
	month.dayCount = dayCount
	month.firstJulianDay = firstJulianDay
	return month
}

func NewLunarMonthFromYm(lunarYear int, lunarMonth int) *LunarMonth {
	year := new(LunarYear)
	year.year = lunarYear
	return NewLunarYear(lunarYear).GetMonth(lunarMonth)
}

func (lunarMonth *LunarMonth) GetYear() int {
	return lunarMonth.year
}

func (lunarMonth *LunarMonth) GetMonth() int {
	return lunarMonth.month
}

func (lunarMonth *LunarMonth) IsLeap() bool {
	return lunarMonth.month < 0
}

func (lunarMonth *LunarMonth) GetDayCount() int {
	return lunarMonth.dayCount
}

func (lunarMonth *LunarMonth) GetFirstJulianDay() float64 {
	return lunarMonth.firstJulianDay
}

func (lunarMonth *LunarMonth) String() string {
	run := ""
	if lunarMonth.IsLeap() {
		run = "闰"
	}
	m := lunarMonth.month
	if m < 0 {
		m = -m
	}
	return fmt.Sprintf("%d年%s%s月(%d)天", lunarMonth.year, run, LunarUtil.MONTH[m], lunarMonth.dayCount)
}
