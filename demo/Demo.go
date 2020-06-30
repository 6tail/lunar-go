package main

import (
	"fmt"
	"lunar-go/HolidayUtil"
	"lunar-go/calendar"
	"time"
)

func main() {
	// 法定假日
	holiday := HolidayUtil.GetHoliday("2020-01-01")
	fmt.Println(holiday)

	holiday = HolidayUtil.GetHolidayByYmd(2020, 5, 1)
	fmt.Println(holiday)

	// 阳历
	solar := calendar.NewSolarFromDate(time.Now())
	fmt.Println(solar.ToFullString())

	// 阳历往后推一天
	solar = solar.Next(1)
	fmt.Println(solar.ToFullString())

	// 法定假日
	holidays := HolidayUtil.GetHolidaysByTargetYmd(2019, 1, 1)
	for i := holidays.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value.(*calendar.Holiday))
	}

	// 阴历
	lunar := calendar.NewLunarFromYmd(1986, 4, 21)
	fmt.Println(lunar.ToFullString())
	fmt.Println(lunar.GetSolar().ToFullString())

	// 日宜
	lunar = calendar.NewLunarFromDate(time.Now())
	for i := lunar.GetDayYi().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 日忌
	for i := lunar.GetDayJi().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 日吉神宜趋
	for i := lunar.GetDayJiShen().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 日凶煞宜忌
	for i := lunar.GetDayXiongSha().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 时辰宜
	for i := lunar.GetTimeYi().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 时辰忌
	for i := lunar.GetTimeJi().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()
}
