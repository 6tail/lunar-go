package calendar

import (
	"container/list"
	"github.com/6tail/lunar-go/LunarUtil"
	"strings"
)

// 月支，按正月起寅排列
var MONTH_ZHI = []string{"", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"}

// 长生十二神
var CHANG_SHENG = []string{"长生", "沐浴", "冠带", "临官", "帝旺", "衰", "病", "死", "墓", "绝", "胎", "养"}

// 长生十二神日干偏移值，五阳干顺推，五阴干逆推
var changShengOffset = map[string]int{
	"甲": 11,
	"丙": 2,
	"戊": 2,
	"庚": 5,
	"壬": 8,
	"乙": 6,
	"丁": 9,
	"己": 9,
	"辛": 0,
	"癸": 3,
}

// 八字
type EightChar struct {
	lunar *Lunar
}

func NewEightChar(lunar *Lunar) *EightChar {
	eightChar := new(EightChar)
	eightChar.lunar = lunar
	return eightChar
}

func (eightChar *EightChar) String() string {
	return eightChar.GetYear() + " " + eightChar.GetMonth() + " " + eightChar.GetDay() + " " + eightChar.GetTime()
}

func (eightChar *EightChar) getShiShenZhi(zhi string) *list.List {
	l := list.New()
	hideGan := LunarUtil.ZHI_HIDE_GAN[zhi]
	for i := 0; i < len(hideGan); i++ {
		l.PushBack(LunarUtil.SHI_SHEN_ZHI[eightChar.GetDayGan()+zhi+hideGan[i]])
	}
	return l
}

func (eightChar *EightChar) getDiShi(zhiIndex int) string {
	offset := changShengOffset[eightChar.GetDayGan()]
	index := offset
	if eightChar.lunar.GetDayGanIndexExact()%2 == 0 {
		index += zhiIndex
	} else {
		index -= zhiIndex
	}
	if index >= 12 {
		index -= 12
	}
	if index < 0 {
		index += 12
	}
	return CHANG_SHENG[index]
}

func (eightChar *EightChar) GetYear() string {
	return eightChar.lunar.GetYearInGanZhiExact()
}

func (eightChar *EightChar) GetYearGan() string {
	return eightChar.lunar.GetYearGanExact()
}

func (eightChar *EightChar) GetYearZhi() string {
	return eightChar.lunar.GetYearZhiExact()
}

func (eightChar *EightChar) GetYearHideGan() []string {
	return LunarUtil.ZHI_HIDE_GAN[eightChar.GetYearZhi()]
}

func (eightChar *EightChar) GetYearWuXing() string {
	return LunarUtil.WU_XING_GAN[eightChar.GetYearGan()] + LunarUtil.WU_XING_ZHI[eightChar.GetYearZhi()]
}

func (eightChar *EightChar) GetYearNaYin() string {
	return LunarUtil.NAYIN[eightChar.GetYear()]
}

func (eightChar *EightChar) GetYearShiShenGan() string {
	return LunarUtil.SHI_SHEN_GAN[eightChar.GetDayGan()+eightChar.GetYearGan()]
}

func (eightChar *EightChar) GetYearShiShenZhi() *list.List {
	return eightChar.getShiShenZhi(eightChar.GetYearZhi())
}

func (eightChar *EightChar) GetYearDiShi() string {
	return eightChar.getDiShi(eightChar.lunar.GetYearZhiIndexExact())
}

func (eightChar *EightChar) GetMonth() string {
	return eightChar.lunar.GetMonthInGanZhiExact()
}

func (eightChar *EightChar) GetMonthGan() string {
	return eightChar.lunar.GetMonthGanExact()
}

func (eightChar *EightChar) GetMonthZhi() string {
	return eightChar.lunar.GetMonthZhiExact()
}

func (eightChar *EightChar) GetMonthHideGan() []string {
	return LunarUtil.ZHI_HIDE_GAN[eightChar.GetMonthZhi()]
}

func (eightChar *EightChar) GetMonthWuXing() string {
	return LunarUtil.WU_XING_GAN[eightChar.GetMonthGan()] + LunarUtil.WU_XING_ZHI[eightChar.GetMonthZhi()]
}

func (eightChar *EightChar) GetMonthNaYin() string {
	return LunarUtil.NAYIN[eightChar.GetMonth()]
}

func (eightChar *EightChar) GetMonthShiShenGan() string {
	return LunarUtil.SHI_SHEN_GAN[eightChar.GetDayGan()+eightChar.GetMonthGan()]
}

func (eightChar *EightChar) GetMonthShiShenZhi() *list.List {
	return eightChar.getShiShenZhi(eightChar.GetMonthZhi())
}

func (eightChar *EightChar) GetMonthDiShi() string {
	return eightChar.getDiShi(eightChar.lunar.GetMonthZhiIndexExact())
}

func (eightChar *EightChar) GetDay() string {
	return eightChar.lunar.GetDayInGanZhiExact()
}

func (eightChar *EightChar) GetDayGan() string {
	return eightChar.lunar.GetDayGanExact()
}

func (eightChar *EightChar) GetDayZhi() string {
	return eightChar.lunar.GetDayZhiExact()
}

func (eightChar *EightChar) GetDayHideGan() []string {
	return LunarUtil.ZHI_HIDE_GAN[eightChar.GetDayZhi()]
}

func (eightChar *EightChar) GetDayWuXing() string {
	return LunarUtil.WU_XING_GAN[eightChar.GetDayGan()] + LunarUtil.WU_XING_ZHI[eightChar.GetDayZhi()]
}

func (eightChar *EightChar) GetDayNaYin() string {
	return LunarUtil.NAYIN[eightChar.GetDay()]
}

func (eightChar *EightChar) GetDayShiShenGan() string {
	return "日主"
}

func (eightChar *EightChar) GetDayShiShenZhi() *list.List {
	return eightChar.getShiShenZhi(eightChar.GetDayZhi())
}

func (eightChar *EightChar) GetDayDiShi() string {
	return eightChar.getDiShi(eightChar.lunar.GetDayZhiIndexExact())
}

func (eightChar *EightChar) GetTime() string {
	return eightChar.lunar.GetTimeInGanZhi()
}

func (eightChar *EightChar) GetTimeGan() string {
	return eightChar.lunar.GetTimeGan()
}

func (eightChar *EightChar) GetTimeZhi() string {
	return eightChar.lunar.GetTimeZhi()
}

func (eightChar *EightChar) GetTimeHideGan() []string {
	return LunarUtil.ZHI_HIDE_GAN[eightChar.GetTimeZhi()]
}

func (eightChar *EightChar) GetTimeWuXing() string {
	return LunarUtil.WU_XING_GAN[eightChar.GetTimeGan()] + LunarUtil.WU_XING_ZHI[eightChar.GetTimeZhi()]
}

func (eightChar *EightChar) GetTimeNaYin() string {
	return LunarUtil.NAYIN[eightChar.GetTime()]
}

func (eightChar *EightChar) GetTimeShiShenGan() string {
	return LunarUtil.SHI_SHEN_GAN[eightChar.GetDayGan()+eightChar.GetTimeGan()]
}

func (eightChar *EightChar) GetTimeShiShenZhi() *list.List {
	return eightChar.getShiShenZhi(eightChar.GetTimeZhi())
}

func (eightChar *EightChar) GetTimeDiShi() string {
	return eightChar.getDiShi(eightChar.lunar.GetTimeZhiIndex())
}

func (eightChar *EightChar) GetTaiYuan() string {
	ganIndex := eightChar.lunar.GetMonthGanIndexExact() + 1
	if ganIndex >= 10 {
		ganIndex -= 10
	}
	zhiIndex := eightChar.lunar.GetMonthZhiIndexExact() + 3
	if zhiIndex >= 12 {
		zhiIndex -= 12
	}
	return LunarUtil.GAN[ganIndex+1] + LunarUtil.ZHI[zhiIndex+1]
}

func (eightChar *EightChar) GetTaiYuanNaYin() string {
	return LunarUtil.NAYIN[eightChar.GetTaiYuan()]
}

func (eightChar *EightChar) GetMingGong() string {
	monthZhiIndex := 0
	timeZhiIndex := 0
	size := len(MONTH_ZHI)
	for i := 0; i < size; i++ {
		zhi := MONTH_ZHI[i]
		if strings.Compare(eightChar.GetMonthZhi(), zhi) == 0 {
			monthZhiIndex = i
		}
		if strings.Compare(eightChar.GetTimeZhi(), zhi) == 0 {
			timeZhiIndex = i
		}
	}
	zhiIndex := 26 - (monthZhiIndex + timeZhiIndex)
	if zhiIndex > 12 {
		zhiIndex -= 12
	}
	jiaZiIndex := LunarUtil.GetJiaZiIndex(eightChar.lunar.GetMonthInGanZhiExact()) - (monthZhiIndex - zhiIndex)
	if jiaZiIndex >= 60 {
		jiaZiIndex -= 60
	}
	if jiaZiIndex < 0 {
		jiaZiIndex += 60
	}
	return LunarUtil.JIA_ZI[jiaZiIndex]
}

func (eightChar *EightChar) GetMingGongNaYin() string {
	return LunarUtil.NAYIN[eightChar.GetMingGong()]
}

func (eightChar *EightChar) GetShenGong() string {
	monthZhiIndex := 0
	timeZhiIndex := 0
	size := len(MONTH_ZHI)
	for i := 0; i < size; i++ {
		zhi := MONTH_ZHI[i]
		if strings.Compare(eightChar.GetMonthZhi(), zhi) == 0 {
			monthZhiIndex = i
		}
		if strings.Compare(eightChar.GetTimeZhi(), zhi) == 0 {
			timeZhiIndex = i
		}
	}
	zhiIndex := (2 + (monthZhiIndex + timeZhiIndex)) % 12
	jiaZiIndex := LunarUtil.GetJiaZiIndex(eightChar.lunar.GetMonthInGanZhiExact()) - (monthZhiIndex - zhiIndex)
	if jiaZiIndex >= 60 {
		jiaZiIndex -= 60
	}
	if jiaZiIndex < 0 {
		jiaZiIndex += 60
	}
	return LunarUtil.JIA_ZI[jiaZiIndex]
}

func (eightChar *EightChar) GetShenGongNaYin() string {
	return LunarUtil.NAYIN[eightChar.GetShenGong()]
}

func (eightChar *EightChar) GetLunar() *Lunar {
	return eightChar.lunar
}

// 获取运
func (eightChar *EightChar) GetYun(gender int) *Yun {
	return NewYun(eightChar, gender)
}

// 获取年柱所在旬
func (eightChar *EightChar) GetYearXun() string {
	return eightChar.lunar.GetYearXunExact()
}

// 获取年柱旬空(空亡)
func (eightChar *EightChar) GetYearXunKong() string {
	return eightChar.lunar.GetYearXunKongExact()
}

// 获取月柱所在旬
func (eightChar *EightChar) GetMonthXun() string {
	return eightChar.lunar.GetMonthXunExact()
}

// 获取月柱旬空(空亡)
func (eightChar *EightChar) GetMonthXunKong() string {
	return eightChar.lunar.GetMonthXunKongExact()
}

// 获取日柱所在旬
func (eightChar *EightChar) GetDayXun() string {
	return eightChar.lunar.GetDayXunExact()
}

// 获取日柱旬空(空亡)
func (eightChar *EightChar) GetDayXunKong() string {
	return eightChar.lunar.GetDayXunKongExact()
}

// 获取时柱所在旬
func (eightChar *EightChar) GetTimeXun() string {
	return eightChar.lunar.GetTimeXun()
}

// 获取时柱旬空(空亡)
func (eightChar *EightChar) GetTimeXunKong() string {
	return eightChar.lunar.GetTimeXunKong()
}
