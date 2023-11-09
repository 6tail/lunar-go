package test

import (
	"github.com/6tail/lunar-go/SolarUtil"
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

func TestSolarWeek11(t *testing.T) {
	week := calendar.NewSolarWeekFromYmd(2022, 3, 6, 0)

	excepted := 11
	got := week.GetIndexInYear()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek12(t *testing.T) {
	month := calendar.NewSolarMonthFromYm(2022, 12)

	excepted := 5
	got := month.GetWeeks(0).Len()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek13(t *testing.T) {
	got := SolarUtil.GetWeek(1582, 10, 15)

	excepted := 5
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek14(t *testing.T) {
	got := SolarUtil.GetWeek(1582, 10, 1)

	excepted := 1
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek15(t *testing.T) {
	got := SolarUtil.GetWeek(1129, 11, 17)

	excepted := 0
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek16(t *testing.T) {
	got := SolarUtil.GetWeek(1129, 11, 1)

	excepted := 5
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek17(t *testing.T) {
	got := SolarUtil.GetWeek(8, 11, 1)

	excepted := 4
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek18(t *testing.T) {
	got := SolarUtil.GetWeek(1582, 9, 30)

	excepted := 0
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek19(t *testing.T) {
	got := SolarUtil.GetWeek(1582, 1, 1)

	excepted := 1
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek20(t *testing.T) {
	got := SolarUtil.GetWeek(1500, 2, 29)

	excepted := 6
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek21(t *testing.T) {
	got := SolarUtil.GetWeek(9865, 7, 26)

	excepted := 3
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek22(t *testing.T) {
	got := SolarUtil.GetWeek(1961, 9, 30)

	excepted := 6
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek23(t *testing.T) {
	got := calendar.NewSolar(1961, 9, 30, 23, 59, 59).GetWeek()

	excepted := 6
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestSolarWeek24(t *testing.T) {
	got := calendar.NewSolar(2021, 9, 15, 20, 0, 0).GetWeek()

	excepted := 3
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
