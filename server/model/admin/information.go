// 自动生成模板Information
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"

	"gorm.io/datatypes"
)

// 信息 结构体  Information
type Information struct {
	global.GVA_MODEL
	Category uint8          `json:"category" form:"category" gorm:"column:category;comment:;"binding:"required"` //信息类别
	Title    string         `json:"title" form:"title" gorm:"column:title;comment:;"binding:"required"`          //标题
	Source   string         `json:"source" form:"source" gorm:"column:source;comment:;"binding:"required"`       //来源
	Content  string         `json:"content" form:"content" gorm:"column:content;comment:;"binding:"required"`    //内容
	Imgs     datatypes.JSON `json:"imgs" form:"imgs" gorm:"column:imgs;comment:;"`                               //图片
}

// TableName 信息 Information自定义表名 information
func (Information) TableName() string {
	return "information"
}
