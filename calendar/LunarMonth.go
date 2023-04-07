package calendar

import (
	"fmt"
	"github.com/6tail/lunar-go/LunarUtil"
)

// LunarMonth 阴历月
type LunarMonth struct {
	year           int
	month          int
	dayCount       int
	index          int
	zhiIndex       int
	firstJulianDay float64
}

func NewLunarMonth(lunarYear int, lunarMonth int, dayCount int, firstJulianDay float64, index int) *LunarMonth {
	month := new(LunarMonth)
	month.year = lunarYear
	month.month = lunarMonth
	month.dayCount = dayCount
	month.firstJulianDay = firstJulianDay
	month.index = index
	month.zhiIndex = (index - 1 + LunarUtil.BASE_MONTH_ZHI_INDEX) % 12
	return month
}

func NewLunarMonthFromYm(lunarYear int, lunarMonth int) *LunarMonth {
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

func (lunarMonth *LunarMonth) GetIndex() int {
	return lunarMonth.index
}

func (lunarMonth *LunarMonth) GetZhiIndex() int {
	return lunarMonth.zhiIndex
}

func (lunarMonth *LunarMonth) GetGan() string {
	return LunarUtil.GAN[lunarMonth.GetGanIndex()+1]
}

func (lunarMonth *LunarMonth) GetZhi() string {
	return LunarUtil.ZHI[lunarMonth.zhiIndex+1]
}

func (lunarMonth *LunarMonth) GetGanZhi() string {
	return lunarMonth.GetGan() + lunarMonth.GetZhi()
}

func (lunarMonth *LunarMonth) GetGanIndex() int {
	offset := (NewLunarYear(lunarMonth.year).GetGanIndex() + 1) % 5 * 2
	return (lunarMonth.index - 1 + offset) % 10
}

func (lunarMonth *LunarMonth) GetPositionXi() string {
	return LunarUtil.POSITION_XI[lunarMonth.GetGanIndex()+1]
}

func (lunarMonth *LunarMonth) GetPositionXiDesc() string {
	return LunarUtil.POSITION_DESC[lunarMonth.GetPositionXi()]
}

func (lunarMonth *LunarMonth) GetPositionYangGui() string {
	return LunarUtil.POSITION_YANG_GUI[lunarMonth.GetGanIndex()+1]
}

func (lunarMonth *LunarMonth) GetPositionYangGuiDesc() string {
	return LunarUtil.POSITION_DESC[lunarMonth.GetPositionYangGui()]
}

func (lunarMonth *LunarMonth) GetPositionYinGui() string {
	return LunarUtil.POSITION_YIN_GUI[lunarMonth.GetGanIndex()+1]
}

func (lunarMonth *LunarMonth) GetPositionYinGuiDesc() string {
	return LunarUtil.POSITION_DESC[lunarMonth.GetPositionYinGui()]
}

func (lunarMonth *LunarMonth) GetPositionFu() string {
	return lunarMonth.GetPositionFuBySect(2)
}

func (lunarMonth *LunarMonth) GetPositionFuBySect(sect int) string {
	offset := lunarMonth.GetGanIndex() + 1
	if 1 == sect {
		return LunarUtil.POSITION_FU[offset]
	}
	return LunarUtil.POSITION_FU_2[offset]
}

func (lunarMonth *LunarMonth) GetPositionFuDesc() string {
	return lunarMonth.GetPositionFuDescBySect(2)
}

func (lunarMonth *LunarMonth) GetPositionFuDescBySect(sect int) string {
	return LunarUtil.POSITION_DESC[lunarMonth.GetPositionFuBySect(sect)]
}

func (lunarMonth *LunarMonth) GetPositionCai() string {
	return LunarUtil.POSITION_CAI[lunarMonth.GetGanIndex()+1]
}

func (lunarMonth *LunarMonth) GetPositionCaiDesc() string {
	return LunarUtil.POSITION_DESC[lunarMonth.GetPositionCai()]
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

func (lunarMonth *LunarMonth) GetPositionTaiSui() string {
	month := lunarMonth.month
	if month < 0 {
		month = -month
	}
	m := month % 4
	p := ""
	switch m {
	case 0:
		p = "巽"
		break
	case 1:
		p = "艮"
		break
	case 3:
		p = "坤"
		break
	default:
		p = LunarUtil.POSITION_GAN[NewSolarFromJulianDay(lunarMonth.GetFirstJulianDay()).GetLunar().GetMonthGanIndex()]
	}
	return p
}

func (lunarMonth *LunarMonth) GetPositionTaiSuiDesc() string {
	return LunarUtil.POSITION_DESC[lunarMonth.GetPositionTaiSui()]
}

func (lunarMonth *LunarMonth) GetNineStar() *NineStar {
	index := NewLunarYear(lunarMonth.year).GetZhiIndex() % 3
	m := lunarMonth.month
	if m < 0 {
		m = -m
	}
	monthZhiIndex := (13 + m) % 12
	n := 27 - index*3
	if monthZhiIndex < LunarUtil.BASE_MONTH_ZHI_INDEX {
		n -= 3
	}
	offset := (n - monthZhiIndex) % 9
	return NewNineStar(offset)
}

func (lunarMonth *LunarMonth) Next(n int) *LunarMonth {
	if 0 == n {
		return NewLunarMonthFromYm(lunarMonth.year, lunarMonth.month)
	} else if n > 0 {
		rest := n
		ny := lunarMonth.year
		iy := ny
		im := lunarMonth.month
		index := 0
		months := NewLunarYear(ny).GetMonths()
		for {
			i := 0
			size := months.Len()
			for o := months.Front(); o != nil; o = o.Next() {
				m := o.Value.(*LunarMonth)
				if m.GetYear() == iy && m.GetMonth() == im {
					index = i
					break
				}
				i++
			}
			more := size - index - 1
			if rest < more {
				break
			}
			rest -= more
			lastMonth := months.Back().Value.(*LunarMonth)
			iy = lastMonth.GetYear()
			im = lastMonth.GetMonth()
			ny++
			months = NewLunarYear(ny).GetMonths()
		}
		i := 0
		offset := index + rest
		for o := months.Front(); o != nil; o = o.Next() {
			if i == offset {
				return o.Value.(*LunarMonth)
			}
			i++
		}
		return nil
	} else {
		rest := -n
		ny := lunarMonth.year
		iy := ny
		im := lunarMonth.month
		index := 0
		months := NewLunarYear(ny).GetMonths()
		for {
			i := 0
			for o := months.Front(); o != nil; o = o.Next() {
				m := o.Value.(*LunarMonth)
				if m.GetYear() == iy && m.GetMonth() == im {
					index = i
					break
				}
				i++
			}
			if rest <= index {
				break
			}
			rest -= index
			firstMonth := months.Front().Value.(*LunarMonth)
			iy = firstMonth.GetYear()
			im = firstMonth.GetMonth()
			ny--
			months = NewLunarYear(ny).GetMonths()
		}
		i := 0
		offset := index - rest
		for o := months.Front(); o != nil; o = o.Next() {
			if i == offset {
				return o.Value.(*LunarMonth)
			}
			i++
		}
		return nil
	}
}
