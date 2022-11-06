package calendar

// FotoFestival 佛历因果犯忌
type FotoFestival struct {
	// 是日何日，如：雷斋日
	name string
	// 犯之因果，如：犯者夺纪
	result string
	// 是否每月同
	everyMonth bool
	// 备注，如：宜先一日即戒
	remark string
}

func NewFotoFestival(name string, result string, everyMonth bool, remark string) *FotoFestival {
	f := new(FotoFestival)
	f.name = name
	f.result = result
	f.everyMonth = everyMonth
	f.remark = remark
	return f
}

func (f *FotoFestival) GetName() string {
	return f.name
}

func (f *FotoFestival) GetResult() string {
	return f.result
}

func (f *FotoFestival) IsEveryMonth() bool {
	return f.everyMonth
}

func (f *FotoFestival) GetRemark() string {
	return f.remark
}

func (f *FotoFestival) ToString() string {
	return f.name
}

func (f *FotoFestival) String() string {
	return f.ToString()
}

func (f *FotoFestival) ToFullString() string {
	s := f.name
	if len(f.result) > 0 {
		s += " " + f.result
	}
	if len(f.remark) > 0 {
		s += " " + f.remark
	}
	return s
}
