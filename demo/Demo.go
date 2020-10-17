package main

import (
	"fmt"
	"github.com/6tail/lunar-go/HolidayUtil"
	"github.com/6tail/lunar-go/calendar"
	"time"
)

func test() {
	// 法定假日
	holiday := HolidayUtil.GetHoliday("2020-01-01")
	fmt.Println(holiday)

	holiday = HolidayUtil.GetHolidayByYmd(2020, 5, 1)
	fmt.Println(holiday)

	// 阳历
	solar := calendar.NewSolarFromDate(time.Now())
	fmt.Println(solar.ToFullString())

	// 儒略日
	solar = calendar.NewSolarFromJulianDay(2459045.5)
	fmt.Println(solar.ToYmdHms())

	solar = calendar.NewSolarFromYmd(2020, 7, 15)
	fmt.Println(solar.GetJulianDay())

	// 节气表
	solar = calendar.NewSolarFromYmd(2022, 7, 15)
	lunar := solar.GetLunar()
	jieQi := lunar.GetJieQiTable()
	for i := lunar.GetJieQiList().Front(); i != nil; i = i.Next() {
		name := i.Value.(string)
		fmt.Println(name, jieQi[name].ToYmdHms())
	}

	// 下个节气
	fmt.Println(lunar.GetNextJieQi())
	// 下个节令
	fmt.Println(lunar.GetNextJie())
	// 下个气令
	fmt.Println(lunar.GetNextQi())

	// 阳历往后推一天
	solar = solar.Next(1)
	fmt.Println(solar.ToFullString())

	// 法定假日
	holidays := HolidayUtil.GetHolidaysByTargetYmd(2019, 1, 1)
	for i := holidays.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value.(*calendar.Holiday))
	}

	// 阳历转八字
	solar = calendar.NewSolar(1988, 2, 15, 23, 30, 0)
	lunar = solar.GetLunar()
	baZi := lunar.GetEightChar()
	fmt.Println(baZi)
	fmt.Println(baZi.GetYear() + " " + baZi.GetMonth() + " " + baZi.GetDay() + " " + baZi.GetTime())

	// 阴历
	lunar = calendar.NewLunarFromYmd(1986, 4, 21)
	fmt.Println(lunar.ToFullString())
	fmt.Println(lunar.GetSolar().ToFullString())

	// 八字
	baZi = lunar.GetEightChar()
	fmt.Println(baZi)
	fmt.Println(baZi.GetYear() + " " + baZi.GetMonth() + " " + baZi.GetDay() + " " + baZi.GetTime())

	// 八字五行
	fmt.Println(baZi.GetYearWuXing() + " " + baZi.GetMonthWuXing() + " " + baZi.GetDayWuXing() + " " + baZi.GetTimeWuXing())

	// 八字天干十神
	fmt.Println("八字天干十神 = " + baZi.GetYearShiShenGan() + " " + baZi.GetMonthShiShenGan() + " " + baZi.GetDayShiShenGan() + " " + baZi.GetTimeShiShenGan())

	// 八字地支十神
	fmt.Println("八字地支十神 = " + baZi.GetYearShiShenZhi().Front().Value.(string) + " " + baZi.GetMonthShiShenZhi().Front().Value.(string) + " " + baZi.GetDayShiShenZhi().Front().Value.(string) + " " + baZi.GetTimeShiShenZhi().Front().Value.(string))

	// 八字年支十神
	fmt.Print("年支十神 = ")
	for i := baZi.GetYearShiShenZhi().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 八字月支十神
	fmt.Print("月支十神 = ")
	for i := baZi.GetMonthShiShenZhi().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 八字日支十神
	fmt.Print("日支十神 = ")
	for i := baZi.GetDayShiShenZhi().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 八字时支十神
	fmt.Print("时支十神 = ")
	for i := baZi.GetTimeShiShenZhi().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	// 指定阳历时间得到八字信息
	solar = calendar.NewSolar(2012, 12, 18, 12, 0, 0)
	lunar = solar.GetLunar()
	baZi = lunar.GetEightChar()
	jieQi = lunar.GetJieQiTable()
	for i := lunar.GetJieQiList().Front(); i != nil; i = i.Next() {
		name := i.Value.(string)
		fmt.Println(name, jieQi[name].ToYmdHms())
	}

	// 男运
	yun := baZi.GetYun(1)

	fmt.Println()
	fmt.Printf("阳历%s出生\n", solar.ToYmdHms())
	fmt.Printf("出生%d年%d个月%d天后起运\n", yun.GetStartYear(), yun.GetStartMonth(), yun.GetStartDay())
	fmt.Printf("阳历%s后起运\n", yun.GetStartSolar().ToYmd())

	fmt.Println()

	// 大运
	daYunArr := yun.GetDaYun()
	for i := 0; i < len(daYunArr); i++ {
		daYun := daYunArr[i]
		fmt.Printf("大运[%d] = %d年 %d岁 %s\n", daYun.GetIndex(), daYun.GetStartYear(), daYun.GetStartAge(), daYun.GetGanZhi())
	}

	fmt.Println()

	// 第1次大运流年
	liuNianArr := daYunArr[1].GetLiuNian()
	for i := 0; i < len(liuNianArr); i++ {
		liuNian := liuNianArr[i]
		fmt.Printf("流年[%d] = %d年 %d岁 %s\n", liuNian.GetIndex(), liuNian.GetYear(), liuNian.GetAge(), liuNian.GetGanZhi())
	}
	fmt.Println()

	// 第1次大运小运
	xiaoYunArr := daYunArr[1].GetXiaoYun()
	for i := 0; i < len(xiaoYunArr); i++ {
		xiaoYun := xiaoYunArr[i]
		fmt.Printf("小运[%d] = %d年 %d岁 %s\n", xiaoYun.GetIndex(), xiaoYun.GetYear(), xiaoYun.GetAge(), xiaoYun.GetGanZhi())
	}
	fmt.Println()

	// 第1次大运首个流年的流月
	liuYueArr := liuNianArr[0].GetLiuYue()
	for i := 0; i < len(liuYueArr); i++ {
		liuYue := liuYueArr[i]
		fmt.Printf("流月[%d] = %s月 %s\n", liuYue.GetIndex(), liuYue.GetMonthInChinese(), liuYue.GetGanZhi())
	}
	fmt.Println()

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

	// 月相
	fmt.Println(lunar.GetYueXiang())

	// 值年九星
	jx := lunar.GetYearNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 值月九星
	jx = lunar.GetMonthNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 值日九星
	jx = lunar.GetDayNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 值时九星
	jx = lunar.GetTimeNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 八字转阳历
	for i := calendar.ListSolarFromBaZi("庚子", "戊子", "己卯", "庚午").Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value.(*calendar.Solar).ToFullString())
	}
	fmt.Println()
}

func test1() {
	solar := calendar.NewSolar(1981, 12, 10, 12, 30, 0)
	lunar := solar.GetLunar()
	fmt.Println(lunar.ToFullString())

	// 值年九星
	jx := lunar.GetYearNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 值月九星
	jx = lunar.GetMonthNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 值日九星
	jx = lunar.GetDayNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 值时九星
	jx = lunar.GetTimeNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())
}

func test2() {
	lunar := calendar.NewLunar(1983, 12, 15, 12, 30, 0)
	fmt.Println(lunar.ToFullString())
	fmt.Println(lunar.GetMonthZhi())

	// 值年九星
	jx := lunar.GetYearNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 值月九星
	jx = lunar.GetMonthNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 值日九星
	jx = lunar.GetDayNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 值时九星
	jx = lunar.GetTimeNineStar()
	fmt.Println(jx)
	fmt.Println(jx.ToFullString())

	// 八字转阳历
	for i := calendar.ListSolarFromBaZi("庚子", "戊子", "己卯", "庚午").Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value.(*calendar.Solar).ToFullString())
	}
	fmt.Println()
}

func main() {
	test()
	test1()
	test2()
}
