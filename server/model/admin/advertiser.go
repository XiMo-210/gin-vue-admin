// 自动生成模板Advertiser
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 广告主 结构体  Advertiser
type Advertiser struct {
	global.GVA_MODEL
	SysUserId    uint   `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:;"`                          //关联的后台管理用户
	Name         string `json:"name" form:"name" gorm:"column:name;comment:;"binding:"required"`                         //名称
	Logo         string `json:"logo" form:"logo" gorm:"column:logo;comment:;"binding:"required"`                         //标志
	Introduction string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:;"binding:"required"` //介绍
	SubjectInfo  string `json:"subjectInfo" form:"subjectInfo" gorm:"column:subject_info;comment:;"binding:"required"`   //主体信息
	License      string `json:"license" form:"license" gorm:"column:license;comment:;"binding:"required"`                //营业执照注册号
	Location     string `json:"location" form:"location" gorm:"column:location;comment:;"binding:"required"`             //所在省市
	Points       uint   `json:"points" form:"points" gorm:"column:points;comment:;"binding:"required"`                   //积分
}

// TableName 广告主 Advertiser自定义表名 advertisers
func (Advertiser) TableName() string {
	return "advertisers"
}
