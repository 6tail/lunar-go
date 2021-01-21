package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestShuJiu1(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 12, 21)
	lunar := solar.GetLunar()
	shuJiu := lunar.GetShuJiu()
	excepted := "一九"
	got := shuJiu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "一九第1天"
	got = shuJiu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShuJiu2(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 12, 22)
	lunar := solar.GetLunar()
	shuJiu := lunar.GetShuJiu()
	excepted := "一九"
	got := shuJiu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "一九第2天"
	got = shuJiu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShuJiu3(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 1, 7)
	lunar := solar.GetLunar()
	shuJiu := lunar.GetShuJiu()
	excepted := "二九"
	got := shuJiu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "二九第8天"
	got = shuJiu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShuJiu4(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2021, 1, 6)
	lunar := solar.GetLunar()
	shuJiu := lunar.GetShuJiu()
	excepted := "二九"
	got := shuJiu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "二九第8天"
	got = shuJiu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShuJiu5(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2021, 1, 8)
	lunar := solar.GetLunar()
	shuJiu := lunar.GetShuJiu()
	excepted := "三九"
	got := shuJiu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "三九第1天"
	got = shuJiu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShuJiu6(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2021, 3, 5)
	lunar := solar.GetLunar()
	shuJiu := lunar.GetShuJiu()
	excepted := "九九"
	got := shuJiu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "九九第3天"
	got = shuJiu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShuJiu7(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2021, 7, 5)
	lunar := solar.GetLunar()
	shuJiu := lunar.GetShuJiu()
	if shuJiu != nil {
		t.Errorf("excepted: %v, got: %v", nil, shuJiu)
	}
}
