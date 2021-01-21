package test

import (
	"github.com/6tail/lunar-go/calendar"
	"testing"
)

func TestLunar1(t *testing.T) {
	lunar := calendar.NewLunar(2019, 3, 27, 0, 0, 0)
	excepted := "二〇一九年三月廿七"
	got := lunar.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "二〇一九年三月廿七 己亥(猪)年 戊辰(龙)月 戊戌(狗)日 子(鼠)时 纳音[平地木 大林木 平地木 桑柘木] 星期三 (七殿泰山王诞) 西方白虎 星宿[参水猿](吉) 彭祖百忌[戊不受田田主不祥 戌不吃犬作怪上床] 喜神方位[巽](东南) 阳贵神方位[艮](东北) 阴贵神方位[坤](西南) 福神方位[坎](正北) 财神方位[坎](正北) 冲[(壬辰)龙] 煞[北]"
	got = lunar.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	solar := lunar.GetSolar()
	excepted = "2019-05-01"
	got = solar.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "2019-05-01 00:00:00 星期三 (劳动节) 金牛座"
	got = solar.ToFullString()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLunar2(t *testing.T) {
	solar := calendar.NewSolar(2020, 3, 1, 12, 0, 0)
	lunar := solar.GetLunar()
	excepted := "二〇二〇年二月初八"
	got := lunar.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestLunar3(t *testing.T) {
	solar := calendar.NewSolar(2020, 1, 10, 12, 0, 0)
	lunar := solar.GetLunar()
	for i := -500; i <= 500; i++ {
		solarDay := solar.Next(i)
		excepted := solarDay.GetLunar().ToFullString()
		got := lunar.Next(i).ToFullString()
		if excepted != got {
			t.Errorf("%v excepted: %v, got: %v", solarDay.String(), excepted, got)
		}
	}
}
