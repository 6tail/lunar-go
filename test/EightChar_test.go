package test

import (
	"github.com/6tail/lunar-go/calendar"
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
