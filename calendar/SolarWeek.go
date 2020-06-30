package calendar

import (
	"container/list"
	"github.com/6tail/lunar-go/SolarUtil"
	"math"
	"strconv"
	"time"
)

type SolarWeek struct {
	year  int
	month int
	day   int
	start int
}

func NewSolarWeek(start int) *SolarWeek {
	return NewSolarWeekFromDate(time.Now(), start)
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
	c := time.Date(solarWeek.year, time.Month(solarWeek.month), 1, 0, 0, 0, 0, time.Local)
	week := int(c.Weekday())
	if week == 0 {
		week = 7
	}
	return int(math.Ceil(float64((solarWeek.day + week - solarWeek.start) / 7)))
}

func (solarWeek *SolarWeek) GetFirstDay() *Solar {
	c := time.Date(solarWeek.year, time.Month(solarWeek.month), solarWeek.day, 0, 0, 0, 0, time.Local)
	week := int(c.Weekday())
	prev := week - solarWeek.start
	if prev < 0 {
		prev += 7
	}
	c.AddDate(0, 0, -prev)
	return NewSolarFromDate(c)
}

func (solarWeek *SolarWeek) GetFirstDayInMonth() *Solar {
	days := solarWeek.GetDays()
	for i := days.Front(); i != nil; i = i.Next() {
		day := i.Value.(Solar)
		if solarWeek.month == day.month {
			return NewSolarFromDate(day.calendar)
		}
	}
	return nil
}

func (solarWeek *SolarWeek) GetDays() *list.List {
	firstDay := solarWeek.GetFirstDay()
	l := list.New()
	l.PushBack(firstDay)
	for i := 1; i < 7; i++ {
		l.PushBack(firstDay.Next(i))
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
	return strconv.Itoa(solarWeek.year) + "." + strconv.Itoa(solarWeek.month) + "." + strconv.Itoa(solarWeek.GetIndex())
}

func (solarWeek *SolarWeek) ToFullString() string {
	return strconv.Itoa(solarWeek.year) + "年" + strconv.Itoa(solarWeek.month) + "月第" + strconv.Itoa(solarWeek.GetIndex()) + "周"
}

func (solarWeek *SolarWeek) Next(weeks int, separateMonth bool) *SolarWeek {
	if 0 == weeks {
		return NewSolarWeekFromYmd(solarWeek.year, solarWeek.month, solarWeek.day, solarWeek.start)
	}
	if separateMonth {
		n := weeks
		c := time.Date(solarWeek.year, time.Month(solarWeek.month), solarWeek.day, 0, 0, 0, 0, time.Local)
		week := NewSolarWeekFromDate(c, solarWeek.start)
		month := solarWeek.month
		plus := false
		if n > 0 {
			plus = true
		}
		for {
			if n == 0 {
				break
			}
			if plus {
				c = c.AddDate(0, 0, 7)
			} else {
				c = c.AddDate(0, 0, -7)
			}
			week = NewSolarWeekFromDate(c, solarWeek.start)
			weekMonth := week.month
			if month != weekMonth {
				index := week.GetIndex()
				if plus {
					if 1 == index {
						firstDay := week.GetFirstDay()
						week = NewSolarWeekFromYmd(firstDay.year, firstDay.month, firstDay.day, solarWeek.start)
						weekMonth = week.month
					} else {
						c = time.Date(week.year, time.Month(week.month), 1, 0, 0, 0, 0, time.Local)
						week = NewSolarWeekFromDate(c, solarWeek.start)
					}
				} else {
					size := SolarUtil.GetWeeksOfMonth(week.year, week.month, solarWeek.start)
					if size == index {
						firstDay := week.GetFirstDay()
						lastDay := firstDay.Next(6)
						week = NewSolarWeekFromYmd(lastDay.year, lastDay.month, lastDay.day, solarWeek.start)
						weekMonth = week.month
					} else {
						c = time.Date(week.year, time.Month(week.month), SolarUtil.GetDaysOfMonth(week.year, week.month), 0, 0, 0, 0, time.Local)
						week = NewSolarWeekFromDate(c, solarWeek.start)
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
		c := time.Date(solarWeek.year, time.Month(solarWeek.month), solarWeek.day, 0, 0, 0, 0, time.Local)
		c.AddDate(0, 0, weeks*7)
		return NewSolarWeekFromDate(c, solarWeek.start)
	}
}
