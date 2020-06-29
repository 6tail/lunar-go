package main

import (
	"fmt"
	"lunar-go/HolidayUtil"
	"lunar-go/calendar"
	"time"
)

func main() {
	holiday := HolidayUtil.GetHoliday("2020-01-01")
	fmt.Println(holiday)

	holiday = HolidayUtil.GetHolidayByYmd(2020,5,1)
	fmt.Println(holiday)

	solar := calendar.NewSolarFromDate(time.Now())
	fmt.Println(solar.ToFullString())

	solar = solar.Next(1)
	fmt.Println(solar.ToFullString())

	holidays := HolidayUtil.GetHolidaysByTargetYmd(2019,1,1)
	for i := holidays.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value.(*calendar.Holiday))
	}

	lunar := calendar.NewLunarFromYmd(1986,4,21)
	fmt.Println(lunar.ToFullString())
	fmt.Println(lunar.GetSolar().ToFullString())
}
