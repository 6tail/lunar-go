package calendar

import (
	"container/list"
	"github.com/6tail/lunar-go/ShouXingUtil"
	"math"
	"strconv"
)

type LunarYear struct {
	year            int
	months          *list.List
	jieQiJulianDays []float64
}

func NewLunarYear(lunarYear int) *LunarYear {
	year := new(LunarYear)
	year.year = lunarYear
	year.months = list.New()
	year.compute()
	return year
}

func (lunarYear *LunarYear) compute() {
	// 节气(中午12点)，长度25
	jq := make([]float64, 25)
	// 合朔，即每月初一(中午12点)，长度16
	hs := make([]float64, 16)
	// 每月天数，长度15
	dayCounts := make([]int, 15)

	year := lunarYear.year - 2000
	// 从上年的大雪到下年的立春
	j := len(JIE_QI_IN_USE)
	lunarYear.jieQiJulianDays = make([]float64, j)
	i := 0
	for i = 0; i < j; i++ {
		// 精确的节气
		t := 36525 * ShouXingUtil.SaLonT((float64(year)+float64(17+i)*15/360)*ShouXingUtil.PI_2)
		t += ShouXingUtil.ONE_THIRD - ShouXingUtil.DtT(t)
		lunarYear.jieQiJulianDays[i] = t + J2000
		// 按中午12点算的节气
		if i > 0 && i < 26 {
			jq[i-1] = math.Round(t)
		}
	}

	// 冬至前的初一
	w := ShouXingUtil.CalcShuo(jq[0])
	if w > jq[0] {
		w -= 29.5306
	}
	// 递推每月初一
	for i = 0; i < 16; i++ {
		hs[i] = ShouXingUtil.CalcShuo(w + 29.5306*float64(i))
	}
	// 每月天数
	for i = 0; i < 15; i++ {
		dayCounts[i] = (int)(hs[i+1] - hs[i])
	}

	leap := -1
	if hs[13] <= jq[24] {
		i := 1
		for hs[i+1] > jq[2*i] && i < 13 {
			i++
		}
		leap = i
	}

	y := lunarYear.year - 1
	m := 11
	for i = 0; i < 15; i++ {
		isLeap := false
		if i == leap {
			isLeap = true
			m--
		}
		month := m
		if isLeap {
			month = -m
		}
		lunarYear.months.PushBack(NewLunarMonth(y, month, dayCounts[i], hs[i]+J2000))
		m++
		if m == 13 {
			m = 1
			y++
		}
	}
}

func (lunarYear *LunarYear) GetYear() int {
	return lunarYear.year
}

func (lunarYear *LunarYear) GetMonths() *list.List {
	return lunarYear.months
}

func (lunarYear *LunarYear) GetJieQiJulianDays() []float64 {
	return lunarYear.jieQiJulianDays
}

func (lunarYear *LunarYear) GetMonth(lunarMonth int) *LunarMonth {
	for i := lunarYear.months.Front(); i != nil; i = i.Next() {
		m := i.Value.(*LunarMonth)
		if m.GetYear() == lunarYear.year && m.GetMonth() == lunarMonth {
			return m
		}
	}
	return nil
}

func (lunarYear *LunarYear) String() string {
	return strconv.Itoa(lunarYear.year) + ""
}

func (lunarYear *LunarYear) ToFullString() string {
	return strconv.Itoa(lunarYear.year) + "年"
}
