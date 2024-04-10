// 自动生成模板PlatformParam
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 平台参数 结构体  PlatformParam
type PlatformParam struct {
	global.GVA_MODEL
	SchoolName        string     `json:"schoolName" form:"schoolName" gorm:"column:school_name;comment:;"binding:"required"`                       //学校名称
	SchoolEmblem      string     `json:"schoolEmblem" form:"schoolEmblem" gorm:"column:school_emblem;comment:;"binding:"required"`                 //校徽
	AdmissionYear     *int       `json:"admissionYear" form:"admissionYear" gorm:"column:admission_year;comment:;"binding:"required"`              //入学年份
	RegisterStartDate *time.Time `json:"registerStartDate" form:"registerStartDate" gorm:"column:register_start_date;comment:;"binding:"required"` //报到开始日
	RegisterEndDate   *time.Time `json:"registerEndDate" form:"registerEndDate" gorm:"column:register_end_date;comment:;"binding:"required"`       //报到结束日
	TermYear          *int       `json:"termYear" form:"termYear" gorm:"column:term_year;comment:;"binding:"required"`                             //当前学年
	Term              string     `json:"term" form:"term" gorm:"column:term;comment:;"binding:"required"`                                          //当前学期
	TermStartDate     *time.Time `json:"termStartDate" form:"termStartDate" gorm:"column:term_start_date;comment:;"binding:"required"`             //学期开始时间
	TermEndDate       *time.Time `json:"termEndDate" form:"termEndDate" gorm:"column:term_end_date;comment:;"binding:"required"`                   //学期结束时间
}

// TableName 平台参数 PlatformParam自定义表名 platform_params
func (PlatformParam) TableName() string {
	return "platform_params"
}
