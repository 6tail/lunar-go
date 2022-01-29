package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestSolarWeek1(t *testing.T) {
	week := calendar.NewSolarWeekFromYmd(2022, 5, 1, 0)

	excepted := 1
	got := week.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek2(t *testing.T) {
	week := calendar.NewSolarWeekFromYmd(2022, 5, 7, 0)

	excepted := 1
	got := week.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek3(t *testing.T) {
	week := calendar.NewSolarWeekFromYmd(2022, 5, 8, 0)

	excepted := 2
	got := week.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek4(t *testing.T) {
	week := calendar.NewSolarWeekFromYmd(2022, 5, 1, 1)

	excepted := 1
	got := week.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek5(t *testing.T) {
	week := calendar.NewSolarWeekFromYmd(2022, 5, 2, 1)

	excepted := 2
	got := week.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek6(t *testing.T) {
	week := calendar.NewSolarWeekFromYmd(2022, 5, 8, 1)

	excepted := 2
	got := week.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek7(t *testing.T) {
	week := calendar.NewSolarWeekFromYmd(2021, 11, 1, 0)

	excepted := 1
	got := week.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek8(t *testing.T) {
	week := calendar.NewSolarWeekFromYmd(2021, 11, 1, 1)

	excepted := 1
	got := week.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek9(t *testing.T) {
	week := calendar.NewSolarWeekFromYmd(2021, 5, 2, 2)

	excepted := 1
	got := week.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek10(t *testing.T) {
	week := calendar.NewSolarWeekFromYmd(2021, 5, 4, 2)

	excepted := 2
	got := week.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
