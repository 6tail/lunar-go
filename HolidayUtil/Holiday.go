package HolidayUtil

import "strings"

// Holiday 节假日
type Holiday struct {
	// 日期，YYYY-MM-DD格式
	day string
	// 名称，如：国庆
	name string
	// 是否调休，即是否要上班
	work bool
	// 关联的节日，YYYY-MM-DD格式
	target string
}

func NewHoliday(day string, name string, work bool, target string) *Holiday {
	holiday := new(Holiday)
	holiday.SetDay(day)
	holiday.SetName(name)
	holiday.SetWork(work)
	holiday.SetTarget(target)
	return holiday
}

func (holiday *Holiday) SetDay(day string) {
	if !strings.Contains(day, "-") {
		holiday.day = day[0:4] + "-" + day[4:6] + "-" + day[6:]
	} else {
		holiday.day = day
	}
}

func (holiday *Holiday) SetName(name string) {
	holiday.name = name
}

func (holiday *Holiday) SetWork(work bool) {
	holiday.work = work
}

func (holiday *Holiday) SetTarget(target string) {
	if !strings.Contains(target, "-") {
		holiday.target = target[0:4] + "-" + target[4:6] + "-" + target[6:]
	} else {
		holiday.target = target
	}
}

func (holiday *Holiday) GetDay() string {
	return holiday.day
}

func (holiday *Holiday) GetName() string {
	return holiday.name
}

func (holiday *Holiday) IsWork() bool {
	return holiday.work
}

func (holiday *Holiday) GetTarget() string {
	return holiday.target
}

func (holiday *Holiday) String() string {
	if nil == holiday {
		return ""
	}
	workDesc := ""
	if holiday.work {
		workDesc = "调休"
	}
	return holiday.day + " " + holiday.name + (workDesc) + " " + holiday.target
}
