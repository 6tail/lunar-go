package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestEightChar1(t *testing.T) {
	lunar := calendar.NewLunar(2019, 12, 12, 11, 22, 0)
	eightChar := lunar.GetEightChar()

	excepted := "己亥"
	got := eightChar.GetYear()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "丁丑"
	got = eightChar.GetMonth()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "戊申"
	got = eightChar.GetDay()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "戊午"
	got = eightChar.GetTime()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShenGong(t *testing.T) {
	lunar := calendar.NewSolar(1995, 12, 18, 10, 28, 0).GetLunar()
	eightChar := lunar.GetEightChar()

	excepted := "壬午"
	got := eightChar.GetShenGong()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShenGong1(t *testing.T) {
	lunar := calendar.NewSolar(1994, 12, 6, 2, 0, 0).GetLunar()
	eightChar := lunar.GetEightChar()

	excepted := "丁丑"
	got := eightChar.GetShenGong()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShenGong2(t *testing.T) {
	lunar := calendar.NewSolar(1990, 12, 11, 6, 0, 0).GetLunar()
	eightChar := lunar.GetEightChar()

	excepted := "庚辰"
	got := eightChar.GetShenGong()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShenGong3(t *testing.T) {
	lunar := calendar.NewSolar(1993, 5, 23, 4, 0, 0).GetLunar()
	eightChar := lunar.GetEightChar()

	excepted := "庚申"
	got := eightChar.GetShenGong()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
