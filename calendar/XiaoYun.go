package calendar

import "github.com/6tail/lunar-go/LunarUtil"

// XiaoYun 小运
type XiaoYun struct {
	// 序数，0-9
	index int
	// 大运
	daYun *DaYun
	// 年
	year int
	// 年龄
	age int
	// 是否顺推
	forward bool
	lunar   *Lunar
}

func NewXiaoYun(daYun *DaYun, index int, forward bool) *XiaoYun {
	xiaoYun := new(XiaoYun)
	xiaoYun.daYun = daYun
	xiaoYun.lunar = daYun.GetLunar()
	xiaoYun.index = index
	xiaoYun.year = daYun.GetStartYear() + index
	xiaoYun.age = daYun.GetStartAge() + index
	xiaoYun.forward = forward
	return xiaoYun
}

func (xiaoYun *XiaoYun) GetIndex() int {
	return xiaoYun.index
}

func (xiaoYun *XiaoYun) GetYear() int {
	return xiaoYun.year
}

func (xiaoYun *XiaoYun) GetAge() int {
	return xiaoYun.age
}

// GetGanZhi 获取干支
func (xiaoYun *XiaoYun) GetGanZhi() string {
	offset := LunarUtil.GetJiaZiIndex(xiaoYun.lunar.GetTimeInGanZhi())
	add := xiaoYun.index + 1
	if xiaoYun.daYun.GetIndex() > 0 {
		add += xiaoYun.daYun.GetStartAge() - 1
	}
	if xiaoYun.forward {
		offset += add
	} else {
		offset -= add
	}

	size := len(LunarUtil.JIA_ZI)
	for {
		if offset >= 0 {
			break
		}
		offset += size
	}
	offset %= size
	return LunarUtil.JIA_ZI[offset]
}

// GetXun 获取所在旬
func (xiaoYun *XiaoYun) GetXun() string {
	return LunarUtil.GetXun(xiaoYun.GetGanZhi())
}

// GetXunKong 获取旬空(空亡)
func (xiaoYun *XiaoYun) GetXunKong() string {
	return LunarUtil.GetXunKong(xiaoYun.GetGanZhi())
}
