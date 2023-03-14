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

func TestYear6(t *testing.T) {
	year := calendar.NewLunarYear(1864)
	excepted := "上元"
	got := year.GetYuan()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestYear7(t *testing.T) {
	year := calendar.NewLunarYear(1884)
	excepted := "二运"
	got := year.GetYun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestYear8(t *testing.T) {
	year := calendar.NewLunarYear(2023)
	excepted := 384
	got := year.GetDayCount()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestYear9(t *testing.T) {
	year := calendar.NewLunarYear(2021)
	excepted := 354
	got := year.GetDayCount()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestYear10(t *testing.T) {
	year := calendar.NewLunarYear(1518)
	excepted := 355
	got := year.GetDayCount()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
