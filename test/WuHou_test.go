package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestWuHou1(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 4, 23)
	lunar := solar.GetLunar()

	excepted := "萍始生"
	got := lunar.GetWuHou()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestWuHou2(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2021, 1, 15)
	lunar := solar.GetLunar()

	excepted := "雉始雊"
	got := lunar.GetWuHou()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestWuHou3(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2017, 1, 5)
	lunar := solar.GetLunar()

	excepted := "雁北乡"
	got := lunar.GetWuHou()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestWuHou4(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 4, 10)
	lunar := solar.GetLunar()

	excepted := "田鼠化为鴽"
	got := lunar.GetWuHou()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestWuHou5(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 6, 11)
	lunar := solar.GetLunar()

	excepted := "鵙始鸣"
	got := lunar.GetWuHou()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestWuHou6(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 6, 1)
	lunar := solar.GetLunar()

	excepted := "麦秋至"
	got := lunar.GetWuHou()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestWuHou7(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 12, 8)
	lunar := solar.GetLunar()

	excepted := "鹖鴠不鸣"
	got := lunar.GetWuHou()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestWuHou8(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 12, 11)
	lunar := solar.GetLunar()

	excepted := "鹖鴠不鸣"
	got := lunar.GetWuHou()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestWuHou9(t *testing.T) {
	solar := calendar.NewSolarFromYmd(1982, 12, 22)
	lunar := solar.GetLunar()

	excepted := "蚯蚓结"
	got := lunar.GetWuHou()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
