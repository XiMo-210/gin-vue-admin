// 自动生成模板Task
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 任务 结构体  Task
type Task struct {
	global.GVA_MODEL
	Category  *int       `json:"category" form:"category" gorm:"column:category;comment:;"binding:"required"`     //任务类型
	Title     string     `json:"title" form:"title" gorm:"column:title;comment:;"binding:"required"`              //任务名称
	Desc      string     `json:"desc" form:"desc" gorm:"column:desc;comment:;"binding:"required"`                 //任务描述
	Campus    string     `json:"campus" form:"campus" gorm:"column:campus;comment:;"binding:"required"`           //任务校区
	College   string     `json:"college" form:"college" gorm:"column:college;comment:;"binding:"required"`        //任务学院
	Reward    *int       `json:"reward" form:"reward" gorm:"column:reward;comment:;"binding:"required"`           //奖励积分
	NeedMain  *bool      `json:"needMain" form:"needMain" gorm:"column:need_main;comment:;"binding:"required"`    //是否需要完成主线任务
	StartTime *time.Time `json:"startTime" form:"startTime" gorm:"column:start_time;comment:;"binding:"required"` //开始时间
	EndTime   *time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:;"binding:"required"`       //结束时间
}

// TableName 任务 Task自定义表名 tasks
func (Task) TableName() string {
	return "tasks"
}
