package calendar

import (
	"container/list"
	"fmt"
	"github.com/6tail/lunar-go/SolarUtil"
	"math"
	"time"
)

// SolarWeek 阳历周
type SolarWeek struct {
	year  int
	month int
	day   int
	start int
}

func NewSolarWeek(start int) *SolarWeek {
	return NewSolarWeekFromDate(time.Now().Local(), start)
}

func NewSolarWeekFromYmd(year int, month int, day int, start int) *SolarWeek {
	solarWeek := new(SolarWeek)
	solarWeek.year = year
	solarWeek.month = month
	solarWeek.day = day
	solarWeek.start = start
	return solarWeek
}

func NewSolarWeekFromDate(date time.Time, start int) *SolarWeek {
	return NewSolarWeekFromYmd(date.Year(), int(date.Month()), date.Day(), start)
}

func (solarWeek *SolarWeek) GetYear() int {
	return solarWeek.year
}

func (solarWeek *SolarWeek) GetMonth() int {
	return solarWeek.month
}

func (solarWeek *SolarWeek) GetDay() int {
	return solarWeek.day
}

func (solarWeek *SolarWeek) GetIndex() int {
	offset := NewSolarFromYmd(solarWeek.year, solarWeek.month, 1).GetWeek() - solarWeek.start
	if offset < 0 {
		offset += 7
	}
	return int(math.Ceil(float64(solarWeek.day+offset) / 7))
}

func (solarWeek *SolarWeek) GetIndexInYear() int {
	offset := NewSolarFromYmd(solarWeek.year, 1, 1).GetWeek() - solarWeek.start
	if offset < 0 {
		offset += 7
	}
	return int(math.Ceil(float64(SolarUtil.GetDaysInYear(solarWeek.year, solarWeek.month, solarWeek.day)+offset) / 7))
}

func (solarWeek *SolarWeek) GetFirstDay() *Solar {
	c := NewSolarFromYmd(solarWeek.year, solarWeek.month, solarWeek.day)
	prev := c.GetWeek() - solarWeek.start
	if prev < 0 {
		prev += 7
	}
	return c.NextDay(-prev)
}

func (solarWeek *SolarWeek) GetFirstDayInMonth() *Solar {
	days := solarWeek.GetDays()
	for i := days.Front(); i != nil; i = i.Next() {
		day := i.Value.(*Solar)
		if solarWeek.month == day.month {
			return day
		}
	}
	return nil
}

func (solarWeek *SolarWeek) GetDays() *list.List {
	firstDay := solarWeek.GetFirstDay()
	l := list.New()
	l.PushBack(firstDay)
	for i := 1; i < 7; i++ {
		l.PushBack(firstDay.NextDay(i))
	}
	return l
}

func (solarWeek *SolarWeek) GetDaysInMonth() *list.List {
	days := solarWeek.GetDays()
	l := list.New()
	for i := days.Front(); i != nil; i = i.Next() {
		day := i.Value.(Solar)
		if solarWeek.month == day.month {
			l.PushBack(day)
		}
	}
	return l
}

func (solarWeek *SolarWeek) String() string {
	return fmt.Sprintf("%d.%d.%d", solarWeek.year, solarWeek.month, solarWeek.GetIndex())
}

func (solarWeek *SolarWeek) ToFullString() string {
	return fmt.Sprintf("%d年%d月第%d周", solarWeek.year, solarWeek.month, solarWeek.GetIndex())
}

func (solarWeek *SolarWeek) Next(weeks int, separateMonth bool) *SolarWeek {
	if 0 == weeks {
		return NewSolarWeekFromYmd(solarWeek.year, solarWeek.month, solarWeek.day, solarWeek.start)
	}
	c := NewSolarFromYmd(solarWeek.year, solarWeek.month, solarWeek.day)
	if separateMonth {
		n := weeks
		week := NewSolarWeekFromYmd(c.GetYear(), c.GetMonth(), c.GetDay(), solarWeek.start)
		month := solarWeek.month
		plus := false
		if n > 0 {
			plus = true
		}
		for 0 != n {
			if plus {
				c = c.NextDay(7)
			} else {
				c = c.NextDay(-7)
			}
			week := NewSolarWeekFromYmd(c.GetYear(), c.GetMonth(), c.GetDay(), solarWeek.start)
			weekMonth := week.GetMonth()
			if month != weekMonth {
				index := week.GetIndex()
				if plus {
					if 1 == index {
						firstDay := week.GetFirstDay()
						week = NewSolarWeekFromYmd(firstDay.GetYear(), firstDay.GetMonth(), firstDay.GetDay(), solarWeek.start)
						weekMonth = week.month
					} else {
						c = NewSolarFromYmd(week.year, week.month, 1)
						week = NewSolarWeekFromYmd(c.GetYear(), c.GetMonth(), c.GetDay(), solarWeek.start)
					}
				} else {
					if SolarUtil.GetWeeksOfMonth(week.year, week.month, solarWeek.start) == index {
						lastDay := week.GetFirstDay().NextDay(6)
						week = NewSolarWeekFromYmd(lastDay.year, lastDay.month, lastDay.day, solarWeek.start)
						weekMonth = week.month
					} else {
						c = NewSolarFromYmd(week.GetYear(), week.GetMonth(), SolarUtil.GetDaysOfMonth(week.year, week.month))
						week = NewSolarWeekFromYmd(c.GetYear(), c.GetMonth(), c.GetDay(), solarWeek.start)
					}
				}
				month = weekMonth
			}
			if plus {
				n -= 1
			} else {
				n -= -1
			}
		}
		return week
	} else {
		c = c.NextDay(weeks * 7)
		return NewSolarWeekFromYmd(c.GetYear(), c.GetMonth(), c.GetDay(), solarWeek.start)
	}
}
