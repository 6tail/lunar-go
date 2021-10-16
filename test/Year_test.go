package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestYear1(t *testing.T) {
	year := calendar.NewLunarYear(2017)
	excepted := "二龙治水"
	got := year.GetZhiShui()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "二人分饼"
	got = year.GetFenBing()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestYear2(t *testing.T) {
	year := calendar.NewLunarYear(2018)
	excepted := "二龙治水"
	got := year.GetZhiShui()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "八人分饼"
	got = year.GetFenBing()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestYear3(t *testing.T) {
	year := calendar.NewLunarYear(5)
	excepted := "三龙治水"
	got := year.GetZhiShui()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "一人分饼"
	got = year.GetFenBing()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestYear4(t *testing.T) {
	year := calendar.NewLunarYear(2021)
	excepted := "十一牛耕田"
	got := year.GetGengTian()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestYear5(t *testing.T) {
	year := calendar.NewLunarYear(2018)
	excepted := "三日得金"
	got := year.GetDeJin()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
