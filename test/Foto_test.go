package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestFoto1(t *testing.T) {
	foto := calendar.NewFotoFromLunar(calendar.NewLunarFromYmd(2021, 10, 14))
	excepted := "二五六五年十月十四 (三元降) (四天王巡行)"
	got := foto.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
