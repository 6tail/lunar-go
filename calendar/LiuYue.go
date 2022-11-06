package calendar

import (
	"github.com/6tail/lunar-go/LunarUtil"
	"strings"
)

// LiuYue 流月
type LiuYue struct {
	// 序数，0-9
	index   int
	liuNian *LiuNian
}

func NewLiuYue(liuNian *LiuNian, index int) *LiuYue {
	liuYue := new(LiuYue)
	liuYue.liuNian = liuNian
	liuYue.index = index
	return liuYue
}

func (liuYue *LiuYue) GetIndex() int {
	return liuYue.index
}

// GetMonthInChinese 获取中文的月
func (liuYue *LiuYue) GetMonthInChinese() string {
	return LunarUtil.MONTH[liuYue.index+1]
}

// GetGanZhi 获取干支
// <p>
// 《五虎遁》
// 甲己之年丙作首，
// 乙庚之年戊为头，
// 丙辛之年寻庚上，
// 丁壬壬寅顺水流，
// 若问戊癸何处走，
// 甲寅之上好追求。
func (liuYue *LiuYue) GetGanZhi() string {
	offset := 0
	ganZhi := liuYue.liuNian.GetGanZhi()
	gz := []rune(ganZhi)
	yearGan := string(gz[:1])
	if strings.Compare("甲", yearGan) == 0 || strings.Compare("己", yearGan) == 0 {
		offset = 2
	}
	gan := LunarUtil.GAN[(liuYue.index+offset)%10+1]
	zhi := LunarUtil.ZHI[(liuYue.index+LunarUtil.BASE_MONTH_ZHI_INDEX)%12+1]
	return gan + zhi
}

// GetXun 获取所在旬
func (liuYue *LiuYue) GetXun() string {
	return LunarUtil.GetXun(liuYue.GetGanZhi())
}

// GetXunKong 获取旬空(空亡)
func (liuYue *LiuYue) GetXunKong() string {
	return LunarUtil.GetXunKong(liuYue.GetGanZhi())
}
