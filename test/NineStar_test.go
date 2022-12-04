package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestNineStar1(t *testing.T) {
	lunar := calendar.NewSolarFromYmd(1985, 2, 19).GetLunar()
	excepted := "六"
	got := lunar.GetYearNineStar().GetNumber()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestNineStar2(t *testing.T) {
	lunar := calendar.NewLunarFromYmd(2022, 1, 1)
	excepted := "六白金开阳"
	got := lunar.GetYearNineStar().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestNineStar3(t *testing.T) {
	lunar := calendar.NewLunarFromYmd(2033, 1, 1)
	excepted := "四绿木天权"
	got := lunar.GetYearNineStar().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
