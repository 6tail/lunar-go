package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestSolar1(t *testing.T) {
	solar := calendar.NewSolar(2020, 5, 24, 13, 0, 0)
	excepted := "二〇二〇年闰四月初二"
	got := solar.GetLunar().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar2(t *testing.T) {
	solar := calendar.NewSolarFromYmd(11, 1, 1)
	excepted := "一〇年腊月初八"
	got := solar.GetLunar().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar3(t *testing.T) {
	solar := calendar.NewSolarFromYmd(11, 3, 1)
	excepted := "一一年二月初八"
	got := solar.GetLunar().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar4(t *testing.T) {
	solar := calendar.NewSolarFromYmd(26, 4, 13)
	excepted := "二六年三月初八"
	got := solar.GetLunar().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
