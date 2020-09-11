package calendar

type JieQi struct {
	// 名称
	name string
	// 阳历日期
	solar *Solar
}

func NewJieQi(name string, solar *Solar) *JieQi {
	jieQi := new(JieQi)
	jieQi.name = name
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
}

// 获取阳历日期
func (jieQi *JieQi) GetSolar() *Solar {
	return jieQi.solar
}

// 设置阳历日期
func (jieQi *JieQi) SetSolar(solar *Solar) {
	jieQi.solar = solar
}
