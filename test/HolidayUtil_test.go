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
