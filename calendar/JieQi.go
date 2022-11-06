package calendar

import (
	"strings"
)

// JieQi 节气
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

// GetName 获取名称
func (jieQi *JieQi) GetName() string {
	return jieQi.name
}

// SetName 设置名称
func (jieQi *JieQi) SetName(name string) {
	jieQi.name = name
	for i, v := range JIE_QI {
		if strings.Compare(v, name) == 0 {
			if i%2 == 0 {
				jieQi.qi = true
			} else {
				jieQi.jie = true
			}
			return
		}
	}
}

// GetSolar 获取阳历日期
func (jieQi *JieQi) GetSolar() *Solar {
	return jieQi.solar
}

// SetSolar 设置阳历日期
func (jieQi *JieQi) SetSolar(solar *Solar) {
	jieQi.solar = solar
}

// IsJie 是否节令
func (jieQi *JieQi) IsJie() bool {
	return jieQi.jie
}

// IsQi 是否气令
func (jieQi *JieQi) IsQi() bool {
	return jieQi.qi
}

func (jieQi *JieQi) String() string {
	return jieQi.name
}
