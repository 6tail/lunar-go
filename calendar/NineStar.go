package calendar

import "github.com/6tail/lunar-go/LunarUtil"

// NUMBER 九数
var NUMBER = []string{"一", "二", "三", "四", "五", "六", "七", "八", "九"}

// COLOR 七色
var COLOR = []string{"白", "黒", "碧", "绿", "黄", "白", "赤", "白", "紫"}

// WU_XING 五行
var WU_XING = []string{"水", "土", "木", "木", "土", "金", "金", "土", "火"}

// POSITION 后天八卦方位
var POSITION = []string{"坎", "坤", "震", "巽", "中", "乾", "兑", "艮", "离"}

// NAME_BEI_DOU 北斗九星
var NAME_BEI_DOU = []string{"天枢", "天璇", "天玑", "天权", "玉衡", "开阳", "摇光", "洞明", "隐元"}

// NAME_XUAN_KONG 玄空九星（玄空风水）
var NAME_XUAN_KONG = []string{"贪狼", "巨门", "禄存", "文曲", "廉贞", "武曲", "破军", "左辅", "右弼"}

// NAME_QI_MEN 奇门九星（奇门遁甲，也称天盘九星）
var NAME_QI_MEN = []string{"天蓬", "天芮", "天冲", "天辅", "天禽", "天心", "天柱", "天任", "天英"}

// BA_MEN_QI_MEN 八门（奇门遁甲）
var BA_MEN_QI_MEN = []string{"休", "死", "伤", "杜", "", "开", "惊", "生", "景"}

// NAME_TAI_YI 太乙九神（太乙神数）
var NAME_TAI_YI = []string{"太乙", "摄提", "轩辕", "招摇", "天符", "青龙", "咸池", "太阴", "天乙"}

// TYPE_TAI_YI 太乙九神对应类型
var TYPE_TAI_YI = []string{"吉神", "凶神", "安神", "安神", "凶神", "吉神", "凶神", "吉神", "吉神"}

// SONG_TAI_YI 太乙九神歌诀（太乙神数）
var SONG_TAI_YI = []string{"门中太乙明，星官号贪狼，赌彩财喜旺，婚姻大吉昌，出入无阻挡，参谒见贤良，此行三五里，黑衣别阴阳。", "门前见摄提，百事必忧疑，相生犹自可，相克祸必临，死门并相会，老妇哭悲啼，求谋并吉事，尽皆不相宜，只可藏隐遁，若动伤身疾。", "出入会轩辕，凡事必缠牵，相生全不美，相克更忧煎，远行多不利，博彩尽输钱，九天玄女法，句句不虚言。", "招摇号木星，当之事莫行，相克行人阻，阴人口舌迎，梦寐多惊惧，屋响斧自鸣，阴阳消息理，万法弗违情。", "五鬼为天符，当门阴女谋，相克无好事，行路阻中途，走失难寻觅，道逢有尼姑，此星当门值，万事有灾除。", "神光跃青龙，财气喜重重，投入有酒食，赌彩最兴隆，更逢相生旺，休言克破凶，见贵安营寨，万事总吉同。", "吾将为咸池，当之尽不宜，出入多不利，相克有灾情，赌彩全输尽，求财空手回，仙人真妙语，愚人莫与知，动用虚惊退，反复逆风吹。", "坐临太阴星，百祸不相侵，求谋悉成就，知交有觅寻，回风归来路，恐有殃伏起，密语中记取，慎乎莫轻行。", "迎来天乙星，相逢百事兴，运用和合庆，茶酒喜相迎，求谋并嫁娶，好合有天成，祸福如神验，吉凶甚分明。"}

// LUCK_XUAN_KONG 吉凶（玄空风水）
var LUCK_XUAN_KONG = []string{"吉", "凶", "凶", "吉", "凶", "吉", "凶", "吉", "吉"}

// LUCK_QI_MEN 吉凶（奇门遁甲）
var LUCK_QI_MEN = []string{"大凶", "大凶", "小吉", "大吉", "大吉", "大吉", "小凶", "小吉", "小凶"}

// YIN_YANG_QI_MEN 阴阳（奇门遁甲）
var YIN_YANG_QI_MEN = []string{"阳", "阴", "阳", "阳", "阳", "阴", "阴", "阳", "阴"}

// NineStar 九星
type NineStar struct {
	index int
}

func NewNineStar(index int) *NineStar {
	nineStar := new(NineStar)
	nineStar.index = index
	return nineStar
}

func (nineStar *NineStar) GetNumber() string {
	return NUMBER[nineStar.index]
}

func (nineStar *NineStar) GetColor() string {
	return COLOR[nineStar.index]
}

func (nineStar *NineStar) GetWuXing() string {
	return WU_XING[nineStar.index]
}

func (nineStar *NineStar) GetPosition() string {
	return POSITION[nineStar.index]
}

func (nineStar *NineStar) GetPositionDesc() string {
	return LunarUtil.POSITION_DESC[nineStar.GetPosition()]
}

func (nineStar *NineStar) GetNameInXuanKong() string {
	return NAME_XUAN_KONG[nineStar.index]
}

func (nineStar *NineStar) GetNameInBeiDou() string {
	return NAME_BEI_DOU[nineStar.index]
}

func (nineStar *NineStar) GetNameInQiMen() string {
	return NAME_QI_MEN[nineStar.index]
}

func (nineStar *NineStar) GetNameInTaiYi() string {
	return NAME_TAI_YI[nineStar.index]
}

func (nineStar *NineStar) GetLuckInQiMen() string {
	return LUCK_QI_MEN[nineStar.index]
}

func (nineStar *NineStar) GetLuckInXuanKong() string {
	return LUCK_XUAN_KONG[nineStar.index]
}

func (nineStar *NineStar) GetYinYangInQiMen() string {
	return YIN_YANG_QI_MEN[nineStar.index]
}

func (nineStar *NineStar) GetTypeInTaiYi() string {
	return TYPE_TAI_YI[nineStar.index]
}

func (nineStar *NineStar) GetBaMenInQiMen() string {
	return BA_MEN_QI_MEN[nineStar.index]
}

func (nineStar *NineStar) GetSongInTaiYi() string {
	return SONG_TAI_YI[nineStar.index]
}

func (nineStar *NineStar) GetIndex() int {
	return nineStar.index
}

func (nineStar *NineStar) String() string {
	return nineStar.GetNumber() + nineStar.GetColor() + nineStar.GetWuXing() + nineStar.GetNameInBeiDou()
}

func (nineStar *NineStar) ToFullString() string {
	s := nineStar.GetNumber()
	s += nineStar.GetColor()
	s += nineStar.GetWuXing()
	s += " "
	s += nineStar.GetPosition()
	s += "("
	s += nineStar.GetPositionDesc()
	s += ") "
	s += nineStar.GetNameInBeiDou()
	s += " 玄空["
	s += nineStar.GetNameInXuanKong()
	s += " "
	s += nineStar.GetLuckInXuanKong()
	s += "] 奇门["
	s += nineStar.GetNameInQiMen()
	s += " "
	s += nineStar.GetLuckInQiMen()
	if len(nineStar.GetBaMenInQiMen()) > 0 {
		s += " "
		s += nineStar.GetBaMenInQiMen()
		s += "门"
	}
	s += " "
	s += nineStar.GetYinYangInQiMen()
	s += "] 太乙["
	s += nineStar.GetNameInTaiYi()
	s += " "
	s += nineStar.GetTypeInTaiYi()
	s += "]"
	return s
}
