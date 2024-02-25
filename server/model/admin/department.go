// 自动生成模板Department
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 部门 结构体  Department
type Department struct {
	global.GVA_MODEL
	OrganizationId *int   `json:"organizationId" form:"organizationId" gorm:"column:organization_id;comment:;"binding:"required"` //所属组织社团
	Name           string `json:"name" form:"name" gorm:"column:name;comment:;"binding:"required"`                                //名称
	Desc           string `json:"desc" form:"desc" gorm:"column:desc;comment:;"binding:"required"`                                //描述
}

// TableName 部门 Department自定义表名 departments
func (Department) TableName() string {
	return "departments"
}
