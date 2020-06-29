package calendar

import (
	"container/list"
	"lunar-go/LunarUtil"
	"lunar-go/SolarUtil"
	"math"
	"strconv"
	"strings"
	"time"
)

const radPerDegree = math.Pi / 180
const degreePerRad = 180 / math.Pi
const secondPerRad = 180 * 3600 / math.Pi

var jieQi = []string{"冬至","小寒","大寒","立春","雨水","惊蛰","春分","清明","谷雨","立夏","小满","芒种","夏至","小暑","大暑","立秋","处暑","白露","秋分","寒露","霜降","立冬","小雪","大雪"}
/** 黄经周期项 */
var e10 = []float64{1.75347045673, 0.00000000000, 0.0000000000, 0.03341656456, 4.66925680417, 6283.0758499914, 0.00034894275, 4.62610241759, 12566.1516999828, 0.00003417571, 2.82886579606, 3.5231183490, 0.00003497056, 2.74411800971, 5753.3848848968, 0.00003135896, 3.62767041758, 77713.7714681205, 0.00002676218, 4.41808351397, 7860.4193924392, 0.00002342687, 6.13516237631, 3930.2096962196, 0.00001273166, 2.03709655772, 529.6909650946, 0.00001324292, 0.74246356352, 11506.7697697936, 0.00000901855, 2.04505443513, 26.2983197998, 0.00001199167, 1.10962944315, 1577.3435424478, 0.00000857223, 3.50849156957, 398.1490034082, 0.00000779786, 1.17882652114, 5223.6939198022, 0.00000990250, 5.23268129594, 5884.9268465832, 0.00000753141, 2.53339053818, 5507.5532386674, 0.00000505264, 4.58292563052, 18849.2275499742, 0.00000492379, 4.20506639861, 775.5226113240, 0.00000356655, 2.91954116867, 0.0673103028, 0.00000284125, 1.89869034186, 796.2980068164, 0.00000242810, 0.34481140906, 5486.7778431750, 0.00000317087, 5.84901952218, 11790.6290886588, 0.00000271039, 0.31488607649, 10977.0788046990, 0.00000206160, 4.80646606059, 2544.3144198834, 0.00000205385, 1.86947813692, 5573.1428014331, 0.00000202261, 2.45767795458, 6069.7767545534, 0.00000126184, 1.08302630210, 20.7753954924, 0.00000155516, 0.83306073807, 213.2990954380, 0.00000115132, 0.64544911683, 0.9803210682, 0.00000102851, 0.63599846727, 4694.0029547076, 0.00000101724, 4.26679821365, 7.1135470008, 0.00000099206, 6.20992940258, 2146.1654164752, 0.00000132212, 3.41118275555, 2942.4634232916, 0.00000097607, 0.68101272270, 155.4203994342, 0.00000085128, 1.29870743025, 6275.9623029906, 0.00000074651, 1.75508916159, 5088.6288397668, 0.00000101895, 0.97569221824, 15720.8387848784, 0.00000084711, 3.67080093025, 71430.6956181291, 0.00000073547, 4.67926565481, 801.8209311238, 0.00000073874, 3.50319443167, 3154.6870848956, 0.00000078756, 3.03698313141, 12036.4607348882, 0.00000079637, 1.80791330700, 17260.1546546904, 0.00000085803, 5.98322631256, 161000.6857376741, 0.00000056963, 2.78430398043, 6286.5989683404, 0.00000061148, 1.81839811024, 7084.8967811152, 0.00000069627, 0.83297596966, 9437.7629348870, 0.00000056116, 4.38694880779, 14143.4952424306, 0.00000062449, 3.97763880587, 8827.3902698748, 0.00000051145, 0.28306864501, 5856.4776591154, 0.00000055577, 3.47006009062, 6279.5527316424, 0.00000041036, 5.36817351402, 8429.2412664666, 0.00000051605, 1.33282746983, 1748.0164130670, 0.00000051992, 0.18914945834, 12139.5535091068, 0.00000049000, 0.48735065033, 1194.4470102246, 0.00000039200, 6.16832995016, 10447.3878396044, 0.00000035566, 1.77597314691, 6812.7668150860, 0.00000036770, 6.04133859347, 10213.2855462110, 0.00000036596, 2.56955238628, 1059.3819301892, 0.00000033291, 0.59309499459, 17789.8456197850, 0.00000035954, 1.70876111898, 2352.8661537718}
/** 黄经泊松1项 */
var e11 = []float64{6283.31966747491, 0.00000000000, 0.0000000000, 0.00206058863, 2.67823455584, 6283.0758499914, 0.00004303430, 2.63512650414, 12566.1516999828, 0.00000425264, 1.59046980729, 3.5231183490, 0.00000108977, 2.96618001993, 1577.3435424478, 0.00000093478, 2.59212835365, 18849.2275499742, 0.00000119261, 5.79557487799, 26.2983197998, 0.00000072122, 1.13846158196, 529.6909650946, 0.00000067768, 1.87472304791, 398.1490034082, 0.00000067327, 4.40918235168, 5507.5532386674, 0.00000059027, 2.88797038460, 5223.6939198022, 0.00000055976, 2.17471680261, 155.4203994342, 0.00000045407, 0.39803079805, 796.2980068164, 0.00000036369, 0.46624739835, 775.5226113240, 0.00000028958, 2.64707383882, 7.1135470008, 0.00000019097, 1.84628332577, 5486.7778431750, 0.00000020844, 5.34138275149, 0.9803210682, 0.00000018508, 4.96855124577, 213.2990954380, 0.00000016233, 0.03216483047, 2544.3144198834, 0.00000017293, 2.99116864949, 6275.9623029906}
/** 黄经泊松2项 */
var e12 = []float64{0.00052918870, 0.00000000000, 0.0000000000, 0.00008719837, 1.07209665242, 6283.0758499914, 0.00000309125, 0.86728818832, 12566.1516999828, 0.00000027339, 0.05297871691, 3.5231183490, 0.00000016334, 5.18826691036, 26.2983197998, 0.00000015752, 3.68457889430, 155.4203994342, 0.00000009541, 0.75742297675, 18849.2275499742, 0.00000008937, 2.05705419118, 77713.7714681205, 0.00000006952, 0.82673305410, 775.5226113240, 0.00000005064, 4.66284525271, 1577.3435424478}
var e13 = []float64{0.00000289226, 5.84384198723, 6283.0758499914, 0.00000034955, 0.00000000000, 0.0000000000, 0.00000016819, 5.48766912348, 12566.1516999828}
var e14 = []float64{0.00000114084, 3.14159265359, 0.0000000000, 0.00000007717, 4.13446589358, 6283.0758499914, 0.00000000765, 3.83803776214, 12566.1516999828}
var e15 = []float64{0.00000000878, 3.14159265359, 0.0000000000}
/** 黄纬周期项 */
var e20 = []float64{0.00000279620, 3.19870156017, 84334.6615813083, 0.00000101643, 5.42248619256, 5507.5532386674, 0.00000080445, 3.88013204458, 5223.6939198022, 0.00000043806, 3.70444689758, 2352.8661537718, 0.00000031933, 4.00026369781, 1577.3435424478, 0.00000022724, 3.98473831560, 1047.7473117547, 0.00000016392, 3.56456119782, 5856.4776591154, 0.00000018141, 4.98367470263, 6283.0758499914, 0.00000014443, 3.70275614914, 9437.7629348870, 0.00000014304, 3.41117857525, 10213.2855462110}
var e21 = []float64{0.00000009030, 3.89729061890, 5507.5532386674, 0.00000006177, 1.73038850355, 5223.6939198022}
/** 离心率 */
var gxcE = []float64{0.016708634, -0.000042037, -0.0000001267}
/** 近点 */
var gxcP = []float64{102.93735 / degreePerRad, 1.71946 / degreePerRad, 0.00046 / degreePerRad}
/** 太平黄经 */
var gxcL = []float64{280.4664567 / degreePerRad, 36000.76982779 / degreePerRad, 0.0003032028 / degreePerRad, 1 / 49931000 / degreePerRad, -1 / 153000000 / degreePerRad}
/** 光行差常数 */
var gxcK = 20.49552 / secondPerRad
var zd = []float64{2.1824391966, -33.757045954, 0.0000362262, 3.7340E-08, -2.8793E-10, -171996, -1742, 92025, 89, 3.5069406862, 1256.663930738, 0.0000105845, 6.9813E-10, -2.2815E-10, -13187, -16, 5736, -31, 1.3375032491, 16799.418221925, -0.0000511866, 6.4626E-08, -5.3543E-10, -2274, -2, 977, -5, 4.3648783932, -67.514091907, 0.0000724525, 7.4681E-08, -5.7586E-10, 2062, 2, -895, 5, 0.0431251803, -628.301955171, 0.0000026820, 6.5935E-10, 5.5705E-11, -1426, 34, 54, -1, 2.3555557435, 8328.691425719, 0.0001545547, 2.5033E-07, -1.1863E-09, 712, 1, -7, 0, 3.4638155059, 1884.965885909, 0.0000079025, 3.8785E-11, -2.8386E-10, -517, 12, 224, -6, 5.4382493597, 16833.175267879, -0.0000874129, 2.7285E-08, -2.4750E-10, -386, -4, 200, 0, 3.6930589926, 25128.109647645, 0.0001033681, 3.1496E-07, -1.7218E-09, -301, 0, 129, -1, 3.5500658664, 628.361975567, 0.0000132664, 1.3575E-09, -1.7245E-10, 217, -5, -95, 3}

type Lunar struct {
	year int
	month int
	day int
	hour int
	minute int
	second int
	dayOffset int
	yearGanIndex int
	yearZhiIndex int
	yearGanIndexByLiChun int
	yearZhiIndexByLiChun int
	yearGanIndexExact int
	yearZhiIndexExact int
	monthGanIndex int
	monthZhiIndex int
	monthGanIndexExact int
	monthZhiIndexExact int
	dayGanIndex int
	dayZhiIndex int
	dayGanIndexExact int
	dayZhiIndexExact int
	timeGanIndex int
	timeZhiIndex int
	weekIndex int
	jieQi map[string]*Solar
	solar *Solar
}

func NewLunar(lunarYear int, lunarMonth int, lunarDay int, hour int,minute int,second int) *Lunar {
	lunar := new(Lunar)
	lunar.year = lunarYear
	lunar.month = lunarMonth
	lunar.day = lunarDay
	lunar.hour = hour
	lunar.minute = minute
	lunar.second = second
	lunar.dayOffset = LunarUtil.ComputeAddDays(lunarYear,lunarMonth,lunarDay)
	lunar.solar = toSolar(lunar)
	compute(lunar)
	return lunar
}

func NewLunarFromYmd(lunarYear int, lunarMonth int, lunarDay int) *Lunar {
	return NewLunar(lunarYear,lunarMonth,lunarDay,0,0,0)
}

func NewLunarFromDate(date time.Time) *Lunar {
	solar := NewSolarFromDate(date)
	y := solar.year
	m := solar.month
	d := solar.day
	startYear := 0
	startMonth := 0
	startDay := 0
	lunarYear := 0
	lunarMonth := 0
	lunarDay := 0
	if y<2000 {
		startYear = SolarUtil.BASE_YEAR
		startMonth = SolarUtil.BASE_MONTH
		startDay = SolarUtil.BASE_DAY
		lunarYear = LunarUtil.BASE_YEAR
		lunarMonth = LunarUtil.BASE_MONTH
		lunarDay = LunarUtil.BASE_DAY
	}else{
		startYear = SolarUtil.BASE_YEAR+99
		startMonth = 1
		startDay = 1
		lunarYear = LunarUtil.BASE_YEAR+99
		lunarMonth = 11
		lunarDay = 25
	}
	diff := 0
	for i:=startYear;i<y;i++ {
		diff += 365
		if SolarUtil.IsLeapYear(i) {
			diff += 1
		}
	}
	for i:=startMonth;i<m;i++ {
		diff += SolarUtil.GetDaysOfMonth(y,i)
	}
	diff += d-startDay
	lunarDay += diff
	lastDate := LunarUtil.GetDaysOfMonth(lunarYear,lunarMonth)
	for{
		if lunarDay<=lastDate {
			break
		}
		lunarDay -= lastDate
		lunarMonth = LunarUtil.NextMonth(lunarYear,lunarMonth)
		if lunarMonth==1 {
			lunarYear++
		}
		lastDate = LunarUtil.GetDaysOfMonth(lunarYear,lunarMonth)
	}
	return NewLunar(lunarYear,lunarMonth,lunarDay,solar.hour,solar.minute,solar.second)
}

func computeJieQi(lunar *Lunar) {
	//儒略日，冬至在阳历上一年，所以这里多减1年以从去年开始
	jd := 365.2422 * float64(lunar.solar.year-2001)
	j := len(jieQi)
	table := make(map[string]*Solar)
	for i := 0; i < j; i++ {
		t := calJieQi(jd + float64(i)*15.2, float64(i)*15-90) + J2000 + 8 / 24
		table[jieQi[i]] = NewSolarFromJulianDay(t)
	}
	lunar.jieQi = table
}

func calJieQi(t1 float64,degree float64) float64{
	// 对于节气计算,应满足t在t1到t1+360天之间,对于Y年第n个节气(n=0是春分),t1可取值Y*365.2422+n*15.2
	t2 := t1
	t := float64(0)
	v := float64(0)
	// 在t1到t2范围内求解(范气360天范围),结果置于t
	t2 += 360
	// 待搜索目标角
	rad := degree * radPerDegree
	// 利用截弦法计算
	// v1,v2为t1,t2时对应的黄经
	v1 := calRad(t1, rad)
	v2 := calRad(t2, rad)
	// 减2pi作用是将周期性角度转为连续角度
	if v1 < v2 {
		v2 -= 2 * math.Pi
	}
	// k是截弦的斜率
	k := float64(1)
	k2 := float64(0)
	// 快速截弦求根,通常截弦三四次就已达所需精度
	for i := 0; i < 10; i++ {
		// 算出斜率
		k2 = (v2 - v1) / (t2 - t1)
		// 差商可能为零,应排除
		if math.Abs(k2) > 1e-15 {
			k = k2
		}
		t = t1 - v1 / k
		// 直线逼近法求根(直线方程的根)
		v = calRad(t, rad)
		// 一次逼近后,v1就已接近0,如果很大,则应减1周
		if v > 1 {
			v -= 2 * math.Pi
		}
		// 已达精度
		if math.Abs(v) < 1e-8 {
			break
		}
		t1 = t2
		v1 = v2
		t2 = t
		// 下一次截弦
		v2 = v
	}
	return t
}

func mrad(rad float64) float64{
	pi2 := 2 * math.Pi
	rad = math.Mod(rad,pi2)
	if rad<0 {
		return rad+pi2
	}else{
		return rad
	}
}

func enn(f []float64,ennt float64) float64{
	v := float64(0)
	j := len(f)
	for i:=0;i<j;i+=3 {
		v += f[i] * math.Cos(f[i + 1] + ennt * f[i + 2])
	}
	return v
}

func gxc(t float64,pos []float64) {
	t1 := t / 36525
	t2 := t1 * t1
	t3 := t2 * t1
	t4 := t3 * t1
	l := gxcL[0] + gxcL[1] * t1 + gxcL[2] * t2 + gxcL[3] * t3 + gxcL[4] * t4
	p := gxcP[0] + gxcP[1] * t1 + gxcP[2] * t2
	e := gxcE[0] + gxcE[1] * t1 + gxcE[2] * t2
	dl := l - pos[0]
	dp := p - pos[0]
	pos[0] -= gxcK * (math.Cos(dl) - e * math.Cos(dp)) / math.Cos(pos[1])
	pos[1] -= gxcK * math.Sin(pos[1]) * (math.Sin(dl) - e * math.Sin(dp))
	pos[0] = mrad(pos[0])
}

func calRad(t float64,rad float64) float64{
	// 计算太阳真位置(先算出日心坐标中地球的位置)
	pos := calEarth(t)
	pos[0] += math.Pi
	// 转为地心坐标
	pos[1] = -pos[1]
	// 补周年光行差
	gxc(t, pos)
	// 补黄经章动
	pos[0] += hjzd(t)
	return mrad(rad - pos[0])
}

func calEarth(t float64) []float64{
	t1 := t / 365250
	r := []float64{0,0}
	t2 := t1 * t1
	t3 := t2 * t1
	t4 := t3 * t1
	t5 := t4 * t1
	r[0] = mrad(enn(e10,t1) + enn(e11,t1) * t1 + enn(e12,t1) * t2 + enn(e13,t1) * t3 + enn(e14,t1) * t4 + enn(e15,t1) * t5)
	r[1] = enn(e20,t1) + enn(e21,t1) * t1
	return r
}

func hjzd(t float64) float64{
	lon := float64(0)
	t1 := t / 36525
	c := float64(0)
	t2 := t1 * t1
	t3 := t2 * t1
	t4 := t3 * t1
	j := len(zd)
	for i := 0; i < j; i += 9 {
		c = zd[i] + zd[i + 1] * t1 + zd[i + 2] * t2 + zd[i + 3] * t3 + zd[i + 4] * t4
		lon += (zd[i + 5] + zd[i + 6] * t1 / 10) * math.Sin(c)
	}
	lon /= secondPerRad * 10000
	return lon
}

func computeYear(lunar *Lunar) {
	lunar.yearGanIndex = (lunar.year+LunarUtil.BASE_YEAR_GANZHI_INDEX)%10
	lunar.yearZhiIndex = (lunar.year+LunarUtil.BASE_YEAR_GANZHI_INDEX)%12

	//以立春作为新一年的开始的干支纪年
	g := lunar.yearGanIndex
	z := lunar.yearZhiIndex

	//精确的干支纪年，以立春交接时刻为准
	gExact := lunar.yearGanIndex
	zExact := lunar.yearZhiIndex

	solar := lunar.solar

	if lunar.year==solar.GetYear() {
		//获取立春的阳历时刻
		liChun := lunar.jieQi["立春"]
		//立春日期判断
		if strings.Compare(solar.ToYmd(),liChun.ToYmd())<0 {
			g--
			if g<0 {
				g += 10
			}
			z--
			if z<0 {
				z += 12
			}
		}
		//立春交接时刻判断
		if strings.Compare(solar.ToYmdHms(),liChun.ToYmdHms())<0 {
			gExact--
			if gExact<0 {
				gExact += 10
			}
			zExact--
			if zExact<0 {
				zExact += 12
			}
		}
	}
	lunar.yearGanIndexByLiChun = g
	lunar.yearZhiIndexByLiChun = z

	lunar.yearGanIndexExact = gExact
	lunar.yearZhiIndexExact = zExact
}


func computeMonth(lunar *Lunar) {
	var start *Solar
	var end *Solar
	//干偏移值（以立春当天起算）
	gOffset := ((lunar.yearGanIndexByLiChun%5+1)*2)%10
	//干偏移值（以立春交接时刻起算）
	gOffsetExact := ((lunar.yearGanIndexExact%5+1)*2)%10

	//序号：大雪到小寒之间-2，小寒到立春之间-1，立春之后0
	index := -2
	for i:=0;i<len(LunarUtil.JIE);i++{
		jie := LunarUtil.JIE[i]
		end = lunar.jieQi[jie]
		ymd := lunar.solar.ToYmd()
		symd := ymd
		if start!=nil {
			symd = start.ToYmd()
		}
		eymd := end.ToYmd()
		if strings.Compare(ymd,symd)>=0 && strings.Compare(ymd,eymd)<0  {
			break
		}
		start = end
		index++
	}
	if index<0 {
		index += 12
	}

	lunar.monthGanIndex = (index+gOffset)%10
	lunar.monthZhiIndex = (index+LunarUtil.BASE_MONTH_ZHI_INDEX)%12

	//序号：大雪到小寒之间-2，小寒到立春之间-1，立春之后0
	indexExact := -2
	for i:=0;i<len(LunarUtil.JIE);i++{
		jie := LunarUtil.JIE[i]
		end = lunar.jieQi[jie]
		tm := lunar.solar.ToYmdHms()
		stime := tm
		if start!=nil {
			stime = start.ToYmdHms()
		}
		etime := end.ToYmdHms()
		if strings.Compare(tm,stime)>=0 && strings.Compare(tm,etime)<0  {
			break
		}
		start = end
		indexExact++
	}
	if indexExact<0 {
		indexExact += 12
	}
	lunar.monthGanIndexExact = (indexExact+gOffsetExact)%10
	lunar.monthZhiIndexExact = (indexExact+LunarUtil.BASE_MONTH_ZHI_INDEX)%12
}

func computeDay(lunar *Lunar) {
	addDays := (lunar.dayOffset + LunarUtil.BASE_DAY_GANZHI_INDEX)%60
	lunar.dayGanIndex = addDays%10
	lunar.dayZhiIndex = addDays%12

	dayGanExact := lunar.dayGanIndex
	dayZhiExact := lunar.dayZhiIndex

	// 晚子时（夜子/子夜）应算作第二天
	hm := padding(lunar.hour)+":"+padding(lunar.minute)
	if strings.Compare(hm,"23:00")>=0 && strings.Compare(hm,"23:59")<=0 {
		dayGanExact++
		if dayGanExact>=10 {
			dayGanExact -= 10
		}
		dayZhiExact++
		if dayZhiExact>=12 {
			dayZhiExact -= 12
		}
	}

	lunar.dayGanIndexExact = dayGanExact
	lunar.dayZhiIndexExact = dayZhiExact
}

func computeTime(lunar *Lunar) {
	hm := padding(lunar.hour)+":"+padding(lunar.minute)
	lunar.timeZhiIndex = LunarUtil.GetTimeZhiIndex(hm)
	lunar.timeGanIndex = (lunar.dayGanIndexExact%5*2+lunar.timeZhiIndex)%10
}

func computeWeek(lunar *Lunar) {
	lunar.weekIndex = (lunar.dayOffset+LunarUtil.BASE_WEEK_INDEX)%7
}

func compute(lunar *Lunar) {
	computeJieQi(lunar)
	computeYear(lunar)
	computeMonth(lunar)
	computeDay(lunar)
	computeTime(lunar)
	computeWeek(lunar)
}

func toSolar(lunar *Lunar) *Solar {
	c := time.Date(SolarUtil.BASE_YEAR,time.Month(SolarUtil.BASE_MONTH),SolarUtil.BASE_DAY,lunar.hour,lunar.minute,lunar.second,0,time.Local)
	c = c.AddDate(0,0,lunar.dayOffset)
	return NewSolarFromDate(c)
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
	return lunar.GetYearGan()+lunar.GetYearZhi()
}

func (lunar *Lunar) GetYearInGanZhiByLiChun() string {
	return lunar.GetYearGanByLiChun()+lunar.GetYearZhiByLiChun()
}

func (lunar *Lunar) GetYearInGanZhiExact() string {
	return lunar.GetYearGanExact()+lunar.GetYearZhiExact()
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
	return lunar.GetMonthGan()+lunar.GetMonthZhi()
}

func (lunar *Lunar) GetMonthInGanZhiExact() string {
	return lunar.GetMonthGanExact()+lunar.GetMonthZhiExact()
}

func (lunar *Lunar) GetDayGan() string {
	return LunarUtil.GAN[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetDayGanExact() string {
	return LunarUtil.GAN[lunar.dayGanIndexExact+1]
}

func (lunar *Lunar) GetDayZhi() string {
	return LunarUtil.ZHI[lunar.dayZhiIndex+1]
}

func (lunar *Lunar) GetDayZhiExact() string {
	return LunarUtil.ZHI[lunar.dayZhiIndexExact+1]
}

func (lunar *Lunar) GetDayInGanZhi() string {
	return lunar.GetDayGan()+lunar.GetDayZhi()
}

func (lunar *Lunar) GetDayInGanZhiExact() string {
	return lunar.GetDayGanExact()+lunar.GetDayZhiExact()
}

func (lunar *Lunar) GetTimeGan() string {
	return LunarUtil.GAN[lunar.timeGanIndex+1]
}

func (lunar *Lunar) GetTimeZhi() string {
	return LunarUtil.ZHI[lunar.timeZhiIndex+1]
}

func (lunar *Lunar) GetTimeInGanZhi() string {
	return lunar.GetTimeGan()+lunar.GetTimeZhi()
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
	y := strconv.Itoa(lunar.year)
	s := ""
	j := len(y)
	for i:=0;i<j;i++ {
		s += LunarUtil.NUMBER[[]rune(y[i:i+1])[0]-'0']
	}
	return s
}

func (lunar *Lunar) GetMonthInChinese() string {
	s := ""
	if lunar.month<0 {
		s += "闰"
		s += LunarUtil.MONTH[-lunar.month]
	}else{
		s += LunarUtil.MONTH[lunar.month]
	}
	return s
}

func (lunar *Lunar) GetDayInChinese() string {
	return LunarUtil.DAY[lunar.day]
}

func (lunar *Lunar) GetSeason() string {
	m := lunar.month
	if m<0 {
		m = -m
	}
	return LunarUtil.SEASON[m]
}

func (lunar *Lunar) GetJie() string {
	j := len(LunarUtil.JIE)
	for i:=0;i<j;i++ {
		jie := LunarUtil.JIE[i]
		d := lunar.jieQi[jie]
		if d.year==lunar.solar.year&&d.month==lunar.solar.month&&d.day==lunar.solar.day {
			return jie
		}
	}
	return ""
}

func (lunar *Lunar) GetQi() string {
	j := len(LunarUtil.QI)
	for i:=0;i<j;i++ {
		qi := LunarUtil.QI[i]
		d := lunar.jieQi[qi]
		if d.year==lunar.solar.year&&d.month==lunar.solar.month&&d.day==lunar.solar.day {
			return qi
		}
	}
	return ""
}

func (lunar *Lunar) GetWeek() int {
	return lunar.weekIndex
}

func (lunar *Lunar) GetWeekInChinese() string {
	return SolarUtil.WEEK[lunar.GetWeek()]
}

func (lunar *Lunar) GetXiu() string {
	return LunarUtil.XIU[lunar.GetDayZhi()+strconv.Itoa(lunar.GetWeek())]
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

func (lunar *Lunar) GetFestivals() *list.List{
	l := list.New()
	if f, ok := LunarUtil.FESTIVAL[strconv.Itoa(lunar.month)+"-"+strconv.Itoa(lunar.day)]; ok {
		l.PushBack(f)
	}
	return l
}

func (lunar *Lunar) GetOtherFestivals() *list.List{
	l := list.New()
	if f, ok := LunarUtil.OTHER_FESTIVAL[strconv.Itoa(lunar.month)+"-"+strconv.Itoa(lunar.day)]; ok {
		for i:=0;i<len(f);i++ {
			l.PushBack(f[i])
		}
	}
	return l
}

func (lunar *Lunar) GetPengZuGan() string {
	return LunarUtil.PENGZU_GAN[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetPengZuZhi() string {
	return LunarUtil.PENGZU_ZHI[lunar.dayZhiIndex+1]
}

func (lunar *Lunar) GetPositionXi() string {
	return LunarUtil.POSITION_XI[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetPositionXiDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetPositionXi()]
}

func (lunar *Lunar) GetPositionYangGui() string {
	return LunarUtil.POSITION_YANG_GUI[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetPositionYangGuiDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetPositionYangGui()]
}

func (lunar *Lunar) GetPositionYinGui() string {
	return LunarUtil.POSITION_YIN_GUI[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetPositionYinGuiDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetPositionYinGui()]
}

func (lunar *Lunar) GetPositionFu() string {
	return LunarUtil.POSITION_FU[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetPositionFuDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetPositionFu()]
}

func (lunar *Lunar) GetPositionCai() string {
	return LunarUtil.POSITION_CAI[lunar.dayGanIndex+1]
}

func (lunar *Lunar) GetPositionCaiDesc() string {
	return LunarUtil.POSITION_DESC[lunar.GetPositionCai()]
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
	j := len(LunarUtil.ZHI)
	for i:=0;i<j;i++ {
		if LunarUtil.ZHI[i] == chong {
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
	return "("+lunar.GetDayChongGan()+lunar.GetDayChong()+")"+lunar.GetDayChongShengXiao()
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

func (lunar *Lunar) GetBaZi() [4]string {
	return [4]string{lunar.GetYearInGanZhiExact(),lunar.GetMonthInGanZhiExact(),lunar.GetDayInGanZhiExact(),lunar.GetTimeInGanZhi()}
}

func (lunar *Lunar) GetBaZiWuXing() [4]string {
	baZi := lunar.GetBaZi()
	j := len(baZi)
	l := [4]string{}
	for i:=0;i<j;i++ {
		ganZhi := baZi[i]
		gan := ganZhi[0:1]
		zhi := ganZhi[1:]
		l[i] = LunarUtil.WU_XING_GAN[gan]+LunarUtil.WU_XING_ZHI[zhi]
	}
	return l
}

func (lunar *Lunar) GetBaZiNaYin() [4]string {
	baZi := lunar.GetBaZi()
	j := len(baZi)
	l := [4]string{}
	for i:=0;i<j;i++ {
		ganZhi := baZi[i]
		l[i] = LunarUtil.NAYIN[ganZhi]
	}
	return l
}

func (lunar *Lunar) GetBaZiShiShenGan() [4]string {
	baZi := lunar.GetBaZi()
	yearGan := baZi[0][0:1]
	monthGan := baZi[1][0:1]
	dayGan := baZi[2][0:1]
	timeGan := baZi[3][0:1]
	l := [4]string{LunarUtil.SHI_SHEN_GAN[dayGan+yearGan], LunarUtil.SHI_SHEN_GAN[dayGan+monthGan], "日主", LunarUtil.SHI_SHEN_GAN[dayGan+timeGan]}
	return l
}

func (lunar *Lunar) GetBaZiShiShenZhi() [4]string {
	baZi := lunar.GetBaZi()
	dayGan := baZi[2][0:1]
	j := len(baZi)
	l := [4]string{}
	for i:=0;i<j;i++ {
		ganZhi := baZi[i]
		zhi := ganZhi[1:]
		l[i] = LunarUtil.SHI_SHEN_ZHI[dayGan+zhi+LunarUtil.ZHI_HIDE_GAN[zhi][0]]
	}
	return l
}

func (lunar *Lunar) GetZhiXing() string {
	offset := lunar.dayZhiIndex-lunar.monthZhiIndex
	if offset<0 {
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
	offset := lunar.dayGanIndex-lunar.dayZhiIndex
	if offset<0 {
		offset += 12
	}
	return LunarUtil.POSITION_TAI_DAY[offset*5+lunar.dayGanIndex]
}

func (lunar *Lunar) GetMonthPositionTai() string {
	if lunar.month<0 {
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
	j := len(LunarUtil.ZHI)
	for i:=0;i<j;i++ {
		if LunarUtil.ZHI[i] == chong {
			return LunarUtil.SHENG_XIAO[i]
		}
	}
	return ""
}

func (lunar *Lunar) GetTimeChongDesc() string {
	return "("+lunar.GetTimeChongGan()+lunar.GetTimeChong()+")"+lunar.GetTimeChongShengXiao()
}

func (lunar *Lunar) GetJieQiTable() map[string]*Solar {
	return lunar.jieQi
}

func (lunar *Lunar) String() string{
	return lunar.GetYearInChinese()+"年"+lunar.GetMonthInChinese()+"月"+lunar.GetDayInChinese()
}

func (lunar *Lunar) ToFullString() string{
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
	s +=lunar.GetDayInGanZhi()
	s +="("
	s +=lunar.GetDayShengXiao()
	s +=")日 "
	s +=lunar.GetTimeZhi()
	s +="("
	s +=lunar.GetTimeShengXiao()
	s +=")时 纳音["
	s +=lunar.GetYearNaYin()
	s +=" "
	s +=lunar.GetMonthNaYin()
	s +=" "
	s +=lunar.GetDayNaYin()
	s +=" "
	s +=lunar.GetTimeNaYin()
	s +="] 星期"
	s +=lunar.GetWeekInChinese()
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

	jq := lunar.GetJie()+lunar.GetQi()
	if len(jq)>0 {
		s +=" ["
		s +=jq
		s +="]"
	}
	s +=" "
	s +=lunar.GetGong()
	s +="方"
	s +=lunar.GetShou()
	s +=" 星宿["
	s +=lunar.GetXiu()
	s +=lunar.GetZheng()
	s +=lunar.GetAnimal()
	s +="]("
	s +=lunar.GetXiuLuck()
	s +=") 彭祖百忌["
	s +=lunar.GetPengZuGan()
	s +=" "
	s +=lunar.GetPengZuZhi()
	s +="] 喜神方位["
	s +=lunar.GetPositionXi()
	s +="]("
	s +=lunar.GetPositionXiDesc()
	s +=") 阳贵神方位["
	s +=lunar.GetPositionYangGui()
	s +="]("
	s +=lunar.GetPositionYangGuiDesc()
	s +=") 阴贵神方位["
	s +=lunar.GetPositionYinGui()
	s +="]("
	s +=lunar.GetPositionYinGuiDesc()
	s +=") 福神方位["
	s +=lunar.GetPositionFu()
	s +="]("
	s +=lunar.GetPositionFuDesc()
	s +=") 财神方位["
	s +=lunar.GetPositionCai()
	s +="]("
	s +=lunar.GetPositionCaiDesc()
	s +=") 冲["
	s +=lunar.GetDayChongDesc()
	s +="] 煞["
	s +=lunar.GetDaySha()
	s +="]"
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

func (lunar *Lunar) GetSolar() *Solar {
	return lunar.solar
}