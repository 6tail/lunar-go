package calendar

import (
	"github.com/6tail/lunar-go/LunarUtil"
)

type Yun struct {
	// 性别(1男，0女)
	gender int
	// 起运年数
	startYear int
	// 起运月数
	startMonth int
	// 起运天数
	startDay int
	// 起运小时数
	startHour int
	// 是否顺推
	forward bool
	lunar   *Lunar
}

func NewYun(eightChar *EightChar, gender int, sect int) *Yun {
	yun := new(Yun)
	yun.lunar = eightChar.GetLunar()
	yun.gender = gender
	yang := 0 == yun.lunar.GetYearGanIndexExact()%2
	man := 1 == yun.gender
	yun.forward = (yang && man) || (!yang && !man)
	yun.computeStart(sect)
	return yun
}

// 起运计算
func (yun *Yun) computeStart(sect int) {
	prev := yun.lunar.GetPrevJie()
	next := yun.lunar.GetNextJie()
	current := yun.lunar.GetSolar()
	start := current
	end := current
	if !yun.forward {
		start = prev.GetSolar()
	}
	if yun.forward {
		end = next.GetSolar()
	}
	year := 0
	month := 0
	day := 0
	hour := 0
	if 2 == sect {
		minutes := end.SubtractMinute(start)
		year = minutes / 4320
		minutes -= year * 4320
		month = minutes / 360
		minutes -= month * 360
		day = minutes / 12
		minutes -= day * 12
		hour = minutes * 2
	} else {
		endTimeZhiIndex := 11
		if end.GetHour() != 23 {
			endTimeZhiIndex = LunarUtil.GetTimeZhiIndex(end.ToYmdHms()[11:16])
		}
		startTimeZhiIndex := 11
		if start.GetHour() != 23 {
			startTimeZhiIndex = LunarUtil.GetTimeZhiIndex(start.ToYmdHms()[11:16])
		}
		// 时辰差
		hourDiff := endTimeZhiIndex - startTimeZhiIndex
		dayDiff := end.Subtract(start)
		if hourDiff < 0 {
			hourDiff += 12
			dayDiff--
		}
		monthDiff := hourDiff * 10 / 30
		month = dayDiff*4 + monthDiff
		day = hourDiff*10 - monthDiff*30
		year = month / 12
		month = month - year*12
	}
	yun.startYear = year
	yun.startMonth = month
	yun.startDay = day
	yun.startHour = hour
}

// GetGender 获取性别
func (yun *Yun) GetGender() int {
	return yun.gender
}

// GetStartYear 获取起运年数
func (yun *Yun) GetStartYear() int {
	return yun.startYear
}

// GetStartMonth 获取起运月数
func (yun *Yun) GetStartMonth() int {
	return yun.startMonth
}

// GetStartDay 获取起运天数
func (yun *Yun) GetStartDay() int {
	return yun.startDay
}

// GetStartHour 获取起运小时数
func (yun *Yun) GetStartHour() int {
	return yun.startHour
}

// IsForward 是否顺推
func (yun *Yun) IsForward() bool {
	return yun.forward
}

func (yun *Yun) GetLunar() *Lunar {
	return yun.lunar
}

// GetStartSolar 获取起运的阳历日期
func (yun *Yun) GetStartSolar() *Solar {
	solar := yun.lunar.GetSolar()
	solar = solar.NextYear(yun.startYear)
	solar = solar.NextMonth(yun.startMonth)
	solar = solar.NextDay(yun.startDay)
	return solar.NextHour(yun.startHour)
}

// GetDaYun 获取10轮大运
func (yun *Yun) GetDaYun() []*DaYun {
	return yun.GetDaYunBy(10)
}

// GetDaYunBy 获取大运
func (yun *Yun) GetDaYunBy(n int) []*DaYun {
	l := make([]*DaYun, n)
	for i := 0; i < n; i++ {
		l[i] = NewDaYun(yun, i)
	}
	return l
}
