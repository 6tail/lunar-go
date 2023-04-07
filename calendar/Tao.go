package calendar

import (
	"container/list"
	"fmt"
	"github.com/6tail/lunar-go/LunarUtil"
	"github.com/6tail/lunar-go/TaoUtil"
	"strings"
)

var BIRTH_YEAR = -2697

// Tao 道历
type Tao struct {
	// 阴历
	lunar *Lunar
}

func NewTaoFromLunar(lunar *Lunar) *Tao {
	t := new(Tao)
	t.lunar = lunar
	return t
}

func NewTao(year int, month int, day int, hour int, minute int, second int) *Tao {
	return NewTaoFromLunar(NewLunar(year+BIRTH_YEAR, month, day, hour, minute, second))
}

func NewTaoFromYmd(year int, month int, day int) *Tao {
	return NewTao(year, month, day, 0, 0, 0)
}

func (t *Tao) GetLunar() *Lunar {
	return t.lunar
}

func (t *Tao) GetYear() int {
	return t.lunar.GetYear() - BIRTH_YEAR
}

func (t *Tao) GetMonth() int {
	return t.lunar.GetMonth()
}

func (t *Tao) GetDay() int {
	return t.lunar.GetDay()
}

func (t *Tao) GetYearInChinese() string {
	y := fmt.Sprintf("%d", t.GetYear())
	s := ""
	j := len(y)
	for i := 0; i < j; i++ {
		s += LunarUtil.NUMBER[[]rune(y[i : i+1])[0]-'0']
	}
	return s
}

func (t *Tao) GetMonthInChinese() string {
	return t.lunar.GetMonthInChinese()
}

func (t *Tao) GetDayInChinese() string {
	return t.lunar.GetDayInChinese()
}

func (t *Tao) GetFestivals() *list.List {
	l := list.New()
	if f, ok := TaoUtil.FESTIVAL[fmt.Sprintf("%d-%d", t.GetMonth(), t.GetDay())]; ok {
		for _, o := range f {
			remark := ""
			size := len(o)
			if size > 1 {
				remark = o[1]
			}
			l.PushBack(NewTaoFestival(o[0], remark))
		}
	}
	jq := t.lunar.GetJieQi()
	if strings.Compare("冬至", jq) == 0 {
		l.PushBack(NewTaoFestival("元始天尊圣诞", ""))
	} else if strings.Compare("夏至", jq) == 0 {
		l.PushBack(NewTaoFestival("灵宝天尊圣诞", ""))
	}
	// 八节日
	if f, ok := TaoUtil.BA_JIE[jq]; ok {
		l.PushBack(NewTaoFestival(f, ""))
	}
	// 八会日
	if f, ok := TaoUtil.BA_HUI[t.lunar.GetDayInGanZhi()]; ok {
		l.PushBack(NewTaoFestival(f, ""))
	}
	return l
}

func (t *Tao) isDayIn(days []string) bool {
	k := fmt.Sprintf("%d-%d", t.GetMonth(), t.GetDay())
	for _, v := range days {
		if strings.Compare(k, v) == 0 {
			return true
		}
	}
	return false
}

func (t *Tao) IsDaySanHui() bool {
	return t.isDayIn(TaoUtil.SAN_HUI)
}

func (t *Tao) IsDaySanYuan() bool {
	return t.isDayIn(TaoUtil.SAN_YUAN)
}

func (t *Tao) IsDayWuLa() bool {
	return t.isDayIn(TaoUtil.WU_LA)
}

func (t *Tao) IsDayBaJie() bool {
	if _, ok := TaoUtil.BA_JIE[t.lunar.GetJieQi()]; ok {
		return true
	}
	return false
}

func (t *Tao) IsDayBaHui() bool {
	if _, ok := TaoUtil.BA_HUI[t.lunar.GetDayInGanZhi()]; ok {
		return true
	}
	return false
}

func (t *Tao) IsDayMingWu() bool {
	return strings.Compare("戊", t.lunar.GetDayGan()) == 0
}

func (t *Tao) IsDayAnWu() bool {
	m := t.GetMonth()
	if m < 0 {
		m = -m
	}
	return strings.Compare(t.lunar.GetDayZhi(), TaoUtil.AN_WU[m-1]) == 0
}

func (t *Tao) IsDayWu() bool {
	return t.IsDayMingWu() || t.IsDayAnWu()
}

func (t *Tao) ToString() string {
	return fmt.Sprintf("%s年%s月%s", t.GetYearInChinese(), t.GetMonthInChinese(), t.GetDayInChinese())
}

func (t *Tao) ToFullString() string {
	return fmt.Sprintf("道歷%s年，天運%s年，%s月，%s日。%s月%s日，%s時。", t.GetYearInChinese(), t.lunar.GetYearInGanZhi(), t.lunar.GetMonthInGanZhi(), t.lunar.GetDayInGanZhi(), t.GetMonthInChinese(), t.GetDayInChinese(), t.lunar.GetTimeZhi())
}

func (t *Tao) String() string {
	return t.ToString()
}
