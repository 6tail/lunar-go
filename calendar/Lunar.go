package calendar

import (
	"container/list"
	"fmt"
	"github.com/6tail/lunar-go/LunarUtil"
	"github.com/6tail/lunar-go/SolarUtil"
	"strings"
	"time"
)

var JIE_QI = []string{"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏", "小满", "芒种", "夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬", "小雪", "大雪"}
var JIE_QI_IN_USE = []string{"DA_XUE", "冬至", "小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏", "小满", "芒种", "夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬", "小雪", "大雪", "DONG_ZHI", "XIAO_HAN", "DA_HAN", "LI_CHUN", "YU_SHUI", "JING_ZHE"}

// 阴历
type Lunar struct {
	year                 int
	month                int
	day                  int
	hour                 int
	minute               int
	second               int
	yearGanIndex         int
	yearZhiIndex         int
	yearGanIndexByLiChun int
	yearZhiIndexByLiChun int
	yearGanIndexExact    int
	yearZhiIndexExact    int
	monthGanIndex        int
	monthZhiIndex        int
	monthGanIndexExact   int
	monthZhiIndexExact   int
	dayGanIndex          int
	dayZhiIndex          int
	dayGanIndexExact     int
	dayZhiIndexExact     int
	dayGanIndexExact2    int
	dayZhiIndexExact2    int
	timeGanIndex         int
	timeZhiIndex         int
	weekIndex            int
	jieQi                map[string]*Solar
	jieQiList            *list.List
	solar                *Solar
	eightChar            *EightChar
}

func NewLunar(lunarYear int, lunarMonth int, lunarDay int, hour int, minute int, second int) *Lunar {
	y := NewLunarYear(lunarYear)
	m := y.GetMonth(lunarMonth)
	if m == nil {
		panic(fmt.Sprintf("wrong lunar year %v month %v", lunarYear, lunarMonth))
	}
	if lunarDay < 1 {
		panic("lunar day must bigger than 0")
	}
	days := m.GetDayCount()
	if lunarDay > days {
		panic(fmt.Sprintf("only %v days in lunar year %v month %v", days, lunarYear, lunarMonth))
	}

	lunar := new(Lunar)
	lunar.year = lunarYear
	lunar.month = lunarMonth
	lunar.day = lunarDay
	lunar.hour = hour
	lunar.minute = minute
	lunar.second = second
	noon := NewSolarFromJulianDay(m.GetFirstJulianDay() + float64(lunarDay-1))
	lunar.solar = NewSolar(noon.GetYear(), noon.GetMonth(), noon.GetDay(), hour, minute, second)
	compute(lunar, y)
	return lunar
}

func NewLunarFromYmd(lunarYear int, lunarMonth int, lunarDay int) *Lunar {
	return NewLunar(lunarYear, lunarMonth, lunarDay, 0, 0, 0)
}

func NewLunarFromDate(date time.Time) *Lunar {
	lunarYear := 0
	lunarMonth := 0
	lunarDay := 0
	solar := NewSolarFromDate(date)
	c := NewExactDateFromYmd(solar.year, solar.month, solar.day)
	ly := NewLunarYear(solar.year)
	for i := ly.months.Front(); i != nil; i = i.Next() {
		m := i.Value.(*LunarMonth)
		day := NewSolarFromJulianDay(m.GetFirstJulianDay())
		firstDay := NewSolar(day.year, day.month, day.day, 0, 0, 0)
		days := int(c.Sub(firstDay.calendar).Hours() / 24)
		if days < m.GetDayCount() {
			lunarYear = m.GetYear()
			lunarMonth = m.GetMonth()
			lunarDay = days + 1
			break
		}
	}
	return NewLunar(lunarYear, lunarMonth, lunarDay, solar.hour, solar.minute, solar.second)
}

func computeJieQi(lunar *Lunar, lunarYear *LunarYear) {
	julianDays := lunarYear.GetJieQiJulianDays()
	size := len(JIE_QI_IN_USE)
	table := make(map[string]*Solar)
	jieQiList := list.New()
	for i := 0; i < size; i++ {
		name := JIE_QI_IN_USE[i]
		table[name] = NewSolarFromJulianDay(julianDays[i])
		jieQiList.PushBack(name)
	}
	lunar.jieQiList = jieQiList
	lunar.jieQi = table
}

func computeYear(lunar *Lunar) {
	//以正月初一开始
	offset := lunar.year - 4
	lunar.yearGanIndex = offset % 10
	lunar.yearZhiIndex = offset % 12

	if lunar.yearGanIndex < 0 {
		lunar.yearGanIndex += 10
	}

	if lunar.yearZhiIndex < 0 {
		lunar.yearZhiIndex += 12
	}

	//以立春作为新一年的开始的干支纪年
	g := lunar.yearGanIndex
	z := lunar.yearZhiIndex

	//精确的干支纪年，以立春交接时刻为准
	gExact := lunar.yearGanIndex
	zExact := lunar.yearZhiIndex

	solarYear := lunar.solar.GetYear()
	solarYmd := lunar.solar.ToYmd()
	solarYmdHms := lunar.solar.ToYmdHms()

	//获取立春的阳历时刻
	liChun := lunar.jieQi["立春"]
	if liChun.GetYear() != solarYear {
		liChun = lunar.jieQi["LI_CHUN"]
	}
	liChunYmd := liChun.ToYmd()
	liChunYmdHms := liChun.ToYmdHms()

	//阳历和阴历年份相同代表正月初一及以后
	if lunar.year == solarYear {
		//立春日期判断
		if strings.Compare(solarYmd, liChunYmd) < 0 {
			g--
			z--
		}
		//立春交接时刻判断
		if strings.Compare(solarYmdHms, liChunYmdHms) < 0 {
			gExact--
			zExact--
		}
	} else if lunar.year < solarYear {
		if strings.Compare(solarYmd, liChunYmd) >= 0 {
			g++
			z++
		}
		if strings.Compare(solarYmdHms, liChunYmdHms) >= 0 {
			gExact++
			zExact++
		}
	}

	if g < 0 {
		g += 10
	}
	if z < 0 {
		z += 12
	}
	if gExact < 0 {
		gExact += 10
	}
	if zExact < 0 {
		zExact += 12
	}

	lunar.yearGanIndexByLiChun = g % 10
	lunar.yearZhiIndexByLiChun = z % 12

	lunar.yearGanIndexExact = gExact % 10
	lunar.yearZhiIndexExact = zExact % 12
}

func computeMonth(lunar *Lunar) {
	var start *Solar
	var end *Solar
	ymd := lunar.solar.ToYmd()
	ymdhms := lunar.solar.ToYmdHms()
	size := len(JIE_QI_IN_USE)

	//序号：大雪以前-3，大雪到小寒之间-2，小寒到立春之间-1，立春之后0
	index := -3
	for i := 0; i < size; i += 2 {
		jie := JIE_QI_IN_USE[i]
		end = lunar.jieQi[jie]
		symd := ymd
		if start != nil {
			symd = start.ToYmd()
		}
		if strings.Compare(ymd, symd) >= 0 && strings.Compare(ymd, end.ToYmd()) < 0 {
			break
		}
		start = end
		index++
	}
	add := 0
	if index < 0 {
		add = 1
	}

	offset := (((lunar.yearGanIndexByLiChun+add)%5 + 1) * 2) % 10
	add = index
	if add < 0 {
		add += 10
	}
	lunar.monthGanIndex = (add + offset) % 10
	add = index
	if add < 0 {
		add += 12
	}
	lunar.monthZhiIndex = (add + LunarUtil.BASE_MONTH_ZHI_INDEX) % 12

	start = nil
	index = -3
	for i := 0; i < size; i += 2 {
		jie := JIE_QI_IN_USE[i]
		end = lunar.jieQi[jie]
		stime := ymdhms
		if start != nil {
			stime = start.ToYmdHms()
		}
		if strings.Compare(ymdhms, stime) >= 0 && strings.Compare(ymdhms, end.ToYmdHms()) < 0 {
			break
		}
		start = end
		index++
	}

	add = 0
	if index < 0 {
		add = 1
	}

	offset = (((lunar.yearGanIndexExact+add)%5 + 1) * 2) % 10
	add = index
	if add < 0 {
		add += 10
	}
	lunar.monthGanIndexExact = (add + offset) % 10
	add = index
	if add < 0 {
		add += 12
	}
	lunar.monthZhiIndexExact = (add + LunarUtil.BASE_MONTH_ZHI_INDEX) % 12
}

func computeDay(lunar *Lunar) {
	noon := NewSolar(lunar.solar.GetYear(), lunar.solar.GetMonth(), lunar.solar.GetDay(), 12, 0, 0)
	offset := int(noon.GetJulianDay() - 11)
	lunar.dayGanIndex = offset % 10
	lunar.dayZhiIndex = offset % 12

	dayGanExact := lunar.dayGanIndex
	dayZhiExact := lunar.dayZhiIndex

	lunar.dayGanIndexExact2 = dayGanExact
	lunar.dayZhiIndexExact2 = dayZhiExact

	hm := fmt.Sprintf("%02d:%02d", lunar.hour, lunar.minute)
	if strings.Compare(hm, "23:00") >= 0 && strings.Compare(hm, "23:59") <= 0 {
		dayGanExact++
		if dayGanExact >= 10 {
			dayGanExact -= 10
		}
		dayZhiExact++
		if dayZhiExact >= 12 {
			dayZhiExact -= 12
		}
	}

	lunar.dayGanIndexExact = dayGanExact
	lunar.dayZhiIndexExact = dayZhiExact
}

func computeTime(lunar *Lunar) {
	lunar.timeZhiIndex = LunarUtil.GetTimeZhiIndex(fmt.Sprintf("%02d:%02d", lunar.hour, lunar.minute))
	lunar.timeGanIndex = (lunar.dayGanIndexExact%5*2 + lunar.timeZhiIndex) % 10
}

func computeWeek(lunar *Lunar) {
	lunar.weekIndex = lunar.solar.GetWeek()
}

func compute(lunar *Lunar, lunarYear *LunarYear) {
	computeJieQi(lunar, lunarYear)
	computeYear(lunar)
	computeMonth(lunar)
	computeDay(lunar)
	computeTime(lunar)
	computeWeek(lunar)
}

// @Deprecated: 该方法已废弃，请使用GetYearGan
func (lunar *Lunar) GetGan() string {
	return lunar.GetYearGan()
}

func (lunar *Lunar) GetYearGan() string {
	return LunarUtil.GAN[lunar.yearGanIndex+1]
}

func (lunar *Lunar) GetYearGanByLiChun() string {
	return LunarUtil.GAN[lunar.yearGanIndexByLiChun+1]
}

func (lunar *Lunar) GetYearGanExact() string {
	return LunarUtil.GAN[lunar.yearGanIndexExact+1]
}

// @Deprecated: 该方法已废弃，请使用GetYearZhi
func (lunar *Lunar) GetZhi() string {
	return lunar.GetYearZhi()
}

func (lunar *Lunar) GetYearZhi() string {
	return LunarUtil.ZHI[lunar.yearZhiIndex+1]
}

func (lunar *Lunar) GetYearZhiByLiChun() string {
	return LunarUtil.ZHI[lunar.yearZhiIndexByLiChun+1]
}

func (lunar *Lunar) GetYearZhiExact() string {
	return LunarUtil.ZHI[lunar.yearZhiIndexExact+1]
}

func (lunar *Lunar) GetYearInGanZhi() string {
	return lunar.GetYearGan() + lunar.GetYearZhi()
}

func (lunar *Lunar) GetYearInGanZhiByLiChun() string {
	return lunar.GetYearGanByLiChun() + lunar.GetYearZhiByLiChun()
}

func (lunar *Lunar) GetYearInGanZhiExact() string {
	return lunar.GetYearGanExact() + lunar.GetYearZhiExact()
}

func (lunar *Lunar) GetMonthGan() string {
	return LunarUtil.GAN[lunar.monthGanIndex+1]
}

func (lunar *Lunar) GetMonthGanExact() string {
	return LunarUtil.GAN[lunar.monthGanIndexExact+1]
}

func (lunar *Lunar) GetMonthZhi() string {
	return LunarUtil.ZHI[lunar.monthZhiIndex+1]
}

func (lunar *Lunar) GetMonthZhiExact() string {
	return LunarUtil.ZHI[lunar.monthZhiIndexExact+1]
}

func (lunar *Lunar) GetMonthInGanZhi() string {
	return lunar.GetMonthGan() + lunar.GetMonthZhi()
}

func (lunar *Lunar) GetMonthInGanZhiExact() string {
	return lunar.GetMonthGanExact() + lunar.GetMonthZhiExact()
}

func (lunar *Lunar) GetDayGan() string {
	return LunarUtil.GAN[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetDayGanExact() string {
	return LunarUtil.GAN[lunar.dayGanIndexExact+1]
}

func (lunar *Lunar) GetDayGanExact2() string {
	return LunarUtil.GAN[lunar.dayGanIndexExact2+1]
}

func (lunar *Lunar) GetDayZhi() string {
	return LunarUtil.ZHI[lunar.dayZhiIndex+1]
}

func (lunar *Lunar) GetDayZhiExact() string {
	return LunarUtil.ZHI[lunar.dayZhiIndexExact+1]
}

func (lunar *Lunar) GetDayZhiExact2() string {
	return LunarUtil.ZHI[lunar.dayZhiIndexExact2+1]
}

func (lunar *Lunar) GetDayInGanZhi() string {
	return lunar.GetDayGan() + lunar.GetDayZhi()
}

func (lunar *Lunar) GetDayInGanZhiExact() string {
	return lunar.GetDayGanExact() + lunar.GetDayZhiExact()
}

func (lunar *Lunar) GetDayInGanZhiExact2() string {
	return lunar.GetDayGanExact2() + lunar.GetDayZhiExact2()
}

func (lunar *Lunar) GetTimeGan() string {
	return LunarUtil.GAN[lunar.timeGanIndex+1]
}

func (lunar *Lunar) GetTimeZhi() string {
	return LunarUtil.ZHI[lunar.timeZhiIndex+1]
}

func (lunar *Lunar) GetTimeInGanZhi() string {
	return lunar.GetTimeGan() + lunar.GetTimeZhi()
}

// @Deprecated: 该方法已废弃，请使用GetYearShengXiao
func (lunar *Lunar) GetShengxiao() string {
	return lunar.GetYearShengXiao()
}

func (lunar *Lunar) GetYearShengXiao() string {
	return LunarUtil.SHENG_XIAO[lunar.yearZhiIndex+1]
}

func (lunar *Lunar) GetYearShengXiaoByLiChun() string {
	return LunarUtil.SHENG_XIAO[lunar.yearZhiIndexByLiChun+1]
}

func (lunar *Lunar) GetYearShengXiaoExact() string {
	return LunarUtil.SHENG_XIAO[lunar.yearZhiIndexExact+1]
}

func (lunar *Lunar) GetMonthShengXiao() string {
	return LunarUtil.SHENG_XIAO[lunar.monthZhiIndex+1]
}

func (lunar *Lunar) GetDayShengXiao() string {
	return LunarUtil.SHENG_XIAO[lunar.dayZhiIndex+1]
}

func (lunar *Lunar) GetTimeShengXiao() string {
	return LunarUtil.SHENG_XIAO[lunar.timeZhiIndex+1]
}

func (lunar *Lunar) GetYearInChinese() string {
	y := fmt.Sprintf("%d", lunar.year)
	s := ""
	j := len(y)
	for i := 0; i < j; i++ {
		s += LunarUtil.NUMBER[[]rune(y[i : i+1])[0]-'0']
	}
	return s
}

func (lunar *Lunar) GetMonthInChinese() string {
	s := ""
	if lunar.month < 0 {
		s += "闰"
		s += LunarUtil.MONTH[-lunar.month]
	} else {
		s += LunarUtil.MONTH[lunar.month]
	}
	return s
}

func (lunar *Lunar) GetDayInChinese() string {
	return LunarUtil.DAY[lunar.day]
}

func (lunar *Lunar) GetSeason() string {
	m := lunar.month
	if m < 0 {
		m = -m
	}
	return LunarUtil.SEASON[m]
}

func convertJieQi(name string) string {
	jq := name
	if strings.Compare("DONG_ZHI", jq) == 0 {
		jq = "冬至"
	} else if strings.Compare("DA_HAN", jq) == 0 {
		jq = "大寒"
	} else if strings.Compare("XIAO_HAN", jq) == 0 {
		jq = "小寒"
	} else if strings.Compare("LI_CHUN", jq) == 0 {
		jq = "立春"
	} else if strings.Compare("DA_XUE", jq) == 0 {
		jq = "大雪"
	} else if strings.Compare("YU_SHUI", jq) == 0 {
		jq = "雨水"
	} else if strings.Compare("JING_ZHE", jq) == 0 {
		jq = "惊蛰"
	}
	return jq
}

func (lunar *Lunar) GetJie() string {
	jie := ""
	j := len(JIE_QI_IN_USE)
	for i := 0; i < j; i += 2 {
		key := JIE_QI_IN_USE[i]
		d := lunar.jieQi[key]
		if d.year == lunar.solar.year && d.month == lunar.solar.month && d.day == lunar.solar.day {
			jie = key
			break
		}
	}
	return convertJieQi(jie)
}

func (lunar *Lunar) GetQi() string {
	qi := ""
	j := len(JIE_QI_IN_USE)
	for i := 1; i < j; i += 2 {
		key := JIE_QI_IN_USE[i]
		d := lunar.jieQi[key]
		if d.year == lunar.solar.year && d.month == lunar.solar.month && d.day == lunar.solar.day {
			qi = key
			break
		}
	}
	return convertJieQi(qi)
}

func (lunar *Lunar) GetWeek() int {
	return lunar.weekIndex
}

func (lunar *Lunar) GetWeekInChinese() string {
	return SolarUtil.WEEK[lunar.GetWeek()]
}

func (lunar *Lunar) GetXiu() string {
	return LunarUtil.XIU[fmt.Sprintf("%s%d", lunar.GetDayZhi(), lunar.GetWeek())]
}

func (lunar *Lunar) GetXiuLuck() string {
	return LunarUtil.XIU_LUCK[lunar.GetXiu()]
}

func (lunar *Lunar) GetXiuSong() string {
	return LunarUtil.XIU_SONG[lunar.GetXiu()]
}

func (lunar *Lunar) GetZheng() string {
	return LunarUtil.ZHENG[lunar.GetXiu()]
}

func (lunar *Lunar) GetAnimal() string {
	return LunarUtil.ANIMAL[lunar.GetXiu()]
}

func (lunar *Lunar) GetGong() string {
	return LunarUtil.GONG[lunar.GetXiu()]
}

func (lunar *Lunar) GetShou() string {
	return LunarUtil.SHOU[lunar.GetGong()]
}

func (lunar *Lunar) GetFestivals() *list.List {
	l := list.New()
	if f, ok := LunarUtil.FESTIVAL[fmt.Sprintf("%d-%d", lunar.month, lunar.day)]; ok {
		l.PushBack(f)
	}
	m := lunar.month
	if m < 0 {
		m = -m
	}
	if m == 12 && lunar.day >= 29 && lunar.year != lunar.Next(1).GetYear() {
		l.PushBack("除夕")
	}
	return l
}

func (lunar *Lunar) GetOtherFestivals() *list.List {
	l := list.New()
	if f, ok := LunarUtil.OTHER_FESTIVAL[fmt.Sprintf("%d-%d", lunar.month, lunar.day)]; ok {
		for _, v := range f {
			l.PushBack(v)
		}
	}
	if strings.Compare(lunar.solar.ToYmd(), lunar.jieQi["清明"].Next(-1).ToYmd()) == 0 {
		l.PushBack("寒食节")
	}
	return l
}

func (lunar *Lunar) GetPengZuGan() string {
	return LunarUtil.PENGZU_GAN[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetPengZuZhi() string {
	return LunarUtil.PENGZU_ZHI[lunar.dayZhiIndex+1]
}

// @Deprecated: 该方法已废弃，请使用GetDayPositionXi
func (lunar *Lunar) GetPositionXi() string {
	return lunar.GetDayPositionXi()
}

// @Deprecated: 该方法已废弃，请使用GetDayPositionXiDesc
func (lunar *Lunar) GetPositionXiDesc() string {
	return lunar.GetDayPositionXiDesc()
}

// @Deprecated: 该方法已废弃，请使用GetDayPositionYangGui
func (lunar *Lunar) GetPositionYangGui() string {
	return lunar.GetDayPositionYangGui()
}

// @Deprecated: 该方法已废弃，请使用GetDayPositionYangGuiDesc
func (lunar *Lunar) GetPositionYangGuiDesc() string {
	return lunar.GetDayPositionYangGuiDesc()
}

// @Deprecated: 该方法已废弃，请使用GetDayPositionYinGui
func (lunar *Lunar) GetPositionYinGui() string {
	return lunar.GetDayPositionYinGui()
}

// @Deprecated: 该方法已废弃，请使用GetDayPositionYinGuiDesc
func (lunar *Lunar) GetPositionYinGuiDesc() string {
	return lunar.GetDayPositionYinGuiDesc()
}

// @Deprecated: 该方法已废弃，请使用GetDayPositionFu
func (lunar *Lunar) GetPositionFu() string {
	return lunar.GetDayPositionFu()
}

// @Deprecated: 该方法已废弃，请使用GetDayPositionFuDesc
func (lunar *Lunar) GetPositionFuDesc() string {
	return lunar.GetDayPositionFuDesc()
}

// @Deprecated: 该方法已废弃，请使用GetDayPositionCai
func (lunar *Lunar) GetPositionCai() string {
	return lunar.GetDayPositionCai()
}

// @Deprecated: 该方法已废弃，请使用GetDayPositionCaiDesc
func (lunar *Lunar) GetPositionCaiDesc() string {
	return lunar.GetDayPositionCaiDesc()
}

func (lunar *Lunar) GetDayPositionXi() string {
	return LunarUtil.POSITION_XI[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetDayPositionXiDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetDayPositionXi()]
}

func (lunar *Lunar) GetDayPositionYangGui() string {
	return LunarUtil.POSITION_YANG_GUI[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetDayPositionYangGuiDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetDayPositionYangGui()]
}

func (lunar *Lunar) GetDayPositionYinGui() string {
	return LunarUtil.POSITION_YIN_GUI[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetDayPositionYinGuiDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetDayPositionYinGui()]
}

func (lunar *Lunar) GetDayPositionFu() string {
	return lunar.GetDayPositionFuBySect(2)
}

func (lunar *Lunar) GetDayPositionFuBySect(sect int) string {
	offset := lunar.dayGanIndex + 1
	if 1 == sect {
		return LunarUtil.POSITION_FU[offset]
	}
	return LunarUtil.POSITION_FU_2[offset]
}

func (lunar *Lunar) GetDayPositionFuDesc() string {
	return lunar.GetDayPositionFuDescBySect(2)
}

func (lunar *Lunar) GetDayPositionFuDescBySect(sect int) string {
	return LunarUtil.POSITION_DESC[lunar.GetDayPositionFuBySect(sect)]
}

func (lunar *Lunar) GetDayPositionCai() string {
	return LunarUtil.POSITION_CAI[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetDayPositionCaiDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetDayPositionCai()]
}

func (lunar *Lunar) GetTimePositionXi() string {
	return LunarUtil.POSITION_XI[lunar.timeGanIndex+1]
}

func (lunar *Lunar) GetTimePositionXiDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetTimePositionXi()]
}

func (lunar *Lunar) GetTimePositionYangGui() string {
	return LunarUtil.POSITION_YANG_GUI[lunar.timeGanIndex+1]
}

func (lunar *Lunar) GetTimePositionYangGuiDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetTimePositionYangGui()]
}

func (lunar *Lunar) GetTimePositionYinGui() string {
	return LunarUtil.POSITION_YIN_GUI[lunar.timeGanIndex+1]
}

func (lunar *Lunar) GetTimePositionYinGuiDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetTimePositionYinGui()]
}

func (lunar *Lunar) GetTimePositionFu() string {
	return LunarUtil.POSITION_FU[lunar.timeGanIndex+1]
}

func (lunar *Lunar) GetTimePositionFuDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetTimePositionFu()]
}

func (lunar *Lunar) GetTimePositionCai() string {
	return LunarUtil.POSITION_CAI[lunar.timeGanIndex+1]
}

func (lunar *Lunar) GetTimePositionCaiDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetTimePositionCai()]
}

// @Deprecated: 该方法已废弃，请使用GetDayChong
func (lunar *Lunar) GetChong() string {
	return lunar.GetDayChong()
}

func (lunar *Lunar) GetDayChong() string {
	return LunarUtil.CHONG[lunar.dayZhiIndex+1]
}

// @Deprecated: 该方法已废弃，请使用GetDayChongGan
func (lunar *Lunar) GetChongGan() string {
	return lunar.GetDayChongGan()
}

func (lunar *Lunar) GetDayChongGan() string {
	return LunarUtil.CHONG_GAN[lunar.dayGanIndex+1]
}

// @Deprecated: 该方法已废弃，请使用GetDayChongGanTie
func (lunar *Lunar) GetChongGanTie() string {
	return lunar.GetDayChongGanTie()
}

func (lunar *Lunar) GetDayChongGanTie() string {
	return LunarUtil.CHONG_GAN_TIE[lunar.dayGanIndex+1]
}

// @Deprecated: 该方法已废弃，请使用GetDayChongShengXiao
func (lunar *Lunar) GetChongShengXiao() string {
	return lunar.GetDayChongShengXiao()
}

func (lunar *Lunar) GetDayChongShengXiao() string {
	chong := lunar.GetDayChong()
	for i, v := range LunarUtil.ZHI {
		if v == chong {
			return LunarUtil.SHENG_XIAO[i]
		}
	}
	return ""
}

// @Deprecated: 该方法已废弃，请使用GetDayChongDesc
func (lunar *Lunar) GetChongDesc() string {
	return lunar.GetDayChongDesc()
}

func (lunar *Lunar) GetDayChongDesc() string {
	return "(" + lunar.GetDayChongGan() + lunar.GetDayChong() + ")" + lunar.GetDayChongShengXiao()
}

// @Deprecated: 该方法已废弃，请使用GetDaySha
func (lunar *Lunar) GetSha() string {
	return lunar.GetDaySha()
}

func (lunar *Lunar) GetDaySha() string {
	return LunarUtil.SHA[lunar.GetDayZhi()]
}

func (lunar *Lunar) GetYearNaYin() string {
	return LunarUtil.NAYIN[lunar.GetYearInGanZhi()]
}

func (lunar *Lunar) GetMonthNaYin() string {
	return LunarUtil.NAYIN[lunar.GetMonthInGanZhi()]
}

func (lunar *Lunar) GetDayNaYin() string {
	return LunarUtil.NAYIN[lunar.GetDayInGanZhi()]
}

func (lunar *Lunar) GetTimeNaYin() string {
	return LunarUtil.NAYIN[lunar.GetTimeInGanZhi()]
}

func (lunar *Lunar) GetEightChar() *EightChar {
	if lunar.eightChar == nil {
		lunar.eightChar = NewEightChar(lunar)
	}
	return lunar.eightChar
}

// @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZi() [4]string {
	baZi := lunar.GetEightChar()
	l := [4]string{}
	l[0] = baZi.GetYear()
	l[1] = baZi.GetMonth()
	l[2] = baZi.GetDay()
	l[3] = baZi.GetTime()
	return l
}

// @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiWuXing() [4]string {
	baZi := lunar.GetEightChar()
	l := [4]string{}
	l[0] = baZi.GetYearWuXing()
	l[1] = baZi.GetMonthWuXing()
	l[2] = baZi.GetDayWuXing()
	l[3] = baZi.GetTimeWuXing()
	return l
}

// @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiNaYin() [4]string {
	baZi := lunar.GetEightChar()
	l := [4]string{}
	l[0] = baZi.GetYearNaYin()
	l[1] = baZi.GetMonthNaYin()
	l[2] = baZi.GetDayNaYin()
	l[3] = baZi.GetTimeNaYin()
	return l
}

// @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiShiShenGan() [4]string {
	baZi := lunar.GetEightChar()
	l := [4]string{}
	l[0] = baZi.GetYearShiShenGan()
	l[1] = baZi.GetMonthShiShenGan()
	l[2] = baZi.GetDayShiShenGan()
	l[3] = baZi.GetTimeShiShenGan()
	return l
}

// @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiShiShenZhi() [4]string {
	baZi := lunar.GetEightChar()
	l := [4]string{}
	l[0] = baZi.GetYearShiShenZhi().Front().Value.(string)
	l[1] = baZi.GetMonthShiShenZhi().Front().Value.(string)
	l[2] = baZi.GetDayShiShenZhi().Front().Value.(string)
	l[3] = baZi.GetTimeShiShenZhi().Front().Value.(string)
	return l
}

// @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiShiShenYearZhi() *list.List {
	return lunar.GetEightChar().GetYearShiShenZhi()
}

// @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiShiShenMonthZhi() *list.List {
	return lunar.GetEightChar().GetMonthShiShenZhi()
}

// @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiShiShenDayZhi() *list.List {
	return lunar.GetEightChar().GetDayShiShenZhi()
}

// @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiShiShenTimeZhi() *list.List {
	return lunar.GetEightChar().GetTimeShiShenZhi()
}

func (lunar *Lunar) GetZhiXing() string {
	offset := lunar.dayZhiIndex - lunar.monthZhiIndex
	if offset < 0 {
		offset += 12
	}
	return LunarUtil.ZHI_XING[offset+1]
}

func (lunar *Lunar) GetDayTianShen() string {
	monthZhi := lunar.GetMonthZhi()
	offset := LunarUtil.ZHI_TIAN_SHEN_OFFSET[monthZhi]
	return LunarUtil.TIAN_SHEN[(lunar.dayZhiIndex+offset)%12+1]
}

func (lunar *Lunar) GetTimeTianShen() string {
	dayZhi := lunar.GetDayZhiExact()
	offset := LunarUtil.ZHI_TIAN_SHEN_OFFSET[dayZhi]
	return LunarUtil.TIAN_SHEN[(lunar.timeZhiIndex+offset)%12+1]
}

func (lunar *Lunar) GetDayTianShenType() string {
	return LunarUtil.TIAN_SHEN_TYPE[lunar.GetDayTianShen()]
}

func (lunar *Lunar) GetTimeTianShenType() string {
	return LunarUtil.TIAN_SHEN_TYPE[lunar.GetTimeTianShen()]
}

func (lunar *Lunar) GetDayTianShenLuck() string {
	return LunarUtil.TIAN_SHEN_TYPE_LUCK[lunar.GetDayTianShenType()]
}

func (lunar *Lunar) GetTimeTianShenLuck() string {
	return LunarUtil.TIAN_SHEN_TYPE_LUCK[lunar.GetTimeTianShenType()]
}

func (lunar *Lunar) GetDayPositionTai() string {
	return LunarUtil.POSITION_TAI_DAY[LunarUtil.GetJiaZiIndex(lunar.GetDayInGanZhi())]
}

func (lunar *Lunar) GetMonthPositionTai() string {
	if lunar.month < 0 {
		return ""
	}
	return LunarUtil.POSITION_TAI_MONTH[lunar.month-1]
}

func (lunar *Lunar) GetTimeChong() string {
	return LunarUtil.CHONG[lunar.timeZhiIndex+1]
}

func (lunar *Lunar) GetTimeSha() string {
	return LunarUtil.SHA[lunar.GetTimeZhi()]
}

func (lunar *Lunar) GetTimeChongGan() string {
	return LunarUtil.CHONG_GAN[lunar.timeGanIndex+1]
}

func (lunar *Lunar) GetTimeChongGanTie() string {
	return LunarUtil.CHONG_GAN_TIE[lunar.timeGanIndex+1]
}

func (lunar *Lunar) GetTimeChongShengXiao() string {
	chong := lunar.GetTimeChong()
	for i, v := range LunarUtil.ZHI {
		if v == chong {
			return LunarUtil.SHENG_XIAO[i]
		}
	}
	return ""
}

func (lunar *Lunar) GetTimeChongDesc() string {
	return "(" + lunar.GetTimeChongGan() + lunar.GetTimeChong() + ")" + lunar.GetTimeChongShengXiao()
}

func (lunar *Lunar) GetJieQiTable() map[string]*Solar {
	return lunar.jieQi
}

func (lunar *Lunar) GetJieQiList() *list.List {
	return lunar.jieQiList
}

func (lunar *Lunar) GetDayYi() *list.List {
	return LunarUtil.GetDayYi(lunar.GetMonthInGanZhiExact(), lunar.GetDayInGanZhi())
}

func (lunar *Lunar) GetDayJi() *list.List {
	return LunarUtil.GetDayJi(lunar.GetMonthInGanZhiExact(), lunar.GetDayInGanZhi())
}

func (lunar *Lunar) GetDayJiShen() *list.List {
	return LunarUtil.GetDayJiShen(lunar.GetMonth(), lunar.GetDayInGanZhi())
}

func (lunar *Lunar) GetDayXiongSha() *list.List {
	return LunarUtil.GetDayXiongSha(lunar.GetMonth(), lunar.GetDayInGanZhi())
}

func (lunar *Lunar) GetTimeYi() *list.List {
	return LunarUtil.GetTimeYi(lunar.GetDayInGanZhiExact(), lunar.GetTimeInGanZhi())
}

func (lunar *Lunar) GetTimeJi() *list.List {
	return LunarUtil.GetTimeJi(lunar.GetDayInGanZhiExact(), lunar.GetTimeInGanZhi())
}

func (lunar *Lunar) GetYueXiang() string {
	return LunarUtil.YUE_XIANG[lunar.day]
}

func (lunar *Lunar) GetYearNineStar() *NineStar {
	index := -(lunar.year - 1900) % 9
	if index < 0 {
		index += 9
	}
	return NewNineStar(index)
}

func (lunar *Lunar) GetMonthNineStar() *NineStar {
	start := 2
	yearZhi := lunar.GetYearZhi()
	if strings.Index("子午卯酉", yearZhi) > -1 {
		start = 8
	} else if strings.Index("辰戌丑未", yearZhi) > -1 {
		start = 5
	}
	// 寅月起，所以需要-2
	monthIndex := lunar.monthZhiIndex - 2
	if monthIndex < 0 {
		monthIndex += 12
	}
	index := start - monthIndex - 1
	for {
		if index >= 0 {
			break
		}
		index += 9
	}
	return NewNineStar(index)
}

func (lunar *Lunar) GetDayNineStar() *NineStar {
	//顺逆
	solarYmd := lunar.solar.ToYmd()
	yuShui := lunar.jieQi["雨水"].ToYmd()
	guYu := lunar.jieQi["谷雨"].ToYmd()
	xiaZhi := lunar.jieQi["夏至"].ToYmd()
	chuShu := lunar.jieQi["处暑"].ToYmd()
	shuangJiang := lunar.jieQi["霜降"].ToYmd()

	start := 6
	asc := false
	if strings.Compare(solarYmd, lunar.jieQi["冬至"].ToYmd()) >= 0 && strings.Compare(solarYmd, yuShui) < 0 {
		asc = true
		start = 1
	} else if strings.Compare(solarYmd, yuShui) >= 0 && strings.Compare(solarYmd, guYu) < 0 {
		asc = true
		start = 7
	} else if strings.Compare(solarYmd, guYu) >= 0 && strings.Compare(solarYmd, xiaZhi) < 0 {
		asc = true
		start = 4
	} else if strings.Compare(solarYmd, xiaZhi) >= 0 && strings.Compare(solarYmd, chuShu) < 0 {
		start = 9
	} else if strings.Compare(solarYmd, chuShu) >= 0 && strings.Compare(solarYmd, shuangJiang) < 0 {
		start = 3
	}
	ganZhiIndex := LunarUtil.GetJiaZiIndex(lunar.GetDayInGanZhi()) % 9
	index := start - ganZhiIndex - 1
	if asc {
		index = start + ganZhiIndex - 1
	}
	if index > 8 {
		index -= 9
	}
	if index < 0 {
		index += 9
	}
	return NewNineStar(index)
}

func (lunar *Lunar) GetTimeNineStar() *NineStar {
	//顺逆
	solarYmd := lunar.solar.ToYmd()
	asc := false
	if strings.Compare(solarYmd, lunar.jieQi["冬至"].ToYmd()) >= 0 && strings.Compare(solarYmd, lunar.jieQi["夏至"].ToYmd()) < 0 {
		asc = true
	}
	start := 3
	if asc {
		start = 7
	}
	dayZhi := lunar.GetDayZhi()
	if strings.Index("子午卯酉", dayZhi) > -1 {
		if asc {
			start = 1
		} else {
			start = 9
		}
	} else if strings.Index("辰戌丑未", dayZhi) > -1 {
		if asc {
			start = 4
		} else {
			start = 6
		}
	}
	index := start - lunar.timeZhiIndex - 1
	if asc {
		index = start + lunar.timeZhiIndex - 1
	}
	if index > 8 {
		index -= 9
	}
	if index < 0 {
		index += 9
	}
	return NewNineStar(index)
}

// 获取下一节（顺推的第一个节）
func (lunar *Lunar) GetNextJie() *JieQi {
	l := len(JIE_QI_IN_USE) / 2
	conditions := make([]string, l)
	for i := 0; i < l; i++ {
		conditions[i] = JIE_QI_IN_USE[i*2]
	}
	return lunar.getNearJieQi(true, conditions)
}

// 获取上一节（逆推的第一个节）
func (lunar *Lunar) GetPrevJie() *JieQi {
	l := len(JIE_QI_IN_USE) / 2
	conditions := make([]string, l)
	for i := 0; i < l; i++ {
		conditions[i] = JIE_QI_IN_USE[i*2]
	}
	return lunar.getNearJieQi(false, conditions)
}

// 获取下一气令（顺推的第一个气令）
func (lunar *Lunar) GetNextQi() *JieQi {
	l := len(JIE_QI_IN_USE) / 2
	conditions := make([]string, l)
	for i := 0; i < l; i++ {
		conditions[i] = JIE_QI_IN_USE[i*2+1]
	}
	return lunar.getNearJieQi(true, conditions)
}

// 获取上一气令（逆推的第一个气令）
func (lunar *Lunar) GetPrevQi() *JieQi {
	l := len(JIE_QI_IN_USE) / 2
	conditions := make([]string, l)
	for i := 0; i < l; i++ {
		conditions[i] = JIE_QI_IN_USE[i*2+1]
	}
	return lunar.getNearJieQi(false, conditions)
}

// 获取下一节气（顺推的第一个节气）
func (lunar *Lunar) GetNextJieQi() *JieQi {
	return lunar.getNearJieQi(true, nil)
}

// 获取上一节气（逆推的第一个节气）
func (lunar *Lunar) GetPrevJieQi() *JieQi {
	return lunar.getNearJieQi(false, nil)
}

// 获取最近的节气，如果未找到匹配的，返回null
func (lunar *Lunar) getNearJieQi(forward bool, conditions []string) *JieQi {
	var name string
	var near *Solar
	filters := map[string]bool{}
	if nil != conditions {
		for _, v := range conditions {
			filters[v] = true
		}
	}
	filter := len(filters) > 0
	today := lunar.solar.ToYmdHms()
	jieQi := lunar.GetJieQiTable()
	for i := lunar.GetJieQiList().Front(); i != nil; i = i.Next() {
		key := i.Value.(string)
		jq := convertJieQi(key)
		if filter {
			if !filters[jq] {
				continue
			}
		}
		solar := jieQi[key]
		day := solar.ToYmdHms()
		if forward {
			if strings.Compare(day, today) < 0 {
				continue
			}
			if nil == near || strings.Compare(day, near.ToYmdHms()) < 0 {
				name = jq
				near = solar
			}
		} else {
			if strings.Compare(day, today) > 0 {
				continue
			}
			if nil == near || strings.Compare(day, near.ToYmdHms()) > 0 {
				name = jq
				near = solar
			}
		}
	}
	if nil == near {
		return nil
	}
	return NewJieQi(name, near)
}

// 获取节气名称，如果无节气，返回空字符串
func (lunar *Lunar) GetJieQi() string {
	name := ""
	jieQi := lunar.GetJieQiTable()
	for i := lunar.GetJieQiList().Front(); i != nil; i = i.Next() {
		jq := i.Value.(string)
		d := jieQi[jq]
		if d.GetYear() == lunar.solar.GetYear() && d.GetMonth() == lunar.solar.GetMonth() && d.GetDay() == lunar.solar.GetDay() {
			name = jq
			break
		}
	}
	return convertJieQi(name)
}

// 获取当天节气对象，如果无节气，返回nil
func (lunar *Lunar) GetCurrentJieQi() *JieQi {
	name := lunar.GetJieQi()
	if len(name) > 0 {
		return NewJieQi(name, lunar.solar)
	}
	return nil
}

// 获取当天节令对象，如果无节令，返回nil
func (lunar *Lunar) GetCurrentJie() *JieQi {
	name := lunar.GetJie()
	if len(name) > 0 {
		return NewJieQi(name, lunar.solar)
	}
	return nil
}

// 获取当天气令对象，如果无气令，返回nil
func (lunar *Lunar) GetCurrentQi() *JieQi {
	name := lunar.GetQi()
	if len(name) > 0 {
		return NewJieQi(name, lunar.solar)
	}
	return nil
}

func (lunar *Lunar) String() string {
	return lunar.GetYearInChinese() + "年" + lunar.GetMonthInChinese() + "月" + lunar.GetDayInChinese()
}

func (lunar *Lunar) ToFullString() string {
	s := ""
	s += lunar.String()
	s += " "
	s += lunar.GetYearInGanZhi()
	s += "("
	s += lunar.GetYearShengXiao()
	s += ")年 "
	s += lunar.GetMonthInGanZhi()
	s += "("
	s += lunar.GetMonthShengXiao()
	s += ")月 "
	s += lunar.GetDayInGanZhi()
	s += "("
	s += lunar.GetDayShengXiao()
	s += ")日 "
	s += lunar.GetTimeZhi()
	s += "("
	s += lunar.GetTimeShengXiao()
	s += ")时 纳音["
	s += lunar.GetYearNaYin()
	s += " "
	s += lunar.GetMonthNaYin()
	s += " "
	s += lunar.GetDayNaYin()
	s += " "
	s += lunar.GetTimeNaYin()
	s += "] 星期"
	s += lunar.GetWeekInChinese()
	for i := lunar.GetFestivals().Front(); i != nil; i = i.Next() {
		s += " ("
		s += i.Value.(string)
		s += ")"
	}
	for i := lunar.GetOtherFestivals().Front(); i != nil; i = i.Next() {
		s += " ("
		s += i.Value.(string)
		s += ")"
	}

	jq := lunar.GetJieQi()
	if len(jq) > 0 {
		s += " ["
		s += jq
		s += "]"
	}
	s += " "
	s += lunar.GetGong()
	s += "方"
	s += lunar.GetShou()
	s += " 星宿["
	s += lunar.GetXiu()
	s += lunar.GetZheng()
	s += lunar.GetAnimal()
	s += "]("
	s += lunar.GetXiuLuck()
	s += ") 彭祖百忌["
	s += lunar.GetPengZuGan()
	s += " "
	s += lunar.GetPengZuZhi()
	s += "] 喜神方位["
	s += lunar.GetDayPositionXi()
	s += "]("
	s += lunar.GetDayPositionXiDesc()
	s += ") 阳贵神方位["
	s += lunar.GetDayPositionYangGui()
	s += "]("
	s += lunar.GetDayPositionYangGuiDesc()
	s += ") 阴贵神方位["
	s += lunar.GetDayPositionYinGui()
	s += "]("
	s += lunar.GetDayPositionYinGuiDesc()
	s += ") 福神方位["
	s += lunar.GetDayPositionFu()
	s += "]("
	s += lunar.GetDayPositionFuDesc()
	s += ") 财神方位["
	s += lunar.GetDayPositionCai()
	s += "]("
	s += lunar.GetDayPositionCaiDesc()
	s += ") 冲["
	s += lunar.GetDayChongDesc()
	s += "] 煞["
	s += lunar.GetDaySha()
	s += "]"
	return s
}

func (lunar *Lunar) GetYear() int {
	return lunar.year
}

func (lunar *Lunar) GetMonth() int {
	return lunar.month
}

func (lunar *Lunar) GetDay() int {
	return lunar.day
}

func (lunar *Lunar) GetHour() int {
	return lunar.hour
}

func (lunar *Lunar) GetMinute() int {
	return lunar.minute
}

func (lunar *Lunar) GetSecond() int {
	return lunar.second
}

func (lunar *Lunar) GetTimeGanIndex() int {
	return lunar.timeGanIndex
}

func (lunar *Lunar) GetTimeZhiIndex() int {
	return lunar.timeZhiIndex
}

func (lunar *Lunar) GetDayGanIndex() int {
	return lunar.dayGanIndex
}

func (lunar *Lunar) GetDayGanIndexExact() int {
	return lunar.dayGanIndexExact
}

func (lunar *Lunar) GetDayGanIndexExact2() int {
	return lunar.dayGanIndexExact2
}

func (lunar *Lunar) GetDayZhiIndex() int {
	return lunar.dayZhiIndex
}

func (lunar *Lunar) GetDayZhiIndexExact() int {
	return lunar.dayZhiIndexExact
}

func (lunar *Lunar) GetDayZhiIndexExact2() int {
	return lunar.dayZhiIndexExact2
}

func (lunar *Lunar) GetMonthGanIndex() int {
	return lunar.monthGanIndex
}

func (lunar *Lunar) GetMonthGanIndexExact() int {
	return lunar.monthGanIndexExact
}

func (lunar *Lunar) GetMonthZhiIndex() int {
	return lunar.monthZhiIndex
}

func (lunar *Lunar) GetMonthZhiIndexExact() int {
	return lunar.monthZhiIndexExact
}

func (lunar *Lunar) GetYearGanIndex() int {
	return lunar.yearGanIndex
}

func (lunar *Lunar) GetYearGanIndexByLiChun() int {
	return lunar.yearGanIndexByLiChun
}

func (lunar *Lunar) GetYearGanIndexExact() int {
	return lunar.yearGanIndexExact
}

func (lunar *Lunar) GetYearZhiIndex() int {
	return lunar.yearZhiIndex
}

func (lunar *Lunar) GetYearZhiIndexByLiChun() int {
	return lunar.yearZhiIndexByLiChun
}

func (lunar *Lunar) GetYearZhiIndexExact() int {
	return lunar.yearZhiIndexExact
}

func (lunar *Lunar) GetSolar() *Solar {
	return lunar.solar
}

// 获取往后推几天的农历日期，如果要往前推，则天数用负数
func (lunar *Lunar) Next(days int) *Lunar {
	return lunar.solar.Next(days).GetLunar()
}

// 获取年所在旬（以正月初一作为新年的开始）
func (lunar *Lunar) GetYearXun() string {
	return LunarUtil.GetXun(lunar.GetYearInGanZhi())
}

// 获取年所在旬（以立春当天作为新年的开始）
func (lunar *Lunar) GetYearXunByLiChun() string {
	return LunarUtil.GetXun(lunar.GetYearInGanZhiByLiChun())
}

// 获取年所在旬（以立春交接时刻作为新年的开始）
func (lunar *Lunar) GetYearXunExact() string {
	return LunarUtil.GetXun(lunar.GetYearInGanZhiExact())
}

// 获取值年空亡（以正月初一作为新年的开始）
func (lunar *Lunar) GetYearXunKong() string {
	return LunarUtil.GetXunKong(lunar.GetYearInGanZhi())
}

// 获取值年空亡（以立春当天作为新年的开始）
func (lunar *Lunar) GetYearXunKongByLiChun() string {
	return LunarUtil.GetXunKong(lunar.GetYearInGanZhiByLiChun())
}

// 获取值年空亡（以立春交接时刻作为新年的开始）
func (lunar *Lunar) GetYearXunKongExact() string {
	return LunarUtil.GetXunKong(lunar.GetYearInGanZhiExact())
}

// 获取月所在旬（以节交接当天起算）
func (lunar *Lunar) GetMonthXun() string {
	return LunarUtil.GetXun(lunar.GetMonthInGanZhi())
}

// 获取月所在旬（以节交接时刻起算）
func (lunar *Lunar) GetMonthXunExact() string {
	return LunarUtil.GetXun(lunar.GetMonthInGanZhiExact())
}

// 获取值月空亡（以节交接当天起算）
func (lunar *Lunar) GetMonthXunKong() string {
	return LunarUtil.GetXunKong(lunar.GetMonthInGanZhi())
}

// 获取值月空亡（以节交接时刻起算）
func (lunar *Lunar) GetMonthXunKongExact() string {
	return LunarUtil.GetXunKong(lunar.GetMonthInGanZhiExact())
}

// 获取日所在旬（以节交接当天起算）
func (lunar *Lunar) GetDayXun() string {
	return LunarUtil.GetXun(lunar.GetDayInGanZhi())
}

// 获取日所在旬（晚子时日柱算明天）
func (lunar *Lunar) GetDayXunExact() string {
	return LunarUtil.GetXun(lunar.GetDayInGanZhiExact())
}

// 获取日所在旬（晚子时日柱算当天）
func (lunar *Lunar) GetDayXunExact2() string {
	return LunarUtil.GetXun(lunar.GetDayInGanZhiExact2())
}

// 获取值日空亡
func (lunar *Lunar) GetDayXunKong() string {
	return LunarUtil.GetXunKong(lunar.GetDayInGanZhi())
}

// 获取值日空亡（晚子时日柱算明天）
func (lunar *Lunar) GetDayXunKongExact() string {
	return LunarUtil.GetXunKong(lunar.GetDayInGanZhiExact())
}

// 获取值日空亡（晚子时日柱算当天）
func (lunar *Lunar) GetDayXunKongExact2() string {
	return LunarUtil.GetXunKong(lunar.GetDayInGanZhiExact2())
}

// 获取时辰所在旬
func (lunar *Lunar) GetTimeXun() string {
	return LunarUtil.GetXun(lunar.GetTimeInGanZhi())
}

// 获取值时空亡
func (lunar *Lunar) GetTimeXunKong() string {
	return LunarUtil.GetXunKong(lunar.GetTimeInGanZhi())
}

// 获取数九，如果不是数九天，返回nil
func (lunar *Lunar) GetShuJiu() *ShuJiu {
	currentCalendar := NewExactDateFromYmd(lunar.solar.GetYear(), lunar.solar.GetMonth(), lunar.solar.GetDay())
	start := lunar.jieQi["DONG_ZHI"]
	startCalendar := NewExactDateFromYmd(start.GetYear(), start.GetMonth(), start.GetDay())
	if currentCalendar.Before(startCalendar) {
		start = lunar.jieQi["冬至"]
		startCalendar = NewExactDateFromYmd(start.GetYear(), start.GetMonth(), start.GetDay())
	}
	endCalendar := startCalendar.AddDate(0, 0, 81)
	if currentCalendar.Before(startCalendar) || !currentCalendar.Before(endCalendar) {
		return nil
	}
	days := int((currentCalendar.Unix() - startCalendar.Unix()) / 86400)
	return NewShuJiu(LunarUtil.NUMBER[days/9+1]+"九", days%9+1)
}

// 获取三伏，如果不是三伏天，返回nil
func (lunar *Lunar) GetFu() *Fu {
	currentCalendar := NewExactDateFromYmd(lunar.solar.GetYear(), lunar.solar.GetMonth(), lunar.solar.GetDay())
	xiaZhi := lunar.jieQi["夏至"]
	liQiu := lunar.jieQi["立秋"]
	startCalendar := NewExactDateFromYmd(xiaZhi.GetYear(), xiaZhi.GetMonth(), xiaZhi.GetDay())
	// 第1个庚日
	add := 6 - xiaZhi.GetLunar().GetDayGanIndex()
	if add < 0 {
		add += 10
	}
	// 第3个庚日，即初伏第1天
	add += 20
	startCalendar = startCalendar.AddDate(0, 0, add)

	// 初伏以前
	if currentCalendar.Before(startCalendar) {
		return nil
	}

	days := int((currentCalendar.Unix() - startCalendar.Unix()) / 86400)
	if days < 10 {
		return NewFu("初伏", days+1)
	}

	// 第4个庚日，中伏第1天
	startCalendar = startCalendar.AddDate(0, 0, 10)
	days = int((currentCalendar.Unix() - startCalendar.Unix()) / 86400)
	if days < 10 {
		return NewFu("中伏", days+1)
	}

	// 第5个庚日，中伏第11天或末伏第1天
	startCalendar = startCalendar.AddDate(0, 0, 10)
	days = int((currentCalendar.Unix() - startCalendar.Unix()) / 86400)

	liQiuCalendar := NewExactDateFromYmd(liQiu.GetYear(), liQiu.GetMonth(), liQiu.GetDay())

	// 末伏
	if !liQiuCalendar.After(startCalendar) {
		if days < 10 {
			return NewFu("末伏", days+1)
		}
	} else {
		// 中伏
		if days < 10 {
			return NewFu("中伏", days+11)
		}
		// 末伏第1天
		startCalendar = startCalendar.AddDate(0, 0, 10)
		days = int((currentCalendar.Unix() - startCalendar.Unix()) / 86400)
		if days < 10 {
			return NewFu("末伏", days+1)
		}
	}
	return nil
}

// 获取六曜
func (lunar *Lunar) GetLiuYao() string {
	month := lunar.month
	if month < 0 {
		month = 0 - month
	}
	return LunarUtil.LIU_YAO[(month+lunar.day-2)%6]
}

// 获取物候
func (lunar *Lunar) GetWuHou() string {
	jq := lunar.GetPrevJieQi()
	name := jq.GetName()
	offset := 0
	for i, v := range JIE_QI {
		if strings.Compare(name, v) == 0 {
			offset = i
			break
		}
	}
	currentCalendar := NewExactDateFromYmd(lunar.solar.GetYear(), lunar.solar.GetMonth(), lunar.solar.GetDay())
	startSolar := jq.GetSolar()
	startCalendar := NewExactDateFromYmd(startSolar.GetYear(), startSolar.GetMonth(), startSolar.GetDay())
	days := int((currentCalendar.Unix() - startCalendar.Unix()) / 86400)
	return LunarUtil.WU_HOU[(offset*3+days/5)%len(LunarUtil.WU_HOU)]
}

// 获取日禄
func (lunar *Lunar) GetDayLu() string {
	gan := LunarUtil.LU[lunar.GetDayGan()]
	lu := gan + "命互禄"
	if zhi, ok := LunarUtil.LU[lunar.GetDayZhi()]; ok {
		lu += " " + zhi + "命进禄"
	}
	return lu
}

// 获取时辰
func (lunar *Lunar) GetTime() *LunarTime {
	return NewLunarTime(lunar.year, lunar.month, lunar.day, lunar.hour, lunar.minute, lunar.second)
}

// 获取当天的时辰列表
func (lunar *Lunar) GetTimes() []*LunarTime {
	l := make([]*LunarTime, 13)
	l[0] = NewLunarTime(lunar.year, lunar.month, lunar.day, 0, 0, 0)
	for i := 0; i < 12; i++ {
		l[i+1] = NewLunarTime(lunar.year, lunar.month, lunar.day, (i+1)*2-1, 0, 0)
	}
	return l
}

// 获取佛历
func (lunar *Lunar) GetFoto() *Foto {
	return NewFotoFromLunar(lunar)
}

// 获取道历
func (lunar *Lunar) GetTao() *Tao {
	return NewTaoFromLunar(lunar)
}
