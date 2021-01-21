package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestLiuYao1(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 4, 23)
	lunar := solar.GetLunar()

	excepted := "佛灭"
	got := lunar.GetLiuYao()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLiuYao2(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2021, 1, 15)
	lunar := solar.GetLunar()

	excepted := "友引"
	got := lunar.GetLiuYao()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLiuYao3(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2017, 1, 5)
	lunar := solar.GetLunar()

	excepted := "先胜"
	got := lunar.GetLiuYao()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLiuYao4(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 4, 10)
	lunar := solar.GetLunar()

	excepted := "友引"
	got := lunar.GetLiuYao()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLiuYao5(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 6, 11)
	lunar := solar.GetLunar()

	excepted := "大安"
	got := lunar.GetLiuYao()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLiuYao6(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 6, 1)
	lunar := solar.GetLunar()

	excepted := "先胜"
	got := lunar.GetLiuYao()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLiuYao7(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 12, 8)
	lunar := solar.GetLunar()

	excepted := "先负"
	got := lunar.GetLiuYao()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLiuYao8(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 12, 11)
	lunar := solar.GetLunar()

	excepted := "赤口"
	got := lunar.GetLiuYao()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
