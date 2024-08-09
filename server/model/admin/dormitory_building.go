// 自动生成模板DormitoryBuilding
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 寝室楼 结构体  DormitoryBuilding
type DormitoryBuilding struct {
	global.GVA_MODEL
	Campus         string `json:"campus" form:"campus" gorm:"column:campus;comment:;"binding:"required"`                          //所属校区
	Name           string `json:"name" form:"name" gorm:"column:name;comment:;"binding:"required"`                                //寝室楼名称
	Gender         string `json:"gender" form:"gender" gorm:"column:gender;comment:;"binding:"required"`                          //居住性别
	ManagerName    string `json:"managerName" form:"managerName" gorm:"column:manager_name;comment:;"binding:"required"`          //管理员性别
	ManagerContact string `json:"managerContact" form:"managerContact" gorm:"column:manager_contact;comment:;"binding:"required"` //管理员联系方式
}

// TableName 寝室楼 DormitoryBuilding自定义表名 dormitory_buildings
func (DormitoryBuilding) TableName() string {
	return "dormitory_buildings"
}
