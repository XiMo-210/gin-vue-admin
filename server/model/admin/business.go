// 自动生成模板Business
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 店铺商家 结构体  Business
type Business struct {
	global.GVA_MODEL
	SysUserId     uint   `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:;"`                              //关联的后台管理用户
	Name          string `json:"name" form:"name" gorm:"column:name;comment:;"binding:"required"`                             //名称
	Pic           string `json:"pic" form:"pic" gorm:"column:pic;comment:;"binding:"required"`                                //封面图
	Address       string `json:"address" form:"address" gorm:"column:address;comment:;"binding:"required"`                    //地址
	Loc           string `json:"loc" form:"loc" gorm:"column:loc;comment:;"binding:"required"`                                //导航地址
	BusinessHours string `json:"businessHours" form:"businessHours" gorm:"column:business_hours;comment:;"binding:"required"` //营业时间
	Contact       string `json:"contact" form:"contact" gorm:"column:contact;comment:;"binding:"required"`                    //联系方式
	Introduction  string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:;"binding:"required"`     //介绍
}

// TableName 店铺商家 Business自定义表名 businesses
func (Business) TableName() string {
	return "businesses"
}
