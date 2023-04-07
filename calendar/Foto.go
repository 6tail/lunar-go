package calendar

import (
	"container/list"
	"fmt"
	"github.com/6tail/lunar-go/FotoUtil"
	"github.com/6tail/lunar-go/LunarUtil"
	"strings"
)

var DEAD_YEAR = -543

// Foto 佛历
type Foto struct {
	// 阴历
	lunar *Lunar
}

func NewFotoFromLunar(lunar *Lunar) *Foto {
	f := new(Foto)
	f.lunar = lunar
	return f
}

func NewFoto(year int, month int, day int, hour int, minute int, second int) *Foto {
	return NewFotoFromLunar(NewLunar(year+DEAD_YEAR-1, month, day, hour, minute, second))
}

func NewFotoFromYmd(year int, month int, day int) *Foto {
	return NewFoto(year, month, day, 0, 0, 0)
}

func (f *Foto) GetLunar() *Lunar {
	return f.lunar
}

func (f *Foto) GetYear() int {
	sy := f.lunar.GetSolar().GetYear()
	y := sy - DEAD_YEAR
	if sy == f.lunar.GetYear() {
		y++
	}
	return y
}

func (f *Foto) GetMonth() int {
	return f.lunar.GetMonth()
}

func (f *Foto) GetDay() int {
	return f.lunar.GetDay()
}

func (f *Foto) GetYearInChinese() string {
	y := fmt.Sprintf("%d", f.GetYear())
	s := ""
	j := len(y)
	for i := 0; i < j; i++ {
		s += LunarUtil.NUMBER[[]rune(y[i : i+1])[0]-'0']
	}
	return s
}

func (f *Foto) GetMonthInChinese() string {
	return f.lunar.GetMonthInChinese()
}

func (f *Foto) GetDayInChinese() string {
	return f.lunar.GetDayInChinese()
}

func (f *Foto) GetFestivals() *list.List {
	l := list.New()
	m := f.GetMonth()
	if m < 0 {
		m = -m
	}
	if fs, ok := FotoUtil.FESTIVAL[fmt.Sprintf("%d-%d", m, f.GetDay())]; ok {
		for _, o := range fs {
			result := ""
			everyMonth := false
			remark := ""
			size := len(o)
			if size > 1 {
				result = o[1]
			}
			if size > 2 && strings.Compare(o[2], "true") == 0 {
				everyMonth = true
			}
			if size > 3 {
				remark = o[3]
			}
			l.PushBack(NewFotoFestival(o[0], result, everyMonth, remark))
		}
	}
	return l
}

func (f *Foto) GetOtherFestivals() *list.List {
	l := list.New()
	if fs, ok := FotoUtil.OTHER_FESTIVAL[fmt.Sprintf("%d-%d", f.GetMonth(), f.GetDay())]; ok {
		for _, v := range fs {
			l.PushBack(v)
		}
	}
	return l
}

func (f *Foto) IsMonthZhai() bool {
	m := f.GetMonth()
	return 1 == m || 5 == m || 9 == m
}

func (f *Foto) IsDayYangGong() bool {
	for i := f.GetFestivals().Front(); i != nil; i = i.Next() {
		o := i.Value.(FotoFestival)
		if strings.Compare("杨公忌", o.GetName()) == 0 {
			return true
		}
	}
	return false
}

func (f *Foto) IsDayZhaiShuoWang() bool {
	d := f.GetDay()
	return 1 == d || 15 == d
}

func (f *Foto) IsDayZhaiSix() bool {
	d := f.GetDay()
	if 8 == d || 14 == d || 15 == d || 23 == d || 29 == d || 30 == d {
		return true
	} else if 28 == d {
		m := NewLunarMonthFromYm(f.lunar.GetYear(), f.GetMonth())
		return nil != m && 30 != m.GetDayCount()
	}
	return false
}

func (f *Foto) IsDayZhaiTen() bool {
	d := f.GetDay()
	return 1 == d || 8 == d || 14 == d || 15 == d || 18 == d || 23 == d || 24 == d || 28 == d || 29 == d || 30 == d
}

func (f *Foto) IsDayZhaiGuanYin() bool {
	k := fmt.Sprintf("%d-%d", f.GetMonth(), f.GetDay())
	for _, v := range FotoUtil.DAY_ZHAI_GUAN_YIN {
		if strings.Compare(k, v) == 0 {
			return true
		}
	}
	return false
}

func (f *Foto) GetXiu() string {
	return FotoUtil.GetXiu(f.GetMonth(), f.GetDay())
}

func (f *Foto) GetXiuLuck() string {
	return LunarUtil.XIU_LUCK[f.GetXiu()]
}

func (f *Foto) GetXiuSong() string {
	return LunarUtil.XIU_SONG[f.GetXiu()]
}

func (f *Foto) GetZheng() string {
	return LunarUtil.ZHENG[f.GetXiu()]
}

func (f *Foto) GetAnimal() string {
	return LunarUtil.ANIMAL[f.GetXiu()]
}

func (f *Foto) GetGong() string {
	return LunarUtil.GONG[f.GetXiu()]
}

func (f *Foto) GetShou() string {
	return LunarUtil.SHOU[f.GetGong()]
}

func (f *Foto) ToString() string {
	return fmt.Sprintf("%s年%s月%s", f.GetYearInChinese(), f.GetMonthInChinese(), f.GetDayInChinese())
}

func (f *Foto) ToFullString() string {
	s := f.ToString()
	for i := f.GetFestivals().Front(); i != nil; i = i.Next() {
		o := i.Value.(*FotoFestival)
		s += fmt.Sprintf(" (%s)", o.ToString())
	}
	return s
}

func (f *Foto) String() string {
	return f.ToString()
}
