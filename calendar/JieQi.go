package calendar

import (
	"github.com/6tail/lunar-go/LunarUtil"
	"strings"
)

type JieQi struct {
	// 名称
	name string
	// 阳历日期
	solar *Solar
	// 是否节令
	jie bool
	// 是否气令
	qi bool
}

func NewJieQi(name string, solar *Solar) *JieQi {
	jieQi := new(JieQi)
	jieQi.SetName(name)
	jieQi.solar = solar
	return jieQi
}

// 获取名称
func (jieQi *JieQi) GetName() string {
	return jieQi.name
}

// 设置名称
func (jieQi *JieQi) SetName(name string) {
	jieQi.name = name
	j := len(LunarUtil.JIE)
	for i := 0; i < j; i++ {
		key := LunarUtil.JIE[i]
		if strings.Compare(key, name) == 0 {
			jieQi.jie = true
			return
		}
	}
	j = len(LunarUtil.QI)
	for i := 0; i < j; i++ {
		key := LunarUtil.QI[i]
		if strings.Compare(key, jieQi.name) == 0 {
			jieQi.qi = true
			return
		}
	}
}

// 获取阳历日期
func (jieQi *JieQi) GetSolar() *Solar {
	return jieQi.solar
}

// 设置阳历日期
func (jieQi *JieQi) SetSolar(solar *Solar) {
	jieQi.solar = solar
}

// 是否节令
func (jieQi *JieQi) IsJie() bool {
	return jieQi.jie
}

// 是否气令
func (jieQi *JieQi) IsQi() bool {
	return jieQi.qi
}

func (jieQi *JieQi) String() string {
	return jieQi.name
}
