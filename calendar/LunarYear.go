package calendar

import (
	"container/list"
	"fmt"
	"github.com/6tail/lunar-go/LunarUtil"
	"github.com/6tail/lunar-go/ShouXingUtil"
	"math"
	"strings"
	"sync"
)

// YUAN 元
var YUAN = []string{"下", "上", "中"}

// YUN 运
var YUN = []string{"七", "八", "九", "一", "二", "三", "四", "五", "六"}

var LEAP_11 = []int{75, 94, 170, 265, 322, 398, 469, 553, 583, 610, 678, 735, 754, 773, 849, 887, 936, 1050, 1069, 1126, 1145, 1164, 1183, 1259, 1278, 1308, 1373, 1403, 1441, 1460, 1498, 1555, 1593, 1612, 1631, 1642, 2033, 2128, 2147, 2242, 2614, 2728, 2910, 3062, 3244, 3339, 3616, 3711, 3730, 3825, 4007, 4159, 4197, 4322, 4341, 4379, 4417, 4531, 4599, 4694, 4713, 4789, 4808, 4971, 5085, 5104, 5161, 5180, 5199, 5294, 5305, 5476, 5677, 5696, 5772, 5791, 5848, 5886, 6049, 6068, 6144, 6163, 6258, 6402, 6440, 6497, 6516, 6630, 6641, 6660, 6679, 6736, 6774, 6850, 6869, 6899, 6918, 6994, 7013, 7032, 7051, 7070, 7089, 7108, 7127, 7146, 7222, 7271, 7290, 7309, 7366, 7385, 7404, 7442, 7461, 7480, 7491, 7499, 7594, 7624, 7643, 7662, 7681, 7719, 7738, 7814, 7863, 7882, 7901, 7939, 7958, 7977, 7996, 8034, 8053, 8072, 8091, 8121, 8159, 8186, 8216, 8235, 8254, 8273, 8311, 8330, 8341, 8349, 8368, 8444, 8463, 8474, 8493, 8531, 8569, 8588, 8626, 8664, 8683, 8694, 8702, 8713, 8721, 8751, 8789, 8808, 8816, 8827, 8846, 8884, 8903, 8922, 8941, 8971, 9036, 9066, 9085, 9104, 9123, 9142, 9161, 9180, 9199, 9218, 9256, 9294, 9313, 9324, 9343, 9362, 9381, 9419, 9438, 9476, 9514, 9533, 9544, 9552, 9563, 9571, 9582, 9601, 9639, 9658, 9666, 9677, 9696, 9734, 9753, 9772, 9791, 9802, 9821, 9886, 9897, 9916, 9935, 9954, 9973, 9992}
var LEAP_12 = []int{37, 56, 113, 132, 151, 189, 208, 227, 246, 284, 303, 341, 360, 379, 417, 436, 458, 477, 496, 515, 534, 572, 591, 629, 648, 667, 697, 716, 792, 811, 830, 868, 906, 925, 944, 963, 982, 1001, 1020, 1039, 1058, 1088, 1153, 1202, 1221, 1240, 1297, 1335, 1392, 1411, 1422, 1430, 1517, 1525, 1536, 1574, 3358, 3472, 3806, 3988, 4751, 4941, 5066, 5123, 5275, 5343, 5438, 5457, 5495, 5533, 5552, 5715, 5810, 5829, 5905, 5924, 6421, 6535, 6793, 6812, 6888, 6907, 7002, 7184, 7260, 7279, 7374, 7556, 7746, 7757, 7776, 7833, 7852, 7871, 7966, 8015, 8110, 8129, 8148, 8224, 8243, 8338, 8406, 8425, 8482, 8501, 8520, 8558, 8596, 8607, 8615, 8645, 8740, 8778, 8835, 8865, 8930, 8960, 8979, 8998, 9017, 9055, 9074, 9093, 9112, 9150, 9188, 9237, 9275, 9332, 9351, 9370, 9408, 9427, 9446, 9457, 9465, 9495, 9560, 9590, 9628, 9647, 9685, 9715, 9742, 9780, 9810, 9818, 9829, 9848, 9867, 9905, 9924, 9943, 9962, 10000}
var YMC = []int{11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var CACHE_YEAR *LunarYear = nil
var lock sync.Mutex

// LunarYear 阴历年
type LunarYear struct {
	year            int
	ganIndex        int
	zhiIndex        int
	months          *list.List
	jieQiJulianDays []float64
}

func NewLunarYear(lunarYear int) *LunarYear {
	lock.Lock()
	var year *LunarYear
	if nil == CACHE_YEAR || CACHE_YEAR.year != lunarYear {
		year = new(LunarYear)
		year.year = lunarYear
		year.months = list.New()
		offset := lunarYear - 4
		yearGanIndex := offset % 10
		yearZhiIndex := offset % 12
		if yearGanIndex < 0 {
			yearGanIndex += 10
		}
		if yearZhiIndex < 0 {
			yearZhiIndex += 12
		}
		year.ganIndex = yearGanIndex
		year.zhiIndex = yearZhiIndex
		year.compute()
		CACHE_YEAR = year
	} else {
		year = CACHE_YEAR
	}
	lock.Unlock()
	return year
}

func contains(arr []int, n int) bool {
	j := len(arr)
	for i := 0; i < j; i++ {
		if n == arr[i] {
			return true
		}
	}
	return false
}

func (lunarYear *LunarYear) compute() {
	// 节气
	jq := make([]float64, 27)
	// 合朔，即每月初一
	hs := make([]float64, 16)
	// 每月天数，长度15
	dayCounts := make([]int, 15)
	months := make([]int, 15)

	currentYear := lunarYear.year
	jd := math.Floor(float64(currentYear-2000)*365.2422 + float64(180))
	// 355是2000.12冬至，得到较靠近jd的冬至估计值
	w := math.Floor((jd-355+183)/365.2422)*365.2422 + 355
	if ShouXingUtil.CalcQi(w) > jd {
		w -= 365.2422
	}
	// 25个节气时刻(北京时间)，从冬至开始到下一个冬至以后
	i := 0
	for i = 0; i < 26; i++ {
		jq[i] = ShouXingUtil.CalcQi(w + 15.2184*float64(i))
	}
	// 从上年的大雪到下年的立春
	j := len(JIE_QI_IN_USE)
	lunarYear.jieQiJulianDays = make([]float64, j)
	for i = 0; i < j; i++ {
		if i == 0 {
			jd = ShouXingUtil.QiAccurate2(jq[0] - 15.2184)
		} else if i <= 26 {
			jd = ShouXingUtil.QiAccurate2(jq[i-1])
		} else {
			jd = ShouXingUtil.QiAccurate2(jq[25] + 15.2184*float64(i-26))
		}
		lunarYear.jieQiJulianDays[i] = jd + J2000
	}

	// 冬至前的初一，今年"首朔"的日月黄经差w
	w = ShouXingUtil.CalcShuo(jq[0])
	if w > jq[0] {
		w -= 29.53
	}
	// 递推每月初一
	for i = 0; i < 16; i++ {
		hs[i] = ShouXingUtil.CalcShuo(w + 29.5306*float64(i))
	}
	// 每月
	for i = 0; i < 15; i++ {
		dayCounts[i] = (int)(hs[i+1] - hs[i])
		months[i] = i
	}

	prevYear := currentYear - 1
	leapIndex := 16
	if contains(LEAP_11, currentYear) {
		leapIndex = 13
	} else if contains(LEAP_12, currentYear) {
		leapIndex = 14
	} else if hs[13] <= jq[24] {
		i = 1
		for {
			if hs[i+1] <= jq[2*i] {
				break
			}
			if i >= 13 {
				break
			}
			i++
		}
		leapIndex = i
	}
	for i = leapIndex; i < 15; i++ {
		months[i] -= 1
	}

	fm := -1
	index := -1
	y := prevYear
	for i = 0; i < 15; i++ {
		dm := hs[i] + J2000
		v2 := months[i]
		mc := YMC[v2%12]
		if 1724360 <= dm && dm < 1729794 {
			mc = YMC[(v2+1)%12]
		} else if 1807724 <= dm && dm < 1808699 {
			mc = YMC[(v2+1)%12]
		} else if dm == 1729794 || dm == 1808699 {
			mc = 12
		}
		if fm == -1 {
			fm = mc
			index = mc
		}
		if mc < fm {
			y += 1
			index = 1
		}
		fm = mc
		if i == leapIndex {
			mc = -mc
		} else if dm == 1729794 || dm == 1808699 {
			mc = -11
		}
		lunarYear.months.PushBack(NewLunarMonth(y, mc, dayCounts[i], hs[i]+J2000, index))
		index++
	}
}

func (lunarYear *LunarYear) GetYear() int {
	return lunarYear.year
}

func (lunarYear *LunarYear) GetGanIndex() int {
	return lunarYear.ganIndex
}

func (lunarYear *LunarYear) GetZhiIndex() int {
	return lunarYear.zhiIndex
}

func (lunarYear *LunarYear) GetGan() string {
	return LunarUtil.GAN[lunarYear.ganIndex+1]
}

func (lunarYear *LunarYear) GetZhi() string {
	return LunarUtil.ZHI[lunarYear.zhiIndex+1]
}

func (lunarYear *LunarYear) GetGanZhi() string {
	return fmt.Sprintf("%s%s", lunarYear.GetGan(), lunarYear.GetZhi())
}

func (lunarYear *LunarYear) GetMonths() *list.List {
	return lunarYear.months
}

func (lunarYear *LunarYear) GetMonthsInYear() *list.List {
	l := list.New()
	for i := lunarYear.months.Front(); i != nil; i = i.Next() {
		m := i.Value.(*LunarMonth)
		if m.GetYear() == lunarYear.year {
			l.PushBack(m)
		}
	}
	return l
}

func (lunarYear *LunarYear) GetDayCount() int {
	n := 0
	for i := lunarYear.months.Front(); i != nil; i = i.Next() {
		m := i.Value.(*LunarMonth)
		if m.GetYear() == lunarYear.year {
			n += m.GetDayCount()
		}
	}
	return n
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
	return fmt.Sprintf("%d", lunarYear.year)
}

func (lunarYear *LunarYear) ToFullString() string {
	return fmt.Sprintf("%d年", lunarYear.year)
}

func (lunarYear *LunarYear) getZaoByGan(index int, name string) string {
	offset := index - NewSolarFromJulianDay(lunarYear.GetMonth(1).GetFirstJulianDay()).GetLunar().GetDayGanIndex()
	if offset < 0 {
		offset += 10
	}
	return strings.Replace(name, "几", LunarUtil.NUMBER[offset+1], 1)
}

func (lunarYear *LunarYear) getZaoByZhi(index int, name string) string {
	offset := index - NewSolarFromJulianDay(lunarYear.GetMonth(1).GetFirstJulianDay()).GetLunar().GetDayZhiIndex()
	if offset < 0 {
		offset += 12
	}
	return strings.Replace(name, "几", LunarUtil.NUMBER[offset+1], 1)
}

func (lunarYear *LunarYear) GetTouLiang() string {
	return lunarYear.getZaoByZhi(0, "几鼠偷粮")
}

func (lunarYear *LunarYear) GetCaoZi() string {
	return lunarYear.getZaoByZhi(0, "草子几分")
}

func (lunarYear *LunarYear) GetGengTian() string {
	return lunarYear.getZaoByZhi(1, "几牛耕田")
}

func (lunarYear *LunarYear) GetHuaShou() string {
	return lunarYear.getZaoByZhi(3, "花收几分")
}

func (lunarYear *LunarYear) GetZhiShui() string {
	return lunarYear.getZaoByZhi(4, "几龙治水")
}

func (lunarYear *LunarYear) GetTuoGu() string {
	return lunarYear.getZaoByZhi(6, "几马驮谷")
}

func (lunarYear *LunarYear) GetQiangMi() string {
	return lunarYear.getZaoByZhi(9, "几鸡抢米")
}

func (lunarYear *LunarYear) GetKanCan() string {
	return lunarYear.getZaoByZhi(9, "几姑看蚕")
}

func (lunarYear *LunarYear) GetGongZhu() string {
	return lunarYear.getZaoByZhi(11, "几屠共猪")
}

func (lunarYear *LunarYear) GetJiaTian() string {
	return lunarYear.getZaoByGan(0, "甲田几分")
}

func (lunarYear *LunarYear) GetFenBing() string {
	return lunarYear.getZaoByGan(2, "几人分饼")
}

func (lunarYear *LunarYear) GetDeJin() string {
	return lunarYear.getZaoByGan(7, "几日得金")
}

func (lunarYear *LunarYear) GetRenBing() string {
	return lunarYear.getZaoByGan(2, lunarYear.getZaoByZhi(2, "几人几丙"))
}

func (lunarYear *LunarYear) GetRenChu() string {
	return lunarYear.getZaoByGan(3, lunarYear.getZaoByZhi(2, "几人几锄"))
}

func (lunarYear *LunarYear) GetYuan() string {
	return fmt.Sprintf("%s元", YUAN[int((lunarYear.year+2696)/60)%3])
}

func (lunarYear *LunarYear) GetYun() string {
	return fmt.Sprintf("%s运", YUN[int((lunarYear.year+2696)/20)%9])
}

func (lunarYear *LunarYear) GetNineStar() *NineStar {
	index := LunarUtil.GetJiaZiIndex(lunarYear.GetGanZhi()) + 1
	yuan := int((lunarYear.year+2696)/60) % 3
	offset := (62 + yuan*3 - index) % 9
	if 0 == offset {
		offset = 9
	}
	return NewNineStar(offset - 1)
}

func (lunarYear *LunarYear) GetPositionXi() string {
	return LunarUtil.POSITION_XI[lunarYear.ganIndex+1]
}

func (lunarYear *LunarYear) GetPositionXiDesc() string {
	return LunarUtil.POSITION_DESC[lunarYear.GetPositionXi()]
}

func (lunarYear *LunarYear) GetPositionYangGui() string {
	return LunarUtil.POSITION_YANG_GUI[lunarYear.ganIndex+1]
}

func (lunarYear *LunarYear) GetPositionYangGuiDesc() string {
	return LunarUtil.POSITION_DESC[lunarYear.GetPositionYangGui()]
}

func (lunarYear *LunarYear) GetPositionYinGui() string {
	return LunarUtil.POSITION_YIN_GUI[lunarYear.ganIndex+1]
}

func (lunarYear *LunarYear) GetPositionYinGuiDesc() string {
	return LunarUtil.POSITION_DESC[lunarYear.GetPositionYinGui()]
}

func (lunarYear *LunarYear) GetPositionFu() string {
	return lunarYear.GetPositionFuBySect(2)
}

func (lunarYear *LunarYear) GetPositionFuBySect(sect int) string {
	offset := lunarYear.ganIndex + 1
	if 1 == sect {
		return LunarUtil.POSITION_FU[offset]
	}
	return LunarUtil.POSITION_FU_2[offset]
}

func (lunarYear *LunarYear) GetPositionFuDesc() string {
	return lunarYear.GetPositionFuDescBySect(2)
}

func (lunarYear *LunarYear) GetPositionFuDescBySect(sect int) string {
	return LunarUtil.POSITION_DESC[lunarYear.GetPositionFuBySect(sect)]
}

func (lunarYear *LunarYear) GetPositionCai() string {
	return LunarUtil.POSITION_CAI[lunarYear.ganIndex+1]
}

func (lunarYear *LunarYear) GetPositionCaiDesc() string {
	return LunarUtil.POSITION_DESC[lunarYear.GetPositionCai()]
}

func (lunarYear *LunarYear) GetPositionTaiSui() string {
	return LunarUtil.POSITION_TAI_SUI_YEAR[lunarYear.zhiIndex]
}

func (lunarYear *LunarYear) GetPositionTaiSuiDesc() string {
	return LunarUtil.POSITION_DESC[lunarYear.GetPositionTaiSui()]
}

func (lunarYear *LunarYear) Next(n int) *LunarYear {
	return NewLunarYear(lunarYear.year + n)
}
