package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestLunarMonth(t *testing.T) {
	month := calendar.NewLunarMonthFromYm(2023, 1)
	excepted := 1
	got := month.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted1 := "甲寅"
	got1 := month.GetGanZhi()
	if excepted1 != got1 {
		t.Errorf("excepted: %v, got: %v", excepted1, got1)
	}
}

func TestLunarMonth1(t *testing.T) {
	month := calendar.NewLunarMonthFromYm(2023, -2)
	excepted := 3
	got := month.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted1 := "丙辰"
	got1 := month.GetGanZhi()
	if excepted1 != got1 {
		t.Errorf("excepted: %v, got: %v", excepted1, got1)
	}
}

func TestLunarMonth2(t *testing.T) {
	month := calendar.NewLunarMonthFromYm(2023, 3)
	excepted := 4
	got := month.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted1 := "丁巳"
	got1 := month.GetGanZhi()
	if excepted1 != got1 {
		t.Errorf("excepted: %v, got: %v", excepted1, got1)
	}
}

func TestLunarMonth3(t *testing.T) {
	month := calendar.NewLunarMonthFromYm(2024, 1)
	excepted := 1
	got := month.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted1 := "丙寅"
	got1 := month.GetGanZhi()
	if excepted1 != got1 {
		t.Errorf("excepted: %v, got: %v", excepted1, got1)
	}
}

func TestLunarMonth4(t *testing.T) {
	month := calendar.NewLunarMonthFromYm(2023, 12)
	excepted := 13
	got := month.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted1 := "丙寅"
	got1 := month.GetGanZhi()
	if excepted1 != got1 {
		t.Errorf("excepted: %v, got: %v", excepted1, got1)
	}
}

func TestLunarMonth5(t *testing.T) {
	month := calendar.NewLunarMonthFromYm(2022, 1)
	excepted := 1
	got := month.GetIndex()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted1 := "壬寅"
	got1 := month.GetGanZhi()
	if excepted1 != got1 {
		t.Errorf("excepted: %v, got: %v", excepted1, got1)
	}
}
