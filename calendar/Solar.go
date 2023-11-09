package calendar

import (
	"container/list"
	"fmt"
	"github.com/6tail/lunar-go/HolidayUtil"
	"github.com/6tail/lunar-go/LunarUtil"
	"github.com/6tail/lunar-go/SolarUtil"
	"math"
	"strings"
	"time"
)

const J2000 = 2451545

// Solar 阳历
type Solar struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
	second int
}

func NewSolar(year int, month int, day int, hour int, minute int, second int) *Solar {
	if 1582 == year && 10 == month {
		if day > 4 && day < 15 {
			panic(fmt.Sprintf("wrong solar year %v month %v day %v", year, month, day))
		}
	}
	if month < 1 || month > 12 {
		panic(fmt.Sprintf("wrong month %v", month))
	}
	if day < 1 || day > 31 {
		panic(fmt.Sprintf("wrong day %v", day))
	}
	if hour < 0 || hour > 23 {
		panic(fmt.Sprintf("wrong hour %v", hour))
	}
	if minute < 0 || minute > 59 {
		panic(fmt.Sprintf("wrong minute %v", minute))
	}
	if second < 0 || second > 59 {
		panic(fmt.Sprintf("wrong second %v", second))
	}
	solar := new(Solar)
	solar.year = year
	solar.month = month
	solar.day = day
	solar.hour = hour
	solar.minute = minute
	solar.second = second
	return solar
}

func NewSolarFromYmd(year int, month int, day int) *Solar {
	return NewSolar(year, month, day, 0, 0, 0)
}

func NewSolarFromDate(date time.Time) *Solar {
	return NewSolar(date.Year(), int(date.Month()), date.Day(), date.Hour(), date.Minute(), date.Second())
}

func NewSolarFromJulianDay(julianDay float64) *Solar {
	d := int(julianDay + 0.5)
	f := julianDay + 0.5 - float64(d)

	if d >= 2299161 {
		c := int((float64(d) - 1867216.25) / 36524.25)
		d += 1 + c - c/4
	}
	d += 1524
	year := int((float64(d) - 122.1) / 365.25)
	d -= int(365.25 * float64(year))
	month := int(float64(d) / 30.601)
	d -= int(30.601 * float64(month))
	day := d
	if month > 13 {
		month -= 13
		year -= 4715
	} else {
		month -= 1
		year -= 4716
	}
	f *= 24
	hour := int(f)

	f -= float64(hour)
	f *= 60
	minute := int(f)

	f -= float64(minute)
	f *= 60
	second := int(math.Round(f))

	if second > 59 {
		second -= 60
		minute++
	}
	if minute > 59 {
		minute -= 60
		hour++
	}
	if hour > 23 {
		hour -= 24
		day += 1
	}

	return NewSolar(year, month, day, hour, minute, second)
}

func ListSolarFromBaZi(yearGanZhi string, monthGanZhi string, dayGanZhi string, timeGanZhi string) *list.List {
	return ListSolarFromBaZiBySect(yearGanZhi, monthGanZhi, dayGanZhi, timeGanZhi, 2)
}

func ListSolarFromBaZiBySect(yearGanZhi string, monthGanZhi string, dayGanZhi string, timeGanZhi string, sect int) *list.List {
	return ListSolarFromBaZiBySectAndBaseYear(yearGanZhi, monthGanZhi, dayGanZhi, timeGanZhi, sect, 1900)
}

func ListSolarFromBaZiBySectAndBaseYear(yearGanZhi string, monthGanZhi string, dayGanZhi string, timeGanZhi string, sect int, baseYear int) *list.List {
	if sect != 1 {
		sect = 2
	}
	l := list.New()
	years := list.New()
	today := NewSolarFromDate(time.Now().Local())
	offsetYear := LunarUtil.GetJiaZiIndex(today.GetLunar().GetYearInGanZhiExact()) - LunarUtil.GetJiaZiIndex(yearGanZhi)
	if offsetYear < 0 {
		offsetYear = offsetYear + 60
	}
	startYear := today.GetYear() - offsetYear - 1
	minYear := baseYear - 2
	for startYear >= minYear {
		years.PushBack(startYear)
		startYear -= 60
	}
	hours := list.New()
	gz := []rune(timeGanZhi)
	timeZhi := string(gz[1:])
	for i, z := range LunarUtil.ZHI {
		if strings.Compare(z, timeZhi) == 0 {
			hours.PushBack((i - 1) * 2)
			break
		}
	}
	if strings.Compare("子", timeZhi) == 0 {
		hours.PushBack(23)
	}
	for m := hours.Front(); m != nil; m = m.Next() {
		hour := m.Value.(int)
		for i := years.Front(); i != nil; i = i.Next() {
			y := i.Value.(int)
			maxYear := y + 3
			year := y
			month := 11
			if year < baseYear {
				year = baseYear
				month = 1
			}
			o := NewSolar(year, month, 1, hour, 0, 0)
			for o.GetYear() <= maxYear {
				lunar := o.GetLunar()
				dgz := lunar.GetDayInGanZhiExact2()
				if sect == 1 {
					dgz = lunar.GetDayInGanZhiExact()
				}
				if strings.Compare(lunar.GetYearInGanZhiExact(), yearGanZhi) == 0 && strings.Compare(lunar.GetMonthInGanZhiExact(), monthGanZhi) == 0 && strings.Compare(dgz, dayGanZhi) == 0 && strings.Compare(lunar.GetTimeInGanZhi(), timeGanZhi) == 0 {
					l.PushBack(o)
					break
				}
				o = o.NextDay(1)
			}
		}
	}
	return l
}

func (solar *Solar) IsLeapYear() bool {
	return SolarUtil.IsLeapYear(solar.year)
}

func (solar *Solar) GetWeek() int {
	return SolarUtil.GetWeek(solar.year, solar.month, solar.day)
}

func (solar *Solar) GetWeekInChinese() string {
	return SolarUtil.WEEK[solar.GetWeek()]
}

// GetXingzuo @Deprecated: 该方法已废弃，请使用GetXingZuo
func (solar *Solar) GetXingzuo() string {
	return solar.GetXingZuo()
}

func (solar *Solar) GetXingZuo() string {
	index := 11
	m := solar.month
	d := solar.day
	y := m*100 + d
	if y >= 321 && y <= 419 {
		index = 0
	} else if y >= 420 && y <= 520 {
		index = 1
	} else if y >= 521 && y <= 621 {
		index = 2
	} else if y >= 622 && y <= 722 {
		index = 3
	} else if y >= 723 && y <= 822 {
		index = 4
	} else if y >= 823 && y <= 922 {
		index = 5
	} else if y >= 923 && y <= 1023 {
		index = 6
	} else if y >= 1024 && y <= 1122 {
		index = 7
	} else if y >= 1123 && y <= 1221 {
		index = 8
	} else if y >= 1222 || y <= 119 {
		index = 9
	} else if y <= 218 {
		index = 10
	}
	return SolarUtil.XINGZUO[index]
}

func (solar *Solar) GetFestivals() *list.List {
	l := list.New()
	//获取几月几日对应的节日
	if f, ok := SolarUtil.FESTIVAL[fmt.Sprintf("%d-%d", solar.month, solar.day)]; ok {
		l.PushBack(f)
	}
	//计算几月第几个星期几对应的节日
	weeks := int(math.Ceil(float64(solar.day) / 7))
	week := solar.GetWeek()
	if f, ok := SolarUtil.WEEK_FESTIVAL[fmt.Sprintf("%d-%d-%d", solar.month, weeks, week)]; ok {
		l.PushBack(f)
	}
	if solar.day+7 > SolarUtil.GetDaysOfMonth(solar.year, solar.month) {
		if f, ok := SolarUtil.WEEK_FESTIVAL[fmt.Sprintf("%d-0-%d", solar.month, week)]; ok {
			l.PushBack(f)
		}
	}
	return l
}

func (solar *Solar) GetOtherFestivals() *list.List {
	l := list.New()
	if f, ok := SolarUtil.OTHER_FESTIVAL[fmt.Sprintf("%d-%d", solar.month, solar.day)]; ok {
		for _, v := range f {
			l.PushBack(v)
		}
	}
	return l
}

func (solar *Solar) GetYear() int {
	return solar.year
}

func (solar *Solar) GetMonth() int {
	return solar.month
}

func (solar *Solar) GetDay() int {
	return solar.day
}

func (solar *Solar) GetHour() int {
	return solar.hour
}

func (solar *Solar) GetMinute() int {
	return solar.minute
}

func (solar *Solar) GetSecond() int {
	return solar.second
}

func (solar *Solar) GetJulianDay() float64 {
	return SolarUtil.GetJulianDay(solar.year, solar.month, solar.day, solar.hour, solar.minute, solar.second)
}

func (solar *Solar) ToYmd() string {
	return fmt.Sprintf("%04d-%02d-%02d", solar.year, solar.month, solar.day)
}

func (solar *Solar) ToYmdHms() string {
	return fmt.Sprintf("%v %02d:%02d:%02d", solar.ToYmd(), solar.hour, solar.minute, solar.second)
}

func (solar *Solar) String() string {
	return solar.ToYmd()
}

func (solar *Solar) ToFullString() string {
	s := solar.ToYmdHms()
	if solar.IsLeapYear() {
		s += " 闰年"
	}
	s += " 星期"
	s += solar.GetWeekInChinese()
	for i := solar.GetFestivals().Front(); i != nil; i = i.Next() {
		s += " ("
		s += i.Value.(string)
		s += ")"
	}
	for i := solar.GetOtherFestivals().Front(); i != nil; i = i.Next() {
		s += " ("
		s += i.Value.(string)
		s += ")"
	}
	s += " "
	s += solar.GetXingZuo()
	s += "座"
	return s
}

func (solar *Solar) Subtract(other *Solar) int {
	return SolarUtil.GetDaysBetween(other.GetYear(), other.GetMonth(), other.GetDay(), solar.GetYear(), solar.GetMonth(), solar.GetDay())
}

func (solar *Solar) SubtractMinute(other *Solar) int {
	days := solar.Subtract(other)
	cm := solar.GetHour()*60 + solar.GetMinute()
	sm := other.GetHour()*60 + other.GetMinute()
	m := cm - sm
	if m < 0 {
		m += 1440
		days--
	}
	m += days * 1440
	return m
}

func (solar *Solar) IsAfter(other *Solar) bool {
	if solar.GetYear() > other.GetYear() {
		return true
	}
	if solar.GetYear() < other.GetYear() {
		return false
	}
	if solar.GetMonth() > other.GetMonth() {
		return true
	}
	if solar.GetMonth() < other.GetMonth() {
		return false
	}
	if solar.GetDay() > other.GetDay() {
		return true
	}
	if solar.GetDay() < other.GetDay() {
		return false
	}
	if solar.GetHour() > other.GetHour() {
		return true
	}
	if solar.GetHour() < other.GetHour() {
		return false
	}
	if solar.GetMinute() > other.GetMinute() {
		return true
	}
	if solar.GetMinute() < other.GetMinute() {
		return false
	}
	return solar.GetSecond() > other.GetSecond()
}

func (solar *Solar) IsBefore(other *Solar) bool {
	return SolarUtil.IsBefore(solar.year, solar.month, solar.day, solar.hour, solar.minute, solar.second, other.GetYear(), other.GetMonth(), other.GetDay(), other.GetHour(), other.GetMinute(), other.GetSecond())
}

func (solar *Solar) NextYear(years int) *Solar {
	y := solar.GetYear() + years
	m := solar.GetMonth()
	d := solar.GetDay()
	if 1582 == y && 10 == m {
		if d > 4 && d < 15 {
			d += 10
		}
	} else if 2 == m {
		if d > 28 {
			if !SolarUtil.IsLeapYear(y) {
				d = 28
			}
		}
	}
	return NewSolar(y, m, d, solar.GetHour(), solar.GetMinute(), solar.GetSecond())
}

func (solar *Solar) NextMonth(months int) *Solar {
	month := NewSolarMonthFromYm(solar.GetYear(), solar.GetMonth()).Next(months)
	y := month.GetYear()
	m := month.GetMonth()
	d := solar.GetDay()
	if 1582 == y && 10 == m {
		if d > 4 && d < 15 {
			d += 10
		}
	} else {
		maxDay := SolarUtil.GetDaysOfMonth(y, m)
		if d > maxDay {
			d = maxDay
		}
	}
	return NewSolar(y, m, d, solar.GetHour(), solar.GetMinute(), solar.GetSecond())
}

func (solar *Solar) NextDay(days int) *Solar {
	y := solar.GetYear()
	m := solar.GetMonth()
	d := solar.GetDay()
	if 1582 == y && 10 == m {
		if d > 4 {
			d -= 10
		}
	}
	if days > 0 {
		d += days
		daysInMonth := SolarUtil.GetDaysOfMonth(y, m)
		for d > daysInMonth {
			d -= daysInMonth
			m++
			if m > 12 {
				m = 1
				y++
			}
			daysInMonth = SolarUtil.GetDaysOfMonth(y, m)
		}
	} else if days < 0 {
		for d+days <= 0 {
			m--
			if m < 1 {
				m = 12
				y--
			}
			d += SolarUtil.GetDaysOfMonth(y, m)
		}
		d += days
	}
	if 1582 == y && 10 == m {
		if d > 4 {
			d += 10
		}
	}
	return NewSolar(y, m, d, solar.GetHour(), solar.GetMinute(), solar.GetSecond())
}

func (solar *Solar) Next(days int, onlyWorkday bool) *Solar {
	if !onlyWorkday {
		return solar.NextDay(days)
	}
	o := NewSolar(solar.GetYear(), solar.GetMonth(), solar.GetDay(), solar.GetHour(), solar.GetMinute(), solar.GetSecond())
	if days != 0 {
		rest := days
		add := 1
		if days < 0 {
			rest = -days
			add = -1
		}
		for rest > 0 {
			o = o.NextDay(add)
			work := true
			holiday := HolidayUtil.GetHolidayByYmd(o.GetYear(), o.GetMonth(), o.GetDay())
			if nil == holiday {
				week := o.GetWeek()
				if 0 == week || 6 == week {
					work = false
				}
			} else {
				work = holiday.IsWork()
			}
			if work {
				rest -= 1
			}
		}
	}
	return o
}

func (solar *Solar) NextHour(hours int) *Solar {
	h := solar.GetHour() + hours
	n := 1
	hour := h
	if h < 0 {
		n = -1
		hour = -h
	}
	days := hour / 24 * n
	hour = (hour % 24) * n
	if hour < 0 {
		hour += 24
		days--
	}
	o := solar.NextDay(days)
	return NewSolar(o.GetYear(), o.GetMonth(), o.GetDay(), hour, o.GetMinute(), o.GetSecond())
}

func (solar *Solar) GetLunar() *Lunar {
	return NewLunarFromSolar(solar)
}

func (solar *Solar) GetSalaryRate() int {
	// 元旦节
	if solar.month == 1 && solar.day == 1 {
		return 3
	}
	// 劳动节
	if solar.month == 5 && solar.day == 1 {
		return 3
	}
	// 国庆
	if solar.month == 10 && solar.day >= 1 && solar.day <= 3 {
		return 3
	}
	lunar := solar.GetLunar()
	// 春节
	if lunar.GetMonth() == 1 && lunar.GetDay() >= 1 && lunar.GetDay() <= 3 {
		return 3
	}
	// 端午
	if lunar.GetMonth() == 5 && lunar.GetDay() == 5 {
		return 3
	}
	// 中秋
	if lunar.GetMonth() == 8 && lunar.GetDay() == 15 {
		return 3
	}
	// 清明
	if strings.Compare("清明", lunar.GetJieQi()) == 0 {
		return 3
	}
	holiday := HolidayUtil.GetHolidayByYmd(solar.year, solar.month, solar.day)
	if nil != holiday {
		// 法定假日非上班
		if !holiday.IsWork() {
			return 2
		}
	} else {
		// 周末
		week := solar.GetWeek()
		if week == 6 || week == 0 {
			return 2
		}
	}
	// 工作日
	return 1
}
