package calendar

import (
	"container/list"
	"github.com/6tail/lunar-go/LunarUtil"
	"github.com/6tail/lunar-go/SolarUtil"
	"math"
	"strconv"
	"strings"
	"time"
)

const J2000 = 2451545

type Solar struct {
	year     int
	month    int
	day      int
	hour     int
	minute   int
	second   int
	calendar time.Time
}

func NewSolar(year int, month int, day int, hour int, minute int, second int) *Solar {
	solar := new(Solar)
	solar.year = year
	solar.month = month
	solar.day = day
	solar.hour = hour
	solar.minute = minute
	solar.second = second
	solar.calendar = time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
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
	return NewSolar(year, month, day, hour, minute, second)
}

func ListSolarFromBaZi(yearGanZhi string, monthGanZhi string, dayGanZhi string, timeGanZhi string) *list.List {
	return ListSolarFromBaZiBySect(yearGanZhi, monthGanZhi, dayGanZhi, timeGanZhi, 2)
}

func ListSolarFromBaZiBySect(yearGanZhi string, monthGanZhi string, dayGanZhi string, timeGanZhi string, sect int) *list.List {
	if sect != 1 {
		sect = 2
	}
	l := list.New()
	today := NewSolarFromDate(time.Now())
	lunar := today.GetLunar()
	offsetYear := LunarUtil.GetJiaZiIndex(lunar.GetYearInGanZhiExact()) - LunarUtil.GetJiaZiIndex(yearGanZhi)
	if offsetYear < 0 {
		offsetYear = offsetYear + 60
	}
	startYear := today.GetYear() - offsetYear
	hour := 0
	gz := []rune(timeGanZhi)
	timeZhi := string(gz[1:])
	j := len(LunarUtil.ZHI)
	for i := 0; i < j; i++ {
		if strings.Compare(LunarUtil.ZHI[i], timeZhi) == 0 {
			hour = (i - 1) * 2
		}
	}
	for {
		if startYear < SolarUtil.BASE_YEAR-1 {
			break
		}
		year := startYear - 1
		counter := 0
		month := 12
		found := false
		for {
			if counter >= 15 {
				break
			}
			if year >= SolarUtil.BASE_YEAR {
				day := 1
				if year == SolarUtil.BASE_YEAR && month == SolarUtil.BASE_MONTH {
					day = SolarUtil.BASE_DAY
				}
				solar := NewSolar(year, month, day, hour, 0, 0)
				lunar = solar.GetLunar()
				if strings.Compare(lunar.GetYearInGanZhiExact(), yearGanZhi) == 0 && strings.Compare(lunar.GetMonthInGanZhiExact(), monthGanZhi) == 0 {
					found = true
					break
				}
			}
			month++
			if month > 12 {
				month = 1
				year++
			}
			counter++
		}
		if found {
			counter = 0
			month--
			if month < 1 {
				month = 12
				year--
			}
			day := 1
			if year == SolarUtil.BASE_YEAR && month == SolarUtil.BASE_MONTH {
				day = SolarUtil.BASE_DAY
			}
			solar := NewSolar(year, month, day, hour, 0, 0)
			for {
				if counter >= 61 {
					break
				}
				lunar = solar.GetLunar()
				dgz := lunar.GetDayInGanZhiExact2()
				if sect == 1 {
					dgz = lunar.GetDayInGanZhiExact()
				}
				if strings.Compare(lunar.GetYearInGanZhiExact(), yearGanZhi) == 0 && strings.Compare(lunar.GetMonthInGanZhiExact(), monthGanZhi) == 0 && strings.Compare(dgz, dayGanZhi) == 0 && strings.Compare(lunar.GetTimeInGanZhi(), timeGanZhi) == 0 {
					l.PushBack(solar)
					break
				}
				solar = solar.Next(1)
				counter++
			}
		}
		startYear -= 60
	}
	return l
}

func padding(n int) string {
	s := ""
	if n < 10 {
		s = "0"
	}
	return s + strconv.Itoa(n)
}

func (solar *Solar) IsLeapYear() bool {
	return SolarUtil.IsLeapYear(solar.year)
}

func (solar *Solar) GetWeek() int {
	return int(solar.calendar.Weekday())
}

func (solar *Solar) GetWeekInChinese() string {
	return SolarUtil.WEEK[solar.GetWeek()]
}

// @Deprecated: 该方法已废弃，请使用GetXingZuo
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
	} else if y >= 521 && y <= 620 {
		index = 2
	} else if y >= 621 && y <= 722 {
		index = 3
	} else if y >= 723 && y <= 822 {
		index = 4
	} else if y >= 823 && y <= 922 {
		index = 5
	} else if y >= 923 && y <= 1022 {
		index = 6
	} else if y >= 1023 && y <= 1121 {
		index = 7
	} else if y >= 1122 && y <= 1221 {
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
	if f, ok := SolarUtil.FESTIVAL[strconv.Itoa(solar.month)+"-"+strconv.Itoa(solar.day)]; ok {
		l.PushBack(f)
	}
	//计算几月第几个星期几对应的节日
	weeks := int(math.Ceil(float64(solar.day) / 7.0))
	week := solar.GetWeek()
	if f, ok := SolarUtil.WEEK_FESTIVAL[strconv.Itoa(solar.month)+"-"+strconv.Itoa(weeks)+"-"+strconv.Itoa(week)]; ok {
		l.PushBack(f)
	}
	return l
}

func (solar *Solar) GetOtherFestivals() *list.List {
	l := list.New()
	if f, ok := SolarUtil.OTHER_FESTIVAL[strconv.Itoa(solar.month)+"-"+strconv.Itoa(solar.day)]; ok {
		for i := 0; i < len(f); i++ {
			l.PushBack(f[i])
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

func (solar *Solar) GetCalendar() time.Time {
	return solar.calendar
}

func (solar *Solar) GetJulianDay() float64 {
	y := solar.year
	m := solar.month
	d := float64(solar.day) + ((float64(solar.second)/60+float64(solar.minute))/60+float64(solar.hour))/24
	n := 0
	g := false
	if y*372+m*31+int(d) >= 588829 {
		g = true
	}
	if m <= 2 {
		m += 12
		y--
	}
	if g {
		n = y / 100
		n = 2 - n + n/4
	}
	return float64(int(365.25*(float64(y)+4716))) + float64(int(30.6001*(float64(m)+1))) + d + float64(n) - 1524.5
}

func (solar *Solar) ToYmd() string {
	return strconv.Itoa(solar.year) + "-" + padding(solar.month) + "-" + padding(solar.day)
}

func (solar *Solar) ToYmdHms() string {
	return solar.ToYmd() + " " + padding(solar.hour) + ":" + padding(solar.minute) + ":" + padding(solar.second)
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

func (solar *Solar) Next(days int) *Solar {
	c := solar.calendar.AddDate(0, 0, days)
	return NewSolar(c.Year(), int(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second())
}

func (solar *Solar) GetLunar() *Lunar {
	return NewLunarFromDate(solar.calendar)
}
