// 自动生成模板Organization
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 组织社团 结构体  Organization
type Organization struct {
	global.GVA_MODEL
	SysUserId    uint   `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:;"`                          //关联的后台管理用户
	Introduction string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:;"binding:"required"` //介绍
	Name         string `json:"name" form:"name" gorm:"column:name;comment:;"binding:"required"`                         //名称
	Pic          string `json:"pic" form:"pic" gorm:"column:pic;comment:;"binding:"required"`                            //封面图
}

// TableName 组织社团 Organization自定义表名 organizations
func (Organization) TableName() string {
	return "organizations"
}
