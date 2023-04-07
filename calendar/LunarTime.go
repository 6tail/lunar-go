package calendar

import (
	"container/list"
	"fmt"
	"github.com/6tail/lunar-go/LunarUtil"
	"strings"
)

// LunarTime 时辰
type LunarTime struct {
	ganIndex int
	zhiIndex int
	lunar    *Lunar
}

func NewLunarTime(lunarYear int, lunarMonth int, lunarDay int, hour int, minute int, second int) *LunarTime {
	lunarTime := new(LunarTime)
	lunarTime.lunar = NewLunar(lunarYear, lunarMonth, lunarDay, hour, minute, second)
	lunarTime.zhiIndex = LunarUtil.GetTimeZhiIndex(fmt.Sprintf("%02d:%02d", hour, minute))
	lunarTime.ganIndex = (lunarTime.lunar.GetDayGanIndexExact()%5*2 + lunarTime.zhiIndex) % 10
	return lunarTime
}

func (lunarTime *LunarTime) GetGan() string {
	return LunarUtil.GAN[lunarTime.ganIndex+1]
}

func (lunarTime *LunarTime) GetZhi() string {
	return LunarUtil.ZHI[lunarTime.zhiIndex+1]
}

func (lunarTime *LunarTime) GetGanZhi() string {
	return lunarTime.GetGan() + lunarTime.GetZhi()
}

func (lunarTime *LunarTime) GetShengXiao() string {
	return LunarUtil.SHENG_XIAO[lunarTime.zhiIndex+1]
}

func (lunarTime *LunarTime) GetPositionXi() string {
	return LunarUtil.POSITION_XI[lunarTime.ganIndex+1]
}

func (lunarTime *LunarTime) GetPositionXiDesc() string {
	return LunarUtil.POSITION_DESC[lunarTime.GetPositionXi()]
}

func (lunarTime *LunarTime) GetPositionYangGui() string {
	return LunarUtil.POSITION_YANG_GUI[lunarTime.ganIndex+1]
}

func (lunarTime *LunarTime) GetPositionYangGuiDesc() string {
	return LunarUtil.POSITION_DESC[lunarTime.GetPositionYangGui()]
}

func (lunarTime *LunarTime) GetPositionYinGui() string {
	return LunarUtil.POSITION_YIN_GUI[lunarTime.ganIndex+1]
}

func (lunarTime *LunarTime) GetPositionYinGuiDesc() string {
	return LunarUtil.POSITION_DESC[lunarTime.GetPositionYinGui()]
}

func (lunarTime *LunarTime) GetPositionFu() string {
	return lunarTime.GetPositionFuBySect(2)
}

func (lunarTime *LunarTime) GetPositionFuBySect(sect int) string {
	offset := lunarTime.ganIndex + 1
	if 1 == sect {
		return LunarUtil.POSITION_FU[offset]
	}
	return LunarUtil.POSITION_FU_2[offset]
}

func (lunarTime *LunarTime) GetPositionFuDesc() string {
	return lunarTime.GetPositionFuDescBySect(2)
}

func (lunarTime *LunarTime) GetPositionFuDescBySect(sect int) string {
	return LunarUtil.POSITION_DESC[lunarTime.GetPositionFuBySect(sect)]
}

func (lunarTime *LunarTime) GetPositionCai() string {
	return LunarUtil.POSITION_CAI[lunarTime.ganIndex+1]
}

func (lunarTime *LunarTime) GetPositionCaiDesc() string {
	return LunarUtil.POSITION_DESC[lunarTime.GetPositionCai()]
}

func (lunarTime *LunarTime) GetNaYin() string {
	return LunarUtil.NAYIN[lunarTime.GetGanZhi()]
}

func (lunarTime *LunarTime) GetTianShen() string {
	return LunarUtil.TIAN_SHEN[(lunarTime.zhiIndex+LunarUtil.ZHI_TIAN_SHEN_OFFSET[lunarTime.lunar.GetDayZhiExact()])%12+1]
}

func (lunarTime *LunarTime) GetTianShenType() string {
	return LunarUtil.TIAN_SHEN_TYPE[lunarTime.GetTianShen()]
}

func (lunarTime *LunarTime) GetTianShenLuck() string {
	return LunarUtil.TIAN_SHEN_TYPE_LUCK[lunarTime.GetTianShenType()]
}

func (lunarTime *LunarTime) GetChong() string {
	return LunarUtil.CHONG[lunarTime.zhiIndex]
}

func (lunarTime *LunarTime) GetSha() string {
	return LunarUtil.SHA[lunarTime.GetZhi()]
}

func (lunarTime *LunarTime) GetChongGan() string {
	return LunarUtil.CHONG_GAN[lunarTime.ganIndex]
}

func (lunarTime *LunarTime) GetChongGanTie() string {
	return LunarUtil.CHONG_GAN_TIE[lunarTime.ganIndex]
}

func (lunarTime *LunarTime) GetChongShengXiao() string {
	chong := lunarTime.GetChong()
	for i, v := range LunarUtil.ZHI {
		if v == chong {
			return LunarUtil.SHENG_XIAO[i]
		}
	}
	return ""
}

func (lunarTime *LunarTime) GetChongDesc() string {
	return "(" + lunarTime.GetChongGan() + lunarTime.GetChong() + ")" + lunarTime.GetChongShengXiao()
}

func (lunarTime *LunarTime) GetYi() *list.List {
	return LunarUtil.GetTimeYi(lunarTime.lunar.GetDayInGanZhiExact(), lunarTime.GetGanZhi())
}

func (lunarTime *LunarTime) GetJi() *list.List {
	return LunarUtil.GetTimeJi(lunarTime.lunar.GetDayInGanZhiExact(), lunarTime.GetGanZhi())
}

func (lunarTime *LunarTime) GetNineStar() *NineStar {
	//顺逆
	solarYmd := lunarTime.lunar.GetSolar().ToYmd()
	jieQi := lunarTime.lunar.GetJieQiTable()
	asc := false
	if strings.Compare(solarYmd, jieQi["冬至"].ToYmd()) >= 0 && strings.Compare(solarYmd, jieQi["夏至"].ToYmd()) < 0 {
		asc = true
	}
	start := 3
	if asc {
		start = 7
	}
	dayZhi := lunarTime.lunar.GetDayZhi()
	if strings.Index("子午卯酉", dayZhi) > -1 {
		if asc {
			start = 1
		} else {
			start = 9
		}
	} else if strings.Index("辰戌丑未", dayZhi) > -1 {
		if asc {
			start = 4
		} else {
			start = 6
		}
	}
	index := start - lunarTime.zhiIndex - 1
	if asc {
		index = start + lunarTime.zhiIndex - 1
	}
	if index > 8 {
		index -= 9
	}
	if index < 0 {
		index += 9
	}
	return NewNineStar(index)
}

func (lunarTime *LunarTime) String() string {
	return lunarTime.ToString()
}

func (lunarTime *LunarTime) ToString() string {
	return lunarTime.GetGanZhi()
}

func (lunarTime *LunarTime) GetGanIndex() int {
	return lunarTime.ganIndex
}

func (lunarTime *LunarTime) GetZhiIndex() int {
	return lunarTime.zhiIndex
}

// GetXun 获取时辰所在旬
func (lunarTime *LunarTime) GetXun() string {
	return LunarUtil.GetXun(lunarTime.GetGanZhi())
}

// GetXunKong 获取值时空亡
func (lunarTime *LunarTime) GetXunKong() string {
	return LunarUtil.GetXunKong(lunarTime.GetGanZhi())
}

// GetMinHm 获取当前时辰的最早时分
func (lunarTime *LunarTime) GetMinHm() string {
	hour := lunarTime.lunar.GetHour()
	if hour < 1 {
		return "00:00"
	} else if hour > 22 {
		return "23:00"
	}
	if hour%2 == 0 {
		hour -= 1
	}
	return fmt.Sprintf("%02d:00", hour)
}

// GetMaxHm 获取当前时辰的最晚时分
func (lunarTime *LunarTime) GetMaxHm() string {
	hour := lunarTime.lunar.GetHour()
	if hour < 1 {
		return "00:59"
	} else if hour > 22 {
		return "23:59"
	}
	if hour%2 != 0 {
		hour += 1
	}
	return fmt.Sprintf("%02d:59", hour)
}
