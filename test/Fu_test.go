package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestFu1(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2011, 7, 14)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	excepted := "初伏"
	got := fu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "初伏第1天"
	got = fu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestFu2(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2011, 7, 23)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	excepted := "初伏"
	got := fu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "初伏第10天"
	got = fu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestFu3(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2011, 7, 24)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	excepted := "中伏"
	got := fu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "中伏第1天"
	got = fu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestFu4(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2011, 8, 12)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	excepted := "中伏"
	got := fu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "中伏第20天"
	got = fu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestFu5(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2011, 8, 13)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	excepted := "末伏"
	got := fu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "末伏第1天"
	got = fu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestFu6(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2011, 8, 22)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	excepted := "末伏"
	got := fu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "末伏第10天"
	got = fu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestFu7(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2011, 7, 13)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	if fu != nil {
		t.Errorf("excepted: %v, got: %v", nil, fu)
	}
}

func TestFu8(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2011, 8, 23)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	if fu != nil {
		t.Errorf("excepted: %v, got: %v", nil, fu)
	}
}

func TestFu9(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2012, 7, 18)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	excepted := "初伏"
	got := fu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "初伏第1天"
	got = fu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestFu10(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2012, 8, 5)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	excepted := "中伏"
	got := fu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "中伏第9天"
	got = fu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestFu11(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2012, 8, 8)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	excepted := "末伏"
	got := fu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "末伏第2天"
	got = fu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestFu12(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 7, 17)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	excepted := "初伏"
	got := fu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "初伏第2天"
	got = fu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestFu13(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 7, 26)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	excepted := "中伏"
	got := fu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "中伏第1天"
	got = fu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestFu14(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 8, 24)
	lunar := solar.GetLunar()
	fu := lunar.GetFu()
	excepted := "末伏"
	got := fu.ToString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "末伏第10天"
	got = fu.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
