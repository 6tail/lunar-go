package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestXingZuo1(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 3, 21)

	excepted := "白羊"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo2(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 4, 19)

	excepted := "白羊"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo3(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 4, 20)

	excepted := "金牛"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo4(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 5, 20)

	excepted := "金牛"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo5(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 5, 21)

	excepted := "双子"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo6(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 6, 21)

	excepted := "双子"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo7(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 6, 22)

	excepted := "巨蟹"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo8(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 7, 22)

	excepted := "巨蟹"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo9(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 7, 23)

	excepted := "狮子"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo10(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 8, 22)

	excepted := "狮子"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo11(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 8, 23)

	excepted := "处女"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo12(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 9, 22)

	excepted := "处女"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo13(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 9, 23)

	excepted := "天秤"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo14(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 10, 23)

	excepted := "天秤"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo15(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 10, 24)

	excepted := "天蝎"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo16(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 11, 22)

	excepted := "天蝎"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo17(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 11, 23)

	excepted := "射手"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo18(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 12, 21)

	excepted := "射手"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo19(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2020, 12, 22)

	excepted := "摩羯"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo20(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2021, 1, 19)

	excepted := "摩羯"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo21(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2021, 1, 20)

	excepted := "水瓶"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo22(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2021, 2, 18)

	excepted := "水瓶"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo23(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2021, 2, 19)

	excepted := "双鱼"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestXingZuo24(t *testing.T) {
	solar := calendar.NewSolarFromYmd(2021, 3, 20)

	excepted := "双鱼"
	got := solar.GetXingZuo()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
