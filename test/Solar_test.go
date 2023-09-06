package test

import (
	"github.com/6tail/lunar-go/SolarUtil"
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

func TestSolar5(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2022, 3, 28)
	excepted := "全国中小学生安全教育日"
	got := solar.GetFestivals().Front().Value
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar6(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2022, 1, 1)
	excepted := "2022-01-02"
	got := solar.NextDay(1).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar7(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2022, 1, 31)
	excepted := "2022-02-01"
	got := solar.NextDay(1).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar8(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2022, 1, 1)
	excepted := "2023-01-01"
	got := solar.NextDay(365).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar9(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2023, 1, 1)
	excepted := "2022-01-01"
	got := solar.NextDay(-365).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar10(t *testing.T) {
	solar := calendar.NewSolarFromYmd(1582, 10, 4)
	excepted := "1582-10-15"
	got := solar.NextDay(1).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar11(t *testing.T) {
	solar := calendar.NewSolarFromYmd(1582, 10, 15)
	excepted := "1582-10-04"
	got := solar.NextDay(-1).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar12(t *testing.T) {
	solar := calendar.NewSolarFromYmd(1582, 10, 4)
	excepted := "1582-11-01"
	got := solar.NextDay(18).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar13(t *testing.T) {
	solar := calendar.NewSolarFromYmd(1582, 11, 1)
	excepted := "1582-10-04"
	got := solar.NextDay(-18).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar14(t *testing.T) {
	solar := calendar.NewSolarFromYmd(1582, 11, 1)
	excepted := "1582-10-15"
	got := solar.NextDay(-17).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar15(t *testing.T) {
	excepted := 355
	got := SolarUtil.GetDaysBetween(1582, 1, 1, 1583, 1, 1)
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar16(t *testing.T) {
	excepted := 18
	got := SolarUtil.GetDaysBetween(1582, 10, 4, 1582, 11, 1)
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar17(t *testing.T) {
	excepted := 1
	got := SolarUtil.GetDaysBetween(1582, 10, 4, 1582, 10, 15)
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar18(t *testing.T) {
	solar := calendar.NewSolarFromYmd(1582, 10, 15)
	excepted := "1582-09-30"
	got := solar.NextDay(-5).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar19(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2023, 8, 31)
	excepted := "2023-09-30"
	got := solar.NextMonth(1).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar20(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2023, 8, 31)
	excepted := "2023-10-31"
	got := solar.NextMonth(2).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar21(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2023, 8, 31)
	excepted := "2025-08-31"
	got := solar.NextYear(2).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolar22(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2023, 8, 31)
	excepted := "2024-02-29"
	got := solar.NextMonth(6).ToYmd()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
