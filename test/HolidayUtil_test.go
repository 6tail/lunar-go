package test

import (
	"github.com/6tail/lunar-go/HolidayUtil"
	"testing"
)

func TestHolidayUtil1(t *testing.T) {
	holiday := HolidayUtil.GetHoliday("2016-10-04")
	excepted := "2016-10-01"
	got := holiday.GetTarget()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestHolidayUtil2(t *testing.T) {
	holiday := HolidayUtil.GetHoliday("2010-01-01")
	excepted := "元旦节"
	got := holiday.GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	HolidayUtil.Fix(nil, "20100101~000000000000000000000000000")
	holiday = HolidayUtil.GetHoliday("2010-01-01")
	if holiday != nil {
		t.Errorf("excepted: nil, got: %v", holiday)
	}
}

func TestHolidayUtil3(t *testing.T) {
	holiday := HolidayUtil.GetHoliday("2022-12-31")
	excepted := "2023-01-01"
	got := holiday.GetTarget()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
