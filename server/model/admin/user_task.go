// 自动生成模板UserTask
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 用户任务记录 结构体  UserTask
type UserTask struct {
	global.GVA_MODEL
	UserId   *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`       //userId字段
	TaskId   *int   `json:"taskId" form:"taskId" gorm:"column:task_id;comment:;"`       //taskId字段
	Pic      string `json:"pic" form:"pic" gorm:"column:pic;comment:;"`                 //pic字段
	Loc      string `json:"loc" form:"loc" gorm:"column:loc;comment:;"`                 //loc字段
	CurStage *bool  `json:"curStage" form:"curStage" gorm:"column:cur_stage;comment:;"` //curStage字段
}

// TableName 用户任务记录 UserTask自定义表名 user_tasks
func (UserTask) TableName() string {
	return "user_tasks"
}
