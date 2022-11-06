package calendar

import (
	"fmt"
)

// Fu 伏
type Fu struct {
	// 名称：初伏、中伏、末伏
	name string
	// 当前入伏第几天，1-20
	index int
}

func NewFu(name string, index int) *Fu {
	fu := new(Fu)
	fu.name = name
	fu.index = index
	return fu
}

func (fu *Fu) GetName() string {
	return fu.name
}

func (fu *Fu) SetName(name string) {
	fu.name = name
}

func (fu *Fu) GetIndex() int {
	return fu.index
}

func (fu *Fu) SetIndex(index int) {
	fu.index = index
}

func (fu *Fu) ToString() string {
	return fu.name
}

func (fu *Fu) ToFullString() string {
	return fmt.Sprintf("%v第%d天", fu.name, fu.index)
}

func (fu *Fu) String() string {
	return fu.ToString()
}
