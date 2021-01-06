package calendar

import "strconv"

type ShuJiu struct {
	// 名称，如一九、二九
	name string
	// 当前数九第几天，1-9
	index int
}

func NewShuJiu(name string, index int) *ShuJiu {
	shuJiu := new(ShuJiu)
	shuJiu.name = name
	shuJiu.index = index
	return shuJiu
}

func (shuJiu *ShuJiu) GetName() string {
	return shuJiu.name
}

func (shuJiu *ShuJiu) SetName(name string) {
	shuJiu.name = name

}

func (shuJiu *ShuJiu) GetIndex() int {
	return shuJiu.index
}

func (shuJiu *ShuJiu) SetIndex(index int) {
	shuJiu.index = index

}

func (shuJiu *ShuJiu) ToString() string {
	return shuJiu.name
}

func (shuJiu *ShuJiu) ToFullString() string {
	return shuJiu.name + "第" + strconv.Itoa(shuJiu.index) + "天"
}

func (shuJiu *ShuJiu) String() string {
	return shuJiu.ToString()
}
