// 自动生成模板ReviewRecords
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 审核记录 结构体  ReviewRecords
type ReviewRecords struct {
	global.GVA_MODEL
	UserTaskId *int   `json:"userTaskId" form:"userTaskId" gorm:"column:user_task_id;comment:;"` //用户任务记录ID
	Stage      *int   `json:"stage" form:"stage" gorm:"column:stage;comment:;"`                  //阶段
	Loc        string `json:"loc" form:"loc" gorm:"column:loc;comment:;"`                        //位置
	Pic        string `json:"pic" form:"pic" gorm:"column:pic;comment:;"`                        //图片
	Reply      string `json:"reply" form:"reply" gorm:"column:reply;comment:;"`                  //回应
	Status     *int   `json:"status" form:"status" gorm:"column:status;comment:;"`               //状态
}

// TableName 审核记录 ReviewRecords自定义表名 review_records
func (ReviewRecords) TableName() string {
	return "review_records"
}
