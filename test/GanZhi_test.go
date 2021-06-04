package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestGanZhi1(t *testing.T) {
	solar := calendar.NewSolar(2020, 1, 1, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "己亥"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "丙子"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi2(t *testing.T) {
	solar := calendar.NewSolar(2020, 1, 6, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "己亥"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "丁丑"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi3(t *testing.T) {
	solar := calendar.NewSolar(2020, 1, 20, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "己亥"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "丁丑"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi4(t *testing.T) {
	solar := calendar.NewSolar(2020, 1, 25, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "庚子"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "己亥"
	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "丁丑"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi5(t *testing.T) {
	// 春节
	solar := calendar.NewSolar(2020, 1, 25, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "庚子"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "己亥"
	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "丁丑"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi6(t *testing.T) {
	solar := calendar.NewSolar(2020, 1, 30, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "庚子"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "己亥"
	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "丁丑"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi7(t *testing.T) {
	solar := calendar.NewSolar(2020, 2, 1, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "庚子"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "己亥"
	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "丁丑"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi8(t *testing.T) {
	solar := calendar.NewSolar(2020, 2, 4, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "庚子"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "己亥"
	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "戊寅"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "丁丑"
	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi9(t *testing.T) {
	solar := calendar.NewSolar(2020, 2, 4, 18, 22, 0)
	lunar := solar.GetLunar()
	excepted := "庚子"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "戊寅"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi10(t *testing.T) {
	solar := calendar.NewSolar(2020, 2, 5, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "庚子"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "戊寅"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi11(t *testing.T) {
	solar := calendar.NewSolar(2020, 5, 22, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "庚子"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "辛巳"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi12(t *testing.T) {
	solar := calendar.NewSolar(2020, 5, 23, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "庚子"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "辛巳"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi13(t *testing.T) {
	solar := calendar.NewSolar(2020, 5, 29, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "庚子"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "辛巳"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi14(t *testing.T) {
	solar := calendar.NewSolar(2020, 6, 1, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "庚子"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "辛巳"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi15(t *testing.T) {
	solar := calendar.NewSolar(2020, 6, 29, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "庚子"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "壬午"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi16(t *testing.T) {
	solar := calendar.NewSolar(2019, 5, 1, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "己亥"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "戊辰"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi17(t *testing.T) {
	solar := calendar.NewSolar(1986, 5, 29, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "丙寅"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "癸巳"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi18(t *testing.T) {
	solar := calendar.NewSolar(1986, 5, 1, 1, 22, 0)
	lunar := solar.GetLunar()
	excepted := "丙寅"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "壬辰"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi19(t *testing.T) {
	solar := calendar.NewSolar(1986, 5, 6, 1, 22, 0)
	lunar := solar.GetLunar()
	excepted := "丙寅"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "癸巳"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "壬辰"
	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi20(t *testing.T) {
	solar := calendar.NewSolar(1986, 5, 6, 23, 22, 0)
	lunar := solar.GetLunar()
	excepted := "丙寅"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "癸巳"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi21(t *testing.T) {
	solar := calendar.NewSolar(2019, 2, 8, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "己亥"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "丙寅"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestGanZhi22(t *testing.T) {
	solar := calendar.NewSolar(2020, 2, 4, 13, 22, 0)
	lunar := solar.GetLunar()
	excepted := "庚子"
	got := lunar.GetYearInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	got = lunar.GetYearInGanZhiByLiChun()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "己亥"
	got = lunar.GetYearInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "戊寅"
	got = lunar.GetMonthInGanZhi()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "丁丑"
	got = lunar.GetMonthInGanZhiExact()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
