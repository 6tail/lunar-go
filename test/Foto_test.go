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

func TestFoto2(t *testing.T) {
	foto := calendar.NewFotoFromLunar(calendar.NewLunarFromYmd(2020, 4, 13))
	excepted := "氐"
	got := foto.GetXiu()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "土"
	got = foto.GetZheng()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "貉"
	got = foto.GetAnimal()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "东"
	got = foto.GetGong()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "青龙"
	got = foto.GetShou()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
