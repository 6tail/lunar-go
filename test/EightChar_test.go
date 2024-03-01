package test

import (
	"github.com/6tail/lunar-go/calendar"
	"strings"
	"testing"
)

func TestEightChar1(t *testing.T) {
	lunar := calendar.NewLunar(2019, 12, 12, 11, 22, 0)
	eightChar := lunar.GetEightChar()

	excepted := "己亥"
	got := eightChar.GetYear()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "丁丑"
	got = eightChar.GetMonth()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "戊申"
	got = eightChar.GetDay()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "戊午"
	got = eightChar.GetTime()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShenGong(t *testing.T) {
	lunar := calendar.NewSolar(1995, 12, 18, 10, 28, 0).GetLunar()
	eightChar := lunar.GetEightChar()

	excepted := "壬午"
	got := eightChar.GetShenGong()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShenGong1(t *testing.T) {
	lunar := calendar.NewSolar(1994, 12, 6, 2, 0, 0).GetLunar()
	eightChar := lunar.GetEightChar()

	excepted := "乙丑"
	got := eightChar.GetShenGong()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShenGong2(t *testing.T) {
	lunar := calendar.NewSolar(1990, 12, 11, 6, 0, 0).GetLunar()
	eightChar := lunar.GetEightChar()

	excepted := "庚辰"
	got := eightChar.GetShenGong()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestShenGong3(t *testing.T) {
	lunar := calendar.NewSolar(1993, 5, 23, 4, 0, 0).GetLunar()
	eightChar := lunar.GetEightChar()

	excepted := "庚申"
	got := eightChar.GetShenGong()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestEightChar10(t *testing.T) {
	lunar := calendar.NewSolar(1988, 2, 15, 23, 30, 0).GetLunar()
	eightChar := lunar.GetEightChar()

	excepted := "戊辰"
	got := eightChar.GetYear()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "甲寅"
	got = eightChar.GetMonth()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "庚子"
	got = eightChar.GetDay()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "戊子"
	got = eightChar.GetTime()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestEightChar11(t *testing.T) {
	lunar := calendar.NewLunar(1987, 12, 28, 23, 30, 0)
	eightChar := lunar.GetEightChar()

	excepted := "戊辰"
	got := eightChar.GetYear()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "甲寅"
	got = eightChar.GetMonth()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "庚子"
	got = eightChar.GetDay()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "戊子"
	got = eightChar.GetTime()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestEightChar12(t *testing.T) {
	solarList := calendar.ListSolarFromBaZi("丙辰", "丁酉", "丙子", "甲午")

	excepted := []string{"1916-10-06 12:00:00", "1976-09-21 12:00:00"}

	got := []string{"", ""}
	solar := solarList.Front()
	got[0] = solar.Value.(*calendar.Solar).ToYmdHms()

	solar = solar.Next()
	got[1] = solar.Value.(*calendar.Solar).ToYmdHms()

	if strings.Join(excepted, ",") != strings.Join(got, ",") {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

}

func TestEightChar13(t *testing.T) {
	solarList := calendar.ListSolarFromBaZi("壬寅", "庚戌", "己未", "乙亥")

	excepted := "2022-11-02 22:00:00"

	solar := solarList.Front()
	got := solar.Value.(*calendar.Solar).ToYmdHms()

	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

}

func TestEightChar14(t *testing.T) {
	solarList := calendar.ListSolarFromBaZi("己卯", "辛未", "甲戌", "壬申")

	excepted := []string{"1939-08-05 16:00:00", "1999-07-21 16:00:00"}

	got := []string{"", ""}
	solar := solarList.Front()
	got[0] = solar.Value.(*calendar.Solar).ToYmdHms()

	solar = solar.Next()
	got[1] = solar.Value.(*calendar.Solar).ToYmdHms()

	if strings.Join(excepted, ",") != strings.Join(got, ",") {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

}

func TestEightChar15(t *testing.T) {
	solarList := calendar.ListSolarFromBaZi("庚子", "戊子", "己卯", "庚午")

	excepted := []string{"1901-01-01 12:00:00", "1960-12-17 12:00:00"}

	got := []string{"", ""}
	solar := solarList.Front()
	got[0] = solar.Value.(*calendar.Solar).ToYmdHms()

	solar = solar.Next()
	got[1] = solar.Value.(*calendar.Solar).ToYmdHms()

	if strings.Join(excepted, ",") != strings.Join(got, ",") {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

}

func TestEightChar16(t *testing.T) {
	solarList := calendar.ListSolarFromBaZi("庚子", "癸未", "乙丑", "丁亥")

	excepted := []string{"1960-08-05 22:00:00", "2020-07-21 22:00:00"}

	got := []string{"", ""}
	solar := solarList.Front()
	got[0] = solar.Value.(*calendar.Solar).ToYmdHms()

	solar = solar.Next()
	got[1] = solar.Value.(*calendar.Solar).ToYmdHms()

	if strings.Join(excepted, ",") != strings.Join(got, ",") {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

}

func TestEightChar17(t *testing.T) {
	solarList := calendar.ListSolarFromBaZiBySectAndBaseYear("癸卯", "甲寅", "癸丑", "甲子", 2, 1843)

	excepted := []string{"1843-02-08 23:00:00", "2023-02-24 23:00:00"}

	got := []string{"", ""}
	solar := solarList.Front()
	got[0] = solar.Value.(*calendar.Solar).ToYmdHms()

	solar = solar.Next()
	got[1] = solar.Value.(*calendar.Solar).ToYmdHms()

	if strings.Join(excepted, ",") != strings.Join(got, ",") {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

}

func TestEightChar18(t *testing.T) {
	solarList := calendar.ListSolarFromBaZi("己亥", "丁丑", "壬寅", "戊申")

	excepted := []string{"1900-01-29 16:00:00", "1960-01-15 16:00:00"}

	got := []string{"", ""}
	solar := solarList.Front()
	got[0] = solar.Value.(*calendar.Solar).ToYmdHms()

	solar = solar.Next()
	got[1] = solar.Value.(*calendar.Solar).ToYmdHms()

	if strings.Join(excepted, ",") != strings.Join(got, ",") {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

}

func TestEightChar19(t *testing.T) {
	solarList := calendar.ListSolarFromBaZi("己亥", "丙子", "癸酉", "庚申")

	excepted := []string{"1959-12-17 16:00:00"}

	got := []string{""}
	solar := solarList.Front()
	got[0] = solar.Value.(*calendar.Solar).ToYmdHms()

	if strings.Join(excepted, ",") != strings.Join(got, ",") {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

}

func TestEightChar20(t *testing.T) {
	lunar := calendar.NewSolar(1990, 1, 27, 0, 30, 0).GetLunar()
	eightChar := lunar.GetEightChar()

	excepted := "丙寅"
	got := eightChar.GetShenGong()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

}
