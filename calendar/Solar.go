package calendar

import (
	"container/list"
	"github.com/6tail/lunar-go/SolarUtil"
	"math"
	"strconv"
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
	julianDay += 0.5

	// 日数的整数部份
	a := int2(julianDay)
	// 日数的小数部分
	f := julianDay - a
	jd := float64(0)

	if a > 2299161 {
		jd = int2((a - 1867216.25) / 36524.25)
		a += 1 + jd - int2(jd/4)
	}
	// 向前移4年零2个月
	a += 1524
	y := int(int2((a - 122.1) / 365.25))
	// 去除整年日数后余下日数
	jd = a - int2(365.25*float64(y))
	m := int(int2(jd / 30.6001))
	// 去除整月日数后余下日数
	d := int(int2(jd - int2(float64(m)*30.6001)))
	y -= 4716
	m--
	if m > 12 {
		m -= 12
	}
	if m <= 2 {
		y++
	}

	// 日的小数转为时分秒
	f *= 24
	h := int(int2(f))

	f -= float64(h)
	f *= 60
	mi := int(int2(f))

	f -= float64(mi)
	f *= 60
	s := int(int2(f))
	return NewSolar(y, m, d, h, mi, s)
}

func int2(v float64) float64 {
	v = math.Floor(v)
	if v < 0 {
		return v + 1
	} else {
		return v
	}
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
	//第几周
	week := solar.GetWeek()
	weekInMonth := int(math.Ceil(float64((solar.day - week) / 7)))
	if week > 0 {
		weekInMonth++
	}
	if f, ok := SolarUtil.WEEK_FESTIVAL[strconv.Itoa(solar.month)+"-"+strconv.Itoa(weekInMonth)+"-"+strconv.Itoa(week)]; ok {
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
	y := float64(solar.year)
	m := float64(solar.month)
	n := float64(0)

	if m <= 2 {
		m += 12
		y--
	}

	// 判断是否为UTC日1582*372+10*31+15
	if solar.year*372+solar.month*31+solar.day >= 588829 {
		n = int2(y / 100)
		// 加百年闰
		n = 2 - n + int2(n/4)
	}

	// 加上年引起的偏移日数
	n += int2(365.2500001 * (y + 4716))
	// 加上月引起的偏移日数及日偏移数
	n += int2(30.6*(m+1)) + float64(solar.day)
	n += ((float64(solar.second)/60+float64(solar.minute))/60+float64(solar.hour))/24 - 1524.5
	return n
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
