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

// Lunar 阴历
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
	if noon.GetYear() != lunarYear {
		y = NewLunarYear(noon.GetYear())
	}
	compute(lunar, y)
	return lunar
}

func NewLunarFromYmd(lunarYear int, lunarMonth int, lunarDay int) *Lunar {
	return NewLunar(lunarYear, lunarMonth, lunarDay, 0, 0, 0)
}

func NewLunarFromSolar(solar *Solar) *Lunar {
	lunarYear := 0
	lunarMonth := 0
	lunarDay := 0
	ly := NewLunarYear(solar.GetYear())
	lunar := new(Lunar)
	for i := ly.months.Front(); i != nil; i = i.Next() {
		m := i.Value.(*LunarMonth)
		days := solar.Subtract(NewSolarFromJulianDay(m.GetFirstJulianDay()))
		if days < m.GetDayCount() {
			lunarYear = m.GetYear()
			lunarMonth = m.GetMonth()
			lunarDay = days + 1
			break
		}
	}
	lunar.year = lunarYear
	lunar.month = lunarMonth
	lunar.day = lunarDay
	lunar.hour = solar.GetHour()
	lunar.minute = solar.GetMinute()
	lunar.second = solar.GetSecond()
	lunar.solar = solar
	compute(lunar, ly)
	return lunar
}

func NewLunarFromDate(date time.Time) *Lunar {
	return NewLunarFromSolar(NewSolarFromDate(date))
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

// GetGan @Deprecated: 该方法已废弃，请使用GetYearGan
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

// GetZhi @Deprecated: 该方法已废弃，请使用GetYearZhi
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

// GetShengxiao @Deprecated: 该方法已废弃，请使用GetYearShengXiao
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
	solarYmd := lunar.solar.ToYmd()
	jq := lunar.jieQi["清明"]
	if strings.Compare(solarYmd, jq.NextDay(-1).ToYmd()) == 0 {
		l.PushBack("寒食节")
	}

	jq = lunar.jieQi["立春"]
	offset := 4 - jq.GetLunar().GetDayGanIndex()
	if offset < 0 {
		offset += 10
	}
	if strings.Compare(solarYmd, jq.NextDay(offset+40).ToYmd()) == 0 {
		l.PushBack("春社")
	}

	jq = lunar.jieQi["立秋"]
	offset = 4 - jq.GetLunar().GetDayGanIndex()
	if offset < 0 {
		offset += 10
	}
	if strings.Compare(solarYmd, jq.NextDay(offset+40).ToYmd()) == 0 {
		l.PushBack("秋社")
	}
	return l
}

func (lunar *Lunar) GetPengZuGan() string {
	return LunarUtil.PENGZU_GAN[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetPengZuZhi() string {
	return LunarUtil.PENGZU_ZHI[lunar.dayZhiIndex+1]
}

// GetPositionXi @Deprecated: 该方法已废弃，请使用GetDayPositionXi
func (lunar *Lunar) GetPositionXi() string {
	return lunar.GetDayPositionXi()
}

// GetPositionXiDesc @Deprecated: 该方法已废弃，请使用GetDayPositionXiDesc
func (lunar *Lunar) GetPositionXiDesc() string {
	return lunar.GetDayPositionXiDesc()
}

// GetPositionYangGui @Deprecated: 该方法已废弃，请使用GetDayPositionYangGui
func (lunar *Lunar) GetPositionYangGui() string {
	return lunar.GetDayPositionYangGui()
}

// GetPositionYangGuiDesc @Deprecated: 该方法已废弃，请使用GetDayPositionYangGuiDesc
func (lunar *Lunar) GetPositionYangGuiDesc() string {
	return lunar.GetDayPositionYangGuiDesc()
}

// GetPositionYinGui @Deprecated: 该方法已废弃，请使用GetDayPositionYinGui
func (lunar *Lunar) GetPositionYinGui() string {
	return lunar.GetDayPositionYinGui()
}

// GetPositionYinGuiDesc @Deprecated: 该方法已废弃，请使用GetDayPositionYinGuiDesc
func (lunar *Lunar) GetPositionYinGuiDesc() string {
	return lunar.GetDayPositionYinGuiDesc()
}

// GetPositionFu @Deprecated: 该方法已废弃，请使用GetDayPositionFu
func (lunar *Lunar) GetPositionFu() string {
	return lunar.GetDayPositionFu()
}

// GetPositionFuDesc @Deprecated: 该方法已废弃，请使用GetDayPositionFuDesc
func (lunar *Lunar) GetPositionFuDesc() string {
	return lunar.GetDayPositionFuDesc()
}

// GetPositionCai @Deprecated: 该方法已废弃，请使用GetDayPositionCai
func (lunar *Lunar) GetPositionCai() string {
	return lunar.GetDayPositionCai()
}

// GetPositionCaiDesc @Deprecated: 该方法已废弃，请使用GetDayPositionCaiDesc
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

func (lunar *Lunar) GetYearPositionTaiSui() string {
	return lunar.GetYearPositionTaiSuiBySect(2)
}

func (lunar *Lunar) GetYearPositionTaiSuiBySect(sect int) string {
	yearZhiIndex := 0
	switch sect {
	case 1:
		yearZhiIndex = lunar.yearZhiIndex
		break
	case 3:
		yearZhiIndex = lunar.yearZhiIndexExact
		break
	default:
		yearZhiIndex = lunar.yearZhiIndexByLiChun
	}
	return LunarUtil.POSITION_TAI_SUI_YEAR[yearZhiIndex]
}

func (lunar *Lunar) GetYearPositionTaiSuiDesc() string {
	return lunar.GetYearPositionTaiSuiDescBySect(2)
}

func (lunar *Lunar) GetYearPositionTaiSuiDescBySect(sect int) string {
	return LunarUtil.POSITION_DESC[lunar.GetYearPositionTaiSuiBySect(sect)]
}

func (lunar *Lunar) getMonthPositionTaiSui(monthZhiIndex int, monthGanIndex int) string {
	p := ""
	m := monthZhiIndex - LunarUtil.BASE_MONTH_ZHI_INDEX
	if m < 0 {
		m += 12
	}
	m = m % 4
	switch m {
	case 0:
		p = "艮"
		break
	case 2:
		p = "坤"
		break
	case 3:
		p = "巽"
		break
	default:
		p = LunarUtil.POSITION_GAN[monthGanIndex]
	}
	return p
}

func (lunar *Lunar) GetMonthPositionTaiSuiBySect(sect int) string {
	monthZhiIndex := 0
	monthGanIndex := 0
	switch sect {
	case 3:
		monthZhiIndex = lunar.monthZhiIndexExact
		monthGanIndex = lunar.monthGanIndexExact
		break
	default:
		monthZhiIndex = lunar.monthZhiIndex
		monthGanIndex = lunar.monthGanIndex
	}
	return lunar.getMonthPositionTaiSui(monthZhiIndex, monthGanIndex)
}

func (lunar *Lunar) GetMonthPositionTaiSui() string {
	return lunar.GetMonthPositionTaiSuiBySect(2)
}

func (lunar *Lunar) GetMonthPositionTaiSuiDesc() string {
	return lunar.GetMonthPositionTaiSuiDescBySect(2)
}

func (lunar *Lunar) GetMonthPositionTaiSuiDescBySect(sect int) string {
	return LunarUtil.POSITION_DESC[lunar.GetMonthPositionTaiSuiBySect(sect)]
}

func (lunar *Lunar) getDayPositionTaiSui(dayInGanZhi string, yearZhiIndex int) string {
	p := ""
	if strings.Contains("甲子,乙丑,丙寅,丁卯,戊辰,已巳", dayInGanZhi) {
		p = "震"
	} else if strings.Contains("丙子,丁丑,戊寅,已卯,庚辰,辛巳", dayInGanZhi) {
		p = "离"
	} else if strings.Contains("戊子,已丑,庚寅,辛卯,壬辰,癸巳", dayInGanZhi) {
		p = "中"
	} else if strings.Contains("庚子,辛丑,壬寅,癸卯,甲辰,乙巳", dayInGanZhi) {
		p = "兑"
	} else if strings.Contains("壬子,癸丑,甲寅,乙卯,丙辰,丁巳", dayInGanZhi) {
		p = "坎"
	} else {
		p = LunarUtil.POSITION_TAI_SUI_YEAR[yearZhiIndex]
	}
	return p
}

func (lunar *Lunar) GetDayPositionTaiSuiBySect(sect int) string {
	dayInGanZhi := ""
	yearZhiIndex := 0
	switch sect {
	case 1:
		dayInGanZhi = lunar.GetDayInGanZhi()
		yearZhiIndex = lunar.yearZhiIndex
		break
	case 3:
		dayInGanZhi = lunar.GetDayInGanZhi()
		yearZhiIndex = lunar.yearZhiIndexExact
		break
	default:
		dayInGanZhi = lunar.GetDayInGanZhiExact2()
		yearZhiIndex = lunar.yearZhiIndexByLiChun
	}
	return lunar.getDayPositionTaiSui(dayInGanZhi, yearZhiIndex)
}

func (lunar *Lunar) GetDayPositionTaiSui() string {
	return lunar.GetDayPositionTaiSuiBySect(2)
}

func (lunar *Lunar) GetDayPositionTaiSuiDesc() string {
	return lunar.GetDayPositionTaiSuiDescBySect(2)
}

func (lunar *Lunar) GetDayPositionTaiSuiDescBySect(sect int) string {
	return LunarUtil.POSITION_DESC[lunar.GetDayPositionTaiSuiBySect(sect)]
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

// GetChong @Deprecated: 该方法已废弃，请使用GetDayChong
func (lunar *Lunar) GetChong() string {
	return lunar.GetDayChong()
}

func (lunar *Lunar) GetDayChong() string {
	return LunarUtil.CHONG[lunar.dayZhiIndex]
}

// GetChongGan @Deprecated: 该方法已废弃，请使用GetDayChongGan
func (lunar *Lunar) GetChongGan() string {
	return lunar.GetDayChongGan()
}

func (lunar *Lunar) GetDayChongGan() string {
	return LunarUtil.CHONG_GAN[lunar.dayGanIndex]
}

// GetChongGanTie @Deprecated: 该方法已废弃，请使用GetDayChongGanTie
func (lunar *Lunar) GetChongGanTie() string {
	return lunar.GetDayChongGanTie()
}

func (lunar *Lunar) GetDayChongGanTie() string {
	return LunarUtil.CHONG_GAN_TIE[lunar.dayGanIndex]
}

// GetChongShengXiao @Deprecated: 该方法已废弃，请使用GetDayChongShengXiao
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

// GetChongDesc @Deprecated: 该方法已废弃，请使用GetDayChongDesc
func (lunar *Lunar) GetChongDesc() string {
	return lunar.GetDayChongDesc()
}

func (lunar *Lunar) GetDayChongDesc() string {
	return "(" + lunar.GetDayChongGan() + lunar.GetDayChong() + ")" + lunar.GetDayChongShengXiao()
}

// GetSha @Deprecated: 该方法已废弃，请使用GetDaySha
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

// GetBaZi @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZi() [4]string {
	baZi := lunar.GetEightChar()
	l := [4]string{}
	l[0] = baZi.GetYear()
	l[1] = baZi.GetMonth()
	l[2] = baZi.GetDay()
	l[3] = baZi.GetTime()
	return l
}

// GetBaZiWuXing @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiWuXing() [4]string {
	baZi := lunar.GetEightChar()
	l := [4]string{}
	l[0] = baZi.GetYearWuXing()
	l[1] = baZi.GetMonthWuXing()
	l[2] = baZi.GetDayWuXing()
	l[3] = baZi.GetTimeWuXing()
	return l
}

// GetBaZiNaYin @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiNaYin() [4]string {
	baZi := lunar.GetEightChar()
	l := [4]string{}
	l[0] = baZi.GetYearNaYin()
	l[1] = baZi.GetMonthNaYin()
	l[2] = baZi.GetDayNaYin()
	l[3] = baZi.GetTimeNaYin()
	return l
}

// GetBaZiShiShenGan @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiShiShenGan() [4]string {
	baZi := lunar.GetEightChar()
	l := [4]string{}
	l[0] = baZi.GetYearShiShenGan()
	l[1] = baZi.GetMonthShiShenGan()
	l[2] = baZi.GetDayShiShenGan()
	l[3] = baZi.GetTimeShiShenGan()
	return l
}

// GetBaZiShiShenZhi @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiShiShenZhi() [4]string {
	baZi := lunar.GetEightChar()
	l := [4]string{}
	l[0] = baZi.GetYearShiShenZhi().Front().Value.(string)
	l[1] = baZi.GetMonthShiShenZhi().Front().Value.(string)
	l[2] = baZi.GetDayShiShenZhi().Front().Value.(string)
	l[3] = baZi.GetTimeShiShenZhi().Front().Value.(string)
	return l
}

// GetBaZiShiShenYearZhi @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiShiShenYearZhi() *list.List {
	return lunar.GetEightChar().GetYearShiShenZhi()
}

// GetBaZiShiShenMonthZhi @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiShiShenMonthZhi() *list.List {
	return lunar.GetEightChar().GetMonthShiShenZhi()
}

// GetBaZiShiShenDayZhi @Deprecated: 该方法已废弃，请使用GetEightChar
func (lunar *Lunar) GetBaZiShiShenDayZhi() *list.List {
	return lunar.GetEightChar().GetDayShiShenZhi()
}

// GetBaZiShiShenTimeZhi @Deprecated: 该方法已废弃，请使用GetEightChar
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
	return LunarUtil.TIAN_SHEN[(lunar.dayZhiIndex+LunarUtil.ZHI_TIAN_SHEN_OFFSET[lunar.GetMonthZhi()])%12+1]
}

func (lunar *Lunar) GetTimeTianShen() string {
	return LunarUtil.TIAN_SHEN[(lunar.timeZhiIndex+LunarUtil.ZHI_TIAN_SHEN_OFFSET[lunar.GetDayZhiExact()])%12+1]
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
	return LunarUtil.CHONG[lunar.timeZhiIndex]
}

func (lunar *Lunar) GetTimeSha() string {
	return LunarUtil.SHA[lunar.GetTimeZhi()]
}

func (lunar *Lunar) GetTimeChongGan() string {
	return LunarUtil.CHONG_GAN[lunar.timeGanIndex]
}

func (lunar *Lunar) GetTimeChongGanTie() string {
	return LunarUtil.CHONG_GAN_TIE[lunar.timeGanIndex]
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
	return lunar.GetDayYiBySect(1)
}

func (lunar *Lunar) GetDayYiBySect(sect int) *list.List {
	monthGanZhi := lunar.GetMonthInGanZhi()
	if 2 == sect {
		monthGanZhi = lunar.GetMonthInGanZhiExact()
	}
	return LunarUtil.GetDayYi(monthGanZhi, lunar.GetDayInGanZhi())
}

func (lunar *Lunar) GetDayJi() *list.List {
	return lunar.GetDayJiBySect(1)
}

func (lunar *Lunar) GetDayJiBySect(sect int) *list.List {
	monthGanZhi := lunar.GetMonthInGanZhi()
	if 2 == sect {
		monthGanZhi = lunar.GetMonthInGanZhiExact()
	}
	return LunarUtil.GetDayJi(monthGanZhi, lunar.GetDayInGanZhi())
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

func (lunar *Lunar) getYearNineStar(yearInGanZhi string) *NineStar {
	indexExact := LunarUtil.GetJiaZiIndex(yearInGanZhi) + 1
	index := LunarUtil.GetJiaZiIndex(lunar.GetYearInGanZhi()) + 1
	yearOffset := indexExact - index
	if yearOffset > 1 {
		yearOffset -= 60
	} else if yearOffset < -1 {
		yearOffset += 60
	}
	yuan := int((lunar.year+yearOffset+2696)/60) % 3
	offset := (62 + yuan*3 - indexExact) % 9
	if 0 == offset {
		offset = 9
	}
	return NewNineStar(offset - 1)
}

func (lunar *Lunar) GetYearNineStarBySect(sect int) *NineStar {
	yearInGanZhi := ""
	switch sect {
	case 1:
		yearInGanZhi = lunar.GetYearInGanZhi()
		break
	case 3:
		yearInGanZhi = lunar.GetYearInGanZhiExact()
		break
	default:
		yearInGanZhi = lunar.GetYearInGanZhiByLiChun()
	}
	return lunar.getYearNineStar(yearInGanZhi)
}

func (lunar *Lunar) GetYearNineStar() *NineStar {
	return lunar.GetYearNineStarBySect(2)
}

func (lunar *Lunar) getMonthNineStar(yearZhiIndex int, monthZhiIndex int) *NineStar {
	index := yearZhiIndex % 3
	n := 27 - index*3
	if monthZhiIndex < LunarUtil.BASE_MONTH_ZHI_INDEX {
		n -= 3
	}
	offset := (n - monthZhiIndex) % 9
	return NewNineStar(offset)
}

func (lunar *Lunar) GetMonthNineStarBySect(sect int) *NineStar {
	yearZhiIndex := 0
	monthZhiIndex := 0
	switch sect {
	case 1:
		yearZhiIndex = lunar.yearZhiIndex
		monthZhiIndex = lunar.monthZhiIndex
		break
	case 3:
		yearZhiIndex = lunar.yearZhiIndexExact
		monthZhiIndex = lunar.monthZhiIndexExact
		break
	default:
		yearZhiIndex = lunar.yearZhiIndexByLiChun
		monthZhiIndex = lunar.monthZhiIndex
	}
	return lunar.getMonthNineStar(yearZhiIndex, monthZhiIndex)
}

func (lunar *Lunar) GetMonthNineStar() *NineStar {
	return lunar.GetMonthNineStarBySect(2)
}

func (lunar *Lunar) GetDayNineStar() *NineStar {
	solarYmd := lunar.solar.ToYmd()
	dongZhi := lunar.jieQi["冬至"]
	dongZhi2 := lunar.jieQi["DONG_ZHI"]
	xiaZhi := lunar.jieQi["夏至"]
	dongZhiIndex := LunarUtil.GetJiaZiIndex(dongZhi.GetLunar().GetDayInGanZhi())
	dongZhiIndex2 := LunarUtil.GetJiaZiIndex(dongZhi2.GetLunar().GetDayInGanZhi())
	xiaZhiIndex := LunarUtil.GetJiaZiIndex(xiaZhi.GetLunar().GetDayInGanZhi())
	solarShunBai := dongZhi
	solarShunBai2 := dongZhi2
	solarNiZi := xiaZhi
	if dongZhiIndex > 29 {
		solarShunBai = dongZhi.NextDay(60 - dongZhiIndex)
	} else {
		solarShunBai = dongZhi.NextDay(-dongZhiIndex)
	}
	solarShunBaiYmd := solarShunBai.ToYmd()
	if dongZhiIndex2 > 29 {
		solarShunBai2 = dongZhi2.NextDay(60 - dongZhiIndex2)
	} else {
		solarShunBai2 = dongZhi2.NextDay(-dongZhiIndex2)
	}
	solarShunBaiYmd2 := solarShunBai2.ToYmd()
	if xiaZhiIndex > 29 {
		solarNiZi = xiaZhi.NextDay(60 - xiaZhiIndex)
	} else {
		solarNiZi = xiaZhi.NextDay(-xiaZhiIndex)
	}
	solarNiZiYmd := solarNiZi.ToYmd()

	offset := 0
	if strings.Compare(solarYmd, solarShunBaiYmd) >= 0 && strings.Compare(solarYmd, solarNiZiYmd) < 0 {
		offset = lunar.solar.Subtract(solarShunBai) % 9
	} else if strings.Compare(solarYmd, solarNiZiYmd) >= 0 && strings.Compare(solarYmd, solarShunBaiYmd2) < 0 {
		offset = 8 - (lunar.solar.Subtract(solarNiZi) % 9)
	} else if strings.Compare(solarYmd, solarShunBaiYmd2) >= 0 {
		offset = lunar.solar.Subtract(solarShunBai2) % 9
	} else if strings.Compare(solarYmd, solarShunBaiYmd) < 0 {
		offset = (8 + solarShunBai.Subtract(lunar.solar)) % 9
	}
	return NewNineStar(offset)
}

func (lunar *Lunar) GetTimeNineStar() *NineStar {
	//顺逆
	solarYmd := lunar.solar.ToYmd()
	asc := false
	if strings.Compare(solarYmd, lunar.jieQi["冬至"].ToYmd()) >= 0 && strings.Compare(solarYmd, lunar.jieQi["夏至"].ToYmd()) < 0 {
		asc = true
	} else if strings.Compare(solarYmd, lunar.jieQi["DONG_ZHI"].ToYmd()) >= 0 {
		asc = true
	}
	start := 2
	if asc {
		start = 6
	}
	dayZhi := lunar.GetDayZhi()
	if strings.Contains("子午卯酉", dayZhi) {
		if asc {
			start = 0
		} else {
			start = 8
		}
	} else if strings.Contains("辰戌丑未", dayZhi) {
		if asc {
			start = 3
		} else {
			start = 5
		}
	}
	index := start + 9 - lunar.timeZhiIndex
	if asc {
		index = start + lunar.timeZhiIndex
	}
	return NewNineStar(index % 9)
}

// GetNextJie 获取下一节令（顺推的第一个节令）
func (lunar *Lunar) GetNextJie() *JieQi {
	return lunar.GetNextJieByWholeDay(false)
}

func (lunar *Lunar) GetNextJieByWholeDay(wholeDay bool) *JieQi {
	l := len(JIE_QI_IN_USE) / 2
	conditions := make([]string, l)
	for i := 0; i < l; i++ {
		conditions[i] = JIE_QI_IN_USE[i*2]
	}
	return lunar.getNearJieQi(true, conditions, wholeDay)
}

// GetPrevJie 获取上一节令（逆推的第一个节令）
func (lunar *Lunar) GetPrevJie() *JieQi {
	return lunar.GetPrevJieByWholeDay(false)
}

func (lunar *Lunar) GetPrevJieByWholeDay(wholeDay bool) *JieQi {
	l := len(JIE_QI_IN_USE) / 2
	conditions := make([]string, l)
	for i := 0; i < l; i++ {
		conditions[i] = JIE_QI_IN_USE[i*2]
	}
	return lunar.getNearJieQi(false, conditions, wholeDay)
}

// GetNextQi 获取下一气令（顺推的第一个气令）
func (lunar *Lunar) GetNextQi() *JieQi {
	return lunar.GetNextQiByWholeDay(false)
}

func (lunar *Lunar) GetNextQiByWholeDay(wholeDay bool) *JieQi {
	l := len(JIE_QI_IN_USE) / 2
	conditions := make([]string, l)
	for i := 0; i < l; i++ {
		conditions[i] = JIE_QI_IN_USE[i*2+1]
	}
	return lunar.getNearJieQi(true, conditions, wholeDay)
}

// GetPrevQi 获取上一气令（逆推的第一个气令）
func (lunar *Lunar) GetPrevQi() *JieQi {
	return lunar.GetPrevQiByWholeDay(false)
}

func (lunar *Lunar) GetPrevQiByWholeDay(wholeDay bool) *JieQi {
	l := len(JIE_QI_IN_USE) / 2
	conditions := make([]string, l)
	for i := 0; i < l; i++ {
		conditions[i] = JIE_QI_IN_USE[i*2+1]
	}
	return lunar.getNearJieQi(false, conditions, wholeDay)
}

// GetNextJieQi 获取下一节气（顺推的第一个节气）
func (lunar *Lunar) GetNextJieQi() *JieQi {
	return lunar.GetNextJieQiByWholeDay(false)
}

func (lunar *Lunar) GetNextJieQiByWholeDay(wholeDay bool) *JieQi {
	return lunar.getNearJieQi(true, nil, wholeDay)
}

// GetPrevJieQi 获取上一节气（逆推的第一个节气）
func (lunar *Lunar) GetPrevJieQi() *JieQi {
	return lunar.GetPrevJieQiByWholeDay(false)
}

func (lunar *Lunar) GetPrevJieQiByWholeDay(wholeDay bool) *JieQi {
	return lunar.getNearJieQi(false, nil, wholeDay)
}

// 获取最近的节气，如果未找到匹配的，返回null
func (lunar *Lunar) getNearJieQi(forward bool, conditions []string, wholeDay bool) *JieQi {
	var name string
	var near *Solar
	filters := map[string]bool{}
	if nil != conditions {
		for _, v := range conditions {
			filters[v] = true
		}
	}
	filter := len(filters) > 0
	today := ""
	if wholeDay {
		today = lunar.solar.ToYmd()
	} else {
		today = lunar.solar.ToYmdHms()
	}
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
		day := ""
		if wholeDay {
			day = solar.ToYmd()
		} else {
			day = solar.ToYmdHms()
		}
		if forward {
			if strings.Compare(day, today) < 0 {
				continue
			}
			if nil == near {
				name = jq
				near = solar
			} else {
				nearDay := ""
				if wholeDay {
					nearDay = near.ToYmd()
				} else {
					nearDay = near.ToYmdHms()
				}
				if strings.Compare(day, nearDay) < 0 {
					name = jq
					near = solar
				}
			}
		} else {
			if strings.Compare(day, today) > 0 {
				continue
			}
			if nil == near {
				name = jq
				near = solar
			} else {
				nearDay := ""
				if wholeDay {
					nearDay = near.ToYmd()
				} else {
					nearDay = near.ToYmdHms()
				}
				if strings.Compare(day, nearDay) > 0 {
					name = jq
					near = solar
				}
			}
		}
	}
	if nil == near {
		return nil
	}
	return NewJieQi(name, near)
}

// GetJieQi 获取节气名称，如果无节气，返回空字符串
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

// GetCurrentJieQi 获取当天节气对象，如果无节气，返回nil
func (lunar *Lunar) GetCurrentJieQi() *JieQi {
	name := lunar.GetJieQi()
	if len(name) > 0 {
		return NewJieQi(name, lunar.solar)
	}
	return nil
}

// GetCurrentJie 获取当天节令对象，如果无节令，返回nil
func (lunar *Lunar) GetCurrentJie() *JieQi {
	name := lunar.GetJie()
	if len(name) > 0 {
		return NewJieQi(name, lunar.solar)
	}
	return nil
}

// GetCurrentQi 获取当天气令对象，如果无气令，返回nil
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

// Next 获取往后推几天的农历日期，如果要往前推，则天数用负数
func (lunar *Lunar) Next(days int) *Lunar {
	return lunar.solar.NextDay(days).GetLunar()
}

// GetYearXun 获取年所在旬（以正月初一作为新年的开始）
func (lunar *Lunar) GetYearXun() string {
	return LunarUtil.GetXun(lunar.GetYearInGanZhi())
}

// GetYearXunByLiChun 获取年所在旬（以立春当天作为新年的开始）
func (lunar *Lunar) GetYearXunByLiChun() string {
	return LunarUtil.GetXun(lunar.GetYearInGanZhiByLiChun())
}

// GetYearXunExact 获取年所在旬（以立春交接时刻作为新年的开始）
func (lunar *Lunar) GetYearXunExact() string {
	return LunarUtil.GetXun(lunar.GetYearInGanZhiExact())
}

// GetYearXunKong 获取值年空亡（以正月初一作为新年的开始）
func (lunar *Lunar) GetYearXunKong() string {
	return LunarUtil.GetXunKong(lunar.GetYearInGanZhi())
}

// GetYearXunKongByLiChun 获取值年空亡（以立春当天作为新年的开始）
func (lunar *Lunar) GetYearXunKongByLiChun() string {
	return LunarUtil.GetXunKong(lunar.GetYearInGanZhiByLiChun())
}

// GetYearXunKongExact 获取值年空亡（以立春交接时刻作为新年的开始）
func (lunar *Lunar) GetYearXunKongExact() string {
	return LunarUtil.GetXunKong(lunar.GetYearInGanZhiExact())
}

// GetMonthXun 获取月所在旬（以节交接当天起算）
func (lunar *Lunar) GetMonthXun() string {
	return LunarUtil.GetXun(lunar.GetMonthInGanZhi())
}

// GetMonthXunExact 获取月所在旬（以节交接时刻起算）
func (lunar *Lunar) GetMonthXunExact() string {
	return LunarUtil.GetXun(lunar.GetMonthInGanZhiExact())
}

// GetMonthXunKong 获取值月空亡（以节交接当天起算）
func (lunar *Lunar) GetMonthXunKong() string {
	return LunarUtil.GetXunKong(lunar.GetMonthInGanZhi())
}

// GetMonthXunKongExact 获取值月空亡（以节交接时刻起算）
func (lunar *Lunar) GetMonthXunKongExact() string {
	return LunarUtil.GetXunKong(lunar.GetMonthInGanZhiExact())
}

// GetDayXun 获取日所在旬（以节交接当天起算）
func (lunar *Lunar) GetDayXun() string {
	return LunarUtil.GetXun(lunar.GetDayInGanZhi())
}

// GetDayXunExact 获取日所在旬（晚子时日柱算明天）
func (lunar *Lunar) GetDayXunExact() string {
	return LunarUtil.GetXun(lunar.GetDayInGanZhiExact())
}

// GetDayXunExact2 获取日所在旬（晚子时日柱算当天）
func (lunar *Lunar) GetDayXunExact2() string {
	return LunarUtil.GetXun(lunar.GetDayInGanZhiExact2())
}

// GetDayXunKong 获取值日空亡
func (lunar *Lunar) GetDayXunKong() string {
	return LunarUtil.GetXunKong(lunar.GetDayInGanZhi())
}

// GetDayXunKongExact 获取值日空亡（晚子时日柱算明天）
func (lunar *Lunar) GetDayXunKongExact() string {
	return LunarUtil.GetXunKong(lunar.GetDayInGanZhiExact())
}

// GetDayXunKongExact2 获取值日空亡（晚子时日柱算当天）
func (lunar *Lunar) GetDayXunKongExact2() string {
	return LunarUtil.GetXunKong(lunar.GetDayInGanZhiExact2())
}

// GetTimeXun 获取时辰所在旬
func (lunar *Lunar) GetTimeXun() string {
	return LunarUtil.GetXun(lunar.GetTimeInGanZhi())
}

// GetTimeXunKong 获取值时空亡
func (lunar *Lunar) GetTimeXunKong() string {
	return LunarUtil.GetXunKong(lunar.GetTimeInGanZhi())
}

// GetShuJiu 获取数九，如果不是数九天，返回nil
func (lunar *Lunar) GetShuJiu() *ShuJiu {
	current := NewSolarFromYmd(lunar.solar.GetYear(), lunar.solar.GetMonth(), lunar.solar.GetDay())
	start := lunar.jieQi["DONG_ZHI"]
	start = NewSolarFromYmd(start.GetYear(), start.GetMonth(), start.GetDay())
	if current.IsBefore(start) {
		start = lunar.jieQi["冬至"]
		start = NewSolarFromYmd(start.GetYear(), start.GetMonth(), start.GetDay())
	}
	end := start.NextDay(81)
	if current.IsBefore(start) || !current.IsBefore(end) {
		return nil
	}
	days := current.Subtract(start)
	return NewShuJiu(LunarUtil.NUMBER[days/9+1]+"九", days%9+1)
}

// GetFu 获取三伏，如果不是三伏天，返回nil
func (lunar *Lunar) GetFu() *Fu {
	current := NewSolarFromYmd(lunar.solar.GetYear(), lunar.solar.GetMonth(), lunar.solar.GetDay())
	xiaZhi := lunar.jieQi["夏至"]
	liQiu := lunar.jieQi["立秋"]
	start := NewSolarFromYmd(xiaZhi.GetYear(), xiaZhi.GetMonth(), xiaZhi.GetDay())
	// 第1个庚日
	add := 6 - xiaZhi.GetLunar().GetDayGanIndex()
	if add < 0 {
		add += 10
	}
	// 第3个庚日，即初伏第1天
	add += 20
	start = start.NextDay(add)

	// 初伏以前
	if current.IsBefore(start) {
		return nil
	}

	days := current.Subtract(start)
	if days < 10 {
		return NewFu("初伏", days+1)
	}

	// 第4个庚日，中伏第1天
	start = start.NextDay(10)
	days = current.Subtract(start)
	if days < 10 {
		return NewFu("中伏", days+1)
	}

	// 第5个庚日，中伏第11天或末伏第1天
	start = start.NextDay(10)
	days = current.Subtract(start)

	liQiuSolar := NewSolarFromYmd(liQiu.GetYear(), liQiu.GetMonth(), liQiu.GetDay())

	// 末伏
	if liQiuSolar.IsAfter(start) {
		// 中伏
		if days < 10 {
			return NewFu("中伏", days+11)
		}
		// 末伏第1天
		start = start.NextDay(10)
		days = current.Subtract(start)
	}
	if days < 10 {
		return NewFu("末伏", days+1)
	}
	return nil
}

// GetLiuYao 获取六曜
func (lunar *Lunar) GetLiuYao() string {
	month := lunar.month
	if month < 0 {
		month = 0 - month
	}
	return LunarUtil.LIU_YAO[(month+lunar.day-2)%6]
}

// GetHou 获取候
func (lunar *Lunar) GetHou() string {
	jq := lunar.GetPrevJieQiByWholeDay(true)
	max := len(LunarUtil.HOU) - 1
	offset := lunar.solar.Subtract(jq.GetSolar()) / 5
	if offset > max {
		offset = max
	}
	return fmt.Sprintf("%s %s", jq.GetName(), LunarUtil.HOU[offset])
}

// GetWuHou 获取物候
func (lunar *Lunar) GetWuHou() string {
	jq := lunar.GetPrevJieQiByWholeDay(true)
	offset := 0
	for i, v := range JIE_QI {
		if strings.Compare(jq.GetName(), v) == 0 {
			offset = i
			break
		}
	}
	index := lunar.solar.Subtract(jq.GetSolar()) / 5
	if index > 2 {
		index = 2
	}
	return LunarUtil.WU_HOU[(offset*3+index)%len(LunarUtil.WU_HOU)]
}

// GetDayLu 获取日禄
func (lunar *Lunar) GetDayLu() string {
	gan := LunarUtil.LU[lunar.GetDayGan()]
	lu := gan + "命互禄"
	if zhi, ok := LunarUtil.LU[lunar.GetDayZhi()]; ok {
		lu += " " + zhi + "命进禄"
	}
	return lu
}

// GetTime 获取时辰
func (lunar *Lunar) GetTime() *LunarTime {
	return NewLunarTime(lunar.year, lunar.month, lunar.day, lunar.hour, lunar.minute, lunar.second)
}

// GetTimes 获取当天的时辰列表
func (lunar *Lunar) GetTimes() []*LunarTime {
	l := make([]*LunarTime, 13)
	l[0] = NewLunarTime(lunar.year, lunar.month, lunar.day, 0, 0, 0)
	for i := 0; i < 12; i++ {
		l[i+1] = NewLunarTime(lunar.year, lunar.month, lunar.day, (i+1)*2-1, 0, 0)
	}
	return l
}

// GetFoto 获取佛历
func (lunar *Lunar) GetFoto() *Foto {
	return NewFotoFromLunar(lunar)
}

// GetTao 获取道历
func (lunar *Lunar) GetTao() *Tao {
	return NewTaoFromLunar(lunar)
}
