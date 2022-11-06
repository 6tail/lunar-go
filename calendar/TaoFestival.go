package calendar

import "fmt"

// TaoFestival 道历节日
type TaoFestival struct {
	// 名称
	name string
	// 备注
	remark string
}

func NewTaoFestival(name string, remark string) *TaoFestival {
	f := new(TaoFestival)
	f.name = name
	f.remark = remark
	return f
}

func (f *TaoFestival) GetName() string {
	return f.name
}

func (f *TaoFestival) GetRemark() string {
	return f.remark
}

func (f *TaoFestival) ToString() string {
	return f.name
}

func (f *TaoFestival) String() string {
	return f.ToString()
}

func (f *TaoFestival) ToFullString() string {
	s := f.name
	if len(f.remark) > 0 {
		s += fmt.Sprintf("[%s]", f.remark)
	}
	return s
}
