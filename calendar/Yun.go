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
	// 是否顺推
	forward bool
	lunar   *Lunar
}

func NewYun(eightChar *EightChar, gender int) *Yun {
	yun := new(Yun)
	yun.lunar = eightChar.GetLunar()
	yun.gender = gender
	yang := 0 == yun.lunar.GetYearGanIndexExact()%2
	man := 1 == yun.gender
	yun.forward = (yang && man) || (!yang && !man)
	yun.computeStart()
	return yun
}

// 起运计算
func (yun *Yun) computeStart() {
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
	endCalendar := NewExactDateFromYmd(end.GetYear(), end.GetMonth(), end.GetDay())
	startCalendar := NewExactDateFromYmd(start.GetYear(), start.GetMonth(), start.GetDay())
	dayDiff := (int)((endCalendar.Unix() - startCalendar.Unix()) / 86400)
	if hourDiff < 0 {
		hourDiff += 12
		dayDiff--
	}
	monthDiff := hourDiff * 10 / 30
	month := dayDiff*4 + monthDiff
	day := hourDiff*10 - monthDiff*30
	year := month / 12
	month = month - year*12
	yun.startYear = year
	yun.startMonth = month
	yun.startDay = day
}

// 获取性别
func (yun *Yun) GetGender() int {
	return yun.gender
}

// 获取起运年数
func (yun *Yun) GetStartYear() int {
	return yun.startYear
}

// 获取起运月数
func (yun *Yun) GetStartMonth() int {
	return yun.startMonth
}

// 获取起运天数
func (yun *Yun) GetStartDay() int {
	return yun.startDay
}

// 是否顺推
func (yun *Yun) IsForward() bool {
	return yun.forward
}

func (yun *Yun) GetLunar() *Lunar {
	return yun.lunar
}

// 获取起运的阳历日期
func (yun *Yun) GetStartSolar() *Solar {
	birth := yun.lunar.GetSolar()
	c := NewExactDateFromYmd(birth.GetYear(), birth.GetMonth(), birth.GetDay())
	c = c.AddDate(yun.startYear, yun.startMonth, yun.startDay)
	return NewSolarFromDate(c)
}

// 获取10轮大运
func (yun *Yun) GetDaYun() []*DaYun {
	return yun.GetDaYunBy(10)
}

// 获取大运
func (yun *Yun) GetDaYunBy(n int) []*DaYun {
	l := make([]*DaYun, n)
	for i := 0; i < n; i++ {
		l[i] = NewDaYun(yun, i)
	}
	return l
}
