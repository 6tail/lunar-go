package test

import (
	"github.com/6tail/lunar-go/SolarUtil"
	"testing"
)

func TestSolarUtil1(t *testing.T) {
	got := SolarUtil.IsLeapYear(2020)
	if !got {
		t.Errorf("excepted: %v, got: %v", true, got)
	}
}

func TestSolarUtil2(t *testing.T) {
	excepted := 29
	got := SolarUtil.GetDaysOfMonth(2020, 2)
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
