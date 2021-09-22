package main

import (
	"fmt"
	"github.com/6tail/lunar-go/HolidayUtil"
	"github.com/6tail/lunar-go/calendar"
	"strings"
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

// 工作日推移
func NextWorkday(solar *calendar.Solar, days int) *calendar.Solar {
	c := calendar.NewExactDateFromYmdHms(solar.GetYear(), solar.GetMonth(), solar.GetDay(), solar.GetHour(), solar.GetMinute(), solar.GetSecond())
	if 0 != days {
		rest := days
		if rest < 0 {
			rest = -days
		}
		add := 1
		if days < 0 {
			add = -1
		}
		for {
			if rest <= 0 {
				break
			}
			c = c.AddDate(0, 0, add)
			work := true
			holiday := HolidayUtil.GetHolidayByYmd(c.Year(), int(c.Month()), c.Day())
			if nil == holiday {
				week := int(c.Weekday())
				if 0 == week || 6 == week {
					work = false
				}
			} else {
				work = holiday.IsWork()
			}
			if work {
				rest--
			}
		}
	}
	return calendar.NewSolar(c.Year(), int(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second())
}

func test3() {
	date := calendar.NewSolarFromYmd(2020, 1, 23)
	fmt.Print(strings.Compare("2020-01-24", date.Next(1).ToYmd()))
	// 仅工作日，跨越春节假期
	fmt.Print(strings.Compare("2020-02-03", NextWorkday(date, 1).ToYmd()))

	date = calendar.NewSolarFromYmd(2020, 2, 3)
	fmt.Print(strings.Compare("2020-01-31", date.Next(-3).ToYmd()))
	// 仅工作日，跨越春节假期
	fmt.Print(strings.Compare("2020-01-21", NextWorkday(date, -3).ToYmd()))

	date = calendar.NewSolarFromYmd(2020, 2, 9)
	fmt.Print(strings.Compare("2020-02-15", date.Next(6).ToYmd()))
	// 仅工作日，跨越周末
	fmt.Print(strings.Compare("2020-02-17", NextWorkday(date, 6).ToYmd()))

	date = calendar.NewSolarFromYmd(2020, 1, 17)
	fmt.Print(strings.Compare("2020-01-18", date.Next(1).ToYmd()))
	// 仅工作日，周日调休按上班算
	fmt.Print(strings.Compare("2020-01-19", NextWorkday(date, 1).ToYmd()))
	fmt.Println()
}

func test4() {
	fmt.Print(strings.Compare("2020-01-01 元旦节 2020-01-01", HolidayUtil.GetHoliday("2020-01-01").String()))

	// 将2020-01-01修改为春节
	HolidayUtil.Fix(nil, "202001011120200101")
	fmt.Print(strings.Compare("2020-01-01 春节 2020-01-01", HolidayUtil.GetHoliday("2020-01-01").String()))

	// 追加2099-01-01为元旦节
	HolidayUtil.Fix(nil, "209901010120990101")
	fmt.Print(strings.Compare("2099-01-01 元旦节 2099-01-01", HolidayUtil.GetHoliday("2099-01-01").String()))

	// 将2020-01-01修改为春节，并追加2099-01-01为元旦节
	HolidayUtil.Fix(nil, "202001011120200101209901010120990101")
	fmt.Print(strings.Compare("2020-01-01 春节 2020-01-01", HolidayUtil.GetHoliday("2020-01-01").String()))
	fmt.Print(strings.Compare("2099-01-01 元旦节 2099-01-01", HolidayUtil.GetHoliday("2099-01-01").String()))

	// 更改节假日名称
	names := HolidayUtil.NAMES
	names[0] = "元旦"
	names[1] = "大年初一"

	HolidayUtil.Fix(names, "")
	fmt.Print(strings.Compare("2020-01-01 大年初一 2020-01-01", HolidayUtil.GetHoliday("2020-01-01").String()))
	fmt.Print(strings.Compare("2099-01-01 元旦 2099-01-01", HolidayUtil.GetHoliday("2099-01-01").String()))

	// 追加节假日名称和数据
	names = make([]string, 12)
	for i := 0; i < len(HolidayUtil.NAMES); i++ {
		names[i] = HolidayUtil.NAMES[i]
	}
	names[9] = "我的生日"
	names[10] = "结婚纪念日"
	names[11] = "她的生日"

	HolidayUtil.Fix(names, "20210529912021052920211111:12021111120211201;120211201")
	fmt.Print(strings.Compare("2021-05-29 我的生日 2021-05-29", HolidayUtil.GetHoliday("2021-05-29").String()))
	fmt.Print(strings.Compare("2021-11-11 结婚纪念日 2021-11-11", HolidayUtil.GetHoliday("2021-11-11").String()))
	fmt.Print(strings.Compare("2021-12-01 她的生日 2021-12-01", HolidayUtil.GetHoliday("2021-12-01").String()))
	fmt.Println()
}

func test5() {
	solar := calendar.NewSolarFromYmd(2020, 11, 26)
	for i := solar.GetFestivals().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	solar = calendar.NewSolarFromYmd(2020, 6, 21)
	for i := solar.GetFestivals().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	solar = calendar.NewSolarFromYmd(2021, 5, 9)
	for i := solar.GetFestivals().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	solar = calendar.NewSolarFromYmd(1986, 11, 27)
	for i := solar.GetFestivals().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	solar = calendar.NewSolarFromYmd(1985, 6, 16)
	for i := solar.GetFestivals().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()

	solar = calendar.NewSolarFromYmd(1984, 5, 13)
	for i := solar.GetFestivals().Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(string))
		fmt.Print(" ")
	}
	fmt.Println()
}

func test6() {
	solar := calendar.NewSolar(2020, 11, 19, 0, 0, 0)
	lunar := solar.GetLunar()
	// 甲午
	fmt.Println(lunar.GetYearXun())
	// 辰巳
	fmt.Println(lunar.GetYearXunKong())
	// 午未
	fmt.Println(lunar.GetMonthXunKong())
	// 戌亥
	fmt.Println(lunar.GetDayXunKong())

	solar = calendar.NewSolar(1990, 12, 23, 8, 37, 0)
	lunar = solar.GetLunar()
	eightChar := lunar.GetEightChar()
	// 子丑
	fmt.Println(eightChar.GetDayXunKong())
}

func test7() {
	// 数九
	lunar := calendar.NewLunarFromDate(time.Now())
	fmt.Println(lunar.GetShuJiu())

	// 一九第1天
	solar := calendar.NewSolarFromYmd(2020, 12, 21)
	lunar = solar.GetLunar()
	fmt.Println(lunar.GetShuJiu().ToFullString())
}

func test8() {
	// 三伏
	// 末伏第9天
	solar := calendar.NewSolarFromYmd(2020, 8, 23)
	lunar := solar.GetLunar()
	fmt.Println(lunar.GetFu().ToFullString())

	// 中伏第20天
	solar = calendar.NewSolarFromYmd(2011, 8, 12)
	lunar = solar.GetLunar()
	fmt.Println(lunar.GetFu().ToFullString())
}

func test9() {
	lunar := calendar.NewLunarFromDate(time.Now())
	// 六曜
	fmt.Println(lunar.GetLiuYao())
	// 物候
	fmt.Println(lunar.GetWuHou())
}

func main() {
	test()
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
	test9()
}
