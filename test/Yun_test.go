package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestYun1(t *testing.T) {
	solar := calendar.NewSolar(1981, 1, 29, 23, 37, 0)
	lunar := solar.GetLunar()
	eightChar := lunar.GetEightChar()
	yun := eightChar.GetYun(0)
	excepted := 8
	got := yun.GetStartYear()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = 0
	got = yun.GetStartMonth()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = 20
	got = yun.GetStartDay()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted2 := "1989-02-18"
	got2 := yun.GetStartSolar().ToYmd()
	if excepted2 != got2 {
		t.Errorf("excepted: %v, got: %v", excepted2, got2)
	}
}

func TestYun2(t *testing.T) {
	lunar := calendar.NewLunar(2019, 12, 12, 11, 22, 0)
	eightChar := lunar.GetEightChar()
	yun := eightChar.GetYun(1)
	excepted := 0
	got := yun.GetStartYear()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = 1
	got = yun.GetStartMonth()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = 0
	got = yun.GetStartDay()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted2 := "2020-02-06"
	got2 := yun.GetStartSolar().ToYmd()
	if excepted2 != got2 {
		t.Errorf("excepted: %v, got: %v", excepted2, got2)
	}
}

func TestYun3(t *testing.T) {
	solar := calendar.NewSolar(2020, 1, 6, 11, 22, 0)
	lunar := solar.GetLunar()
	eightChar := lunar.GetEightChar()
	yun := eightChar.GetYun(1)
	excepted := 0
	got := yun.GetStartYear()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = 1
	got = yun.GetStartMonth()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = 0
	got = yun.GetStartDay()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted2 := "2020-02-06"
	got2 := yun.GetStartSolar().ToYmd()
	if excepted2 != got2 {
		t.Errorf("excepted: %v, got: %v", excepted2, got2)
	}
}
