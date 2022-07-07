package models

import (
	"gorm.io/gorm"
	"manageService/internal/app"
)

type AlarmRule struct {
	gorm.Model
	// cpu relative threshold
	CpuNoteworthyThreshold float32
	CpuSeriousThreshold float32
	CpuDeadlyThreshold float32

	// mem relative threshold
	MemNoteworthyThreshold float32
	MemSeriousThreshold float32
	MemDeadlyThreshold float32

	Granularity string
	AggregateFunction string
	ContactEmail string
}

func (rule AlarmRule) Insert() int {
	create := app.DB.Create(&rule)
	if create.Error != nil {
		return 0
	}
	return 1
}

func (rule AlarmRule) FindLatest() (int32, AlarmRule) {
	// 边缘条件需要加上去，如无法查询到记录则返回400的Code
	result := app.DB.Last(&rule)
	if result.Error != nil {
		return 400, AlarmRule{}
	}
	return 200, rule
}

func (rule AlarmRule) FindAll() []AlarmRule {
	var alarmRules []AlarmRule
	app.DB.Find(&alarmRules)
	return alarmRules
}