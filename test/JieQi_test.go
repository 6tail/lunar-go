package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestJieQi7(t *testing.T) {
	lunar := calendar.NewLunarFromYmd(2012, 9, 1)
	excepted := "2012-09-07 13:29:00"
	got := lunar.GetJieQiTable()["白露"].ToYmdHms()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestJieQi8(t *testing.T) {
	lunar := calendar.NewLunarFromYmd(2050, 12, 1)
	excepted := "2050-12-07 06:41:00"
	got := lunar.GetJieQiTable()["大雪"].ToYmdHms()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestJieQi9(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2021, 12, 21)
	lunar := solar.GetLunar()
	excepted := "冬至"
	got := lunar.GetJieQi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = ""
	got = lunar.GetJie()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "冬至"
	got = lunar.GetQi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
