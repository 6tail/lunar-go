package calendar

import (
	"container/list"
	"fmt"
	"github.com/6tail/lunar-go/LunarUtil"
	"github.com/6tail/lunar-go/ShouXingUtil"
	"math"
)

var LEAP_11 = []int{75, 94, 170, 238, 265, 322, 389, 469, 553, 583, 610, 678, 735, 754, 773, 849, 887, 936, 1050, 1069, 1126, 1145, 1164, 1183, 1259, 1278, 1308, 1373, 1403, 1441, 1460, 1498, 1555, 1593, 1612, 1631, 1642, 2033, 2128, 2147, 2242, 2614, 2728, 2910, 3062, 3244, 3339, 3616, 3711, 3730, 3825, 4007, 4159, 4197, 4322, 4341, 4379, 4417, 4531, 4599, 4694, 4713, 4789, 4808, 4971, 5085, 5104, 5161, 5180, 5199, 5294, 5305, 5476, 5677, 5696, 5772, 5791, 5848, 5886, 6049, 6068, 6144, 6163, 6258, 6402, 6440, 6497, 6516, 6630, 6641, 6660, 6679, 6736, 6774, 6850, 6869, 6899, 6918, 6994, 7013, 7032, 7051, 7070, 7089, 7108, 7127, 7146, 7222, 7271, 7290, 7309, 7366, 7385, 7404, 7442, 7461, 7480, 7491, 7499, 7594, 7624, 7643, 7662, 7681, 7719, 7738, 7814, 7863, 7882, 7901, 7939, 7958, 7977, 7996, 8034, 8053, 8072, 8091, 8121, 8159, 8186, 8216, 8235, 8254, 8273, 8311, 8330, 8341, 8349, 8368, 8444, 8463, 8474, 8493, 8531, 8569, 8588, 8626, 8664, 8683, 8694, 8702, 8713, 8721, 8751, 8789, 8808, 8816, 8827, 8846, 8884, 8903, 8922, 8941, 8971, 9036, 9066, 9085, 9104, 9123, 9142, 9161, 9180, 9199, 9218, 9256, 9294, 9313, 9324, 9343, 9362, 9381, 9419, 9438, 9476, 9514, 9533, 9544, 9552, 9563, 9571, 9582, 9601, 9639, 9658, 9666, 9677, 9696, 9734, 9753, 9772, 9791, 9802, 9821, 9886, 9897, 9916, 9935, 9954, 9973, 9992}
var LEAP_12 = []int{37, 56, 113, 132, 151, 189, 208, 227, 246, 284, 303, 341, 360, 379, 417, 436, 458, 477, 496, 515, 534, 572, 591, 629, 648, 667, 697, 716, 792, 811, 830, 868, 906, 925, 944, 963, 982, 1001, 1020, 1039, 1058, 1088, 1153, 1202, 1221, 1240, 1297, 1335, 1392, 1411, 1422, 1430, 1517, 1525, 1536, 1574, 3358, 3472, 3806, 3988, 4751, 4941, 5066, 5123, 5275, 5343, 5438, 5457, 5495, 5533, 5552, 5715, 5810, 5829, 5905, 5924, 6421, 6535, 6793, 6812, 6888, 6907, 7002, 7184, 7260, 7279, 7374, 7556, 7746, 7757, 7776, 7833, 7852, 7871, 7966, 8015, 8110, 8129, 8148, 8224, 8243, 8338, 8406, 8425, 8482, 8501, 8520, 8558, 8596, 8607, 8615, 8645, 8740, 8778, 8835, 8865, 8930, 8960, 8979, 8998, 9017, 9055, 9074, 9093, 9112, 9150, 9188, 9237, 9275, 9332, 9351, 9370, 9408, 9427, 9446, 9457, 9465, 9495, 9560, 9590, 9628, 9647, 9685, 9715, 9742, 9780, 9810, 9818, 9829, 9848, 9867, 9905, 9924, 9943, 9962, 10000}
var LEAP = map[int]int{}
var CACHE = map[int]*LunarYear{}

// 阴历年
type LunarYear struct {
	year            int
	months          *list.List
	jieQiJulianDays []float64
}

func NewLunarYear(lunarYear int) *LunarYear {
	if len(LEAP) < 1 {
		for _, value := range LEAP_11 {
			LEAP[value] = 13
		}
		for _, value := range LEAP_12 {
			LEAP[value] = 14
		}
	}
	obj, exists := CACHE[lunarYear]
	if !exists {
		obj = new(LunarYear)
		obj.year = lunarYear
		obj.months = list.New()
		obj.compute()
		CACHE[lunarYear] = obj
	}
	return obj
}

func (lunarYear *LunarYear) compute() {
	// 节气(中午12点)，长度27
	jq := make([]float64, 27)
	// 合朔，即每月初一(中午12点)，长度16
	hs := make([]float64, 16)
	// 每月天数，长度15
	dayCounts := make([]int, 15)

	currentYear := lunarYear.year
	year := currentYear - 2000

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
		if i > 0 && i < 28 {
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

	currentYearLeap, exists := LEAP[currentYear]
	if !exists {
		currentYearLeap = -1
		if hs[13] <= jq[24] {
			i := 1
			for {
				if hs[i+1] <= jq[2*i] {
					break
				}
				if i >= 13 {
					break
				}
				i++
			}
			currentYearLeap = i
		}
	}

	prevYear := currentYear - 1
	prevYearLeap, exists := LEAP[prevYear]
	if !exists {
		prevYearLeap = -1
	} else {
		prevYearLeap -= 12
	}
	y := prevYear
	m := 11
	for i = 0; i < 15; i++ {
		cm := m
		isNextLeap := false
		if y == currentYear && i == currentYearLeap {
			cm = -cm
		} else if y == prevYear && i == prevYearLeap {
			cm = -cm
		}
		if y == currentYear && i+1 == currentYearLeap {
			isNextLeap = true
		} else if y == prevYear && i+1 == prevYearLeap {
			isNextLeap = true
		}
		lunarYear.months.PushBack(NewLunarMonth(y, cm, dayCounts[i], hs[i]+J2000))
		if !isNextLeap {
			m++
		}
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
	return fmt.Sprintf("%d", lunarYear.year)
}

func (lunarYear *LunarYear) ToFullString() string {
	return fmt.Sprintf("%d年", lunarYear.year)
}

func (lunarYear *LunarYear) GetZhiShui() string {
	offset := 4 - NewSolarFromJulianDay(lunarYear.GetMonth(1).GetFirstJulianDay()).GetLunar().GetDayZhiIndex()
	if offset < 0 {
		offset += 12
	}
	return LunarUtil.NUMBER[offset+1] + "龙治水"
}

func (lunarYear *LunarYear) GetFenBing() string {
	offset := 2 - NewSolarFromJulianDay(lunarYear.GetMonth(1).GetFirstJulianDay()).GetLunar().GetDayGanIndex()
	if offset < 0 {
		offset += 10
	}
	return LunarUtil.NUMBER[offset+1] + "人分饼"
}

func (lunarYear *LunarYear) GetGengTian() string {
	offset := 1 - NewSolarFromJulianDay(lunarYear.GetMonth(1).GetFirstJulianDay()).GetLunar().GetDayZhiIndex()
	if offset < 0 {
		offset += 12
	}
	return LunarUtil.NUMBER[offset+1] + "牛耕田"
}

func (lunarYear *LunarYear) GetDeJin() string {
	offset := 7 - NewSolarFromJulianDay(lunarYear.GetMonth(1).GetFirstJulianDay()).GetLunar().GetDayGanIndex()
	if offset < 0 {
		offset += 10
	}
	return LunarUtil.NUMBER[offset+1] + "日得金"
}
