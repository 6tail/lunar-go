package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestTao1(t *testing.T) {
	tao := calendar.NewTaoFromLunar(calendar.NewLunar(2021, 10, 17, 18, 0, 0))
	excepted := "四七一八年十月十七"
	got := tao.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "道歷四七一八年，天運辛丑年，己亥月，癸酉日。十月十七日，酉時。"
	got = tao.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestTao2(t *testing.T) {
	tao := calendar.NewTaoFromYmd(4718, 10, 18)
	s := ""
	for i := tao.GetFestivals().Front(); i != nil; i = i.Next() {
		f := i.Value.(*calendar.TaoFestival)
		s += f.ToString()
	}
	excepted := "地母娘娘圣诞四时会"
	got := s
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	tao = calendar.NewLunarFromYmd(2021, 10, 18).GetTao()
	s = ""
	for i := tao.GetFestivals().Front(); i != nil; i = i.Next() {
		f := i.Value.(*calendar.TaoFestival)
		s += f.ToString()
	}
	excepted = "地母娘娘圣诞四时会"
	got = s
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
