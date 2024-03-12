// 自动生成模板Ad
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
	"time"
)

// 广告 结构体  Ad
type Ad struct {
	global.GVA_MODEL
	AdvertiserId    *int           `json:"advertiserId" form:"advertiserId" gorm:"column:advertiser_id;comment:;"`                   //所属广告主
	Name            string         `json:"name" form:"name" gorm:"column:name;comment:;"binding:"required"`                          //名称
	Text            string         `json:"text" form:"text" gorm:"column:text;comment:;"binding:"required"`                          //文本
	Img             string         `json:"img" form:"img" gorm:"column:img;comment:;"binding:"required"`                             //图片
	Video           datatypes.JSON `json:"video" form:"video" gorm:"column:video;comment:;"binding:"required"`                       //视频
	Link            string         `json:"link" form:"link" gorm:"column:link;comment:;"binding:"required"`                          //跳转链接
	StartDate       *time.Time     `json:"startDate" form:"startDate" gorm:"column:start_date;comment:;"binding:"required"`          //开始日期
	EndDate         *time.Time     `json:"endDate" form:"endDate" gorm:"column:end_date;comment:;"binding:"required"`                //结束日期
	StartHour       *int           `json:"startHour" form:"startHour" gorm:"column:start_hour;comment:;"binding:"required"`          //开始时段
	EndHour         *int           `json:"endHour" form:"endHour" gorm:"column:end_hour;comment:;"binding:"required"`                //结束时段
	Target          datatypes.JSON `json:"target" form:"target" gorm:"column:target;comment:;"binding:"required"`                    //目标受众
	Keywords        datatypes.JSON `json:"keywords" form:"keywords" gorm:"column:keywords;comment:;"binding:"required"`              //关键词
	CostCategory    *int           `json:"costCategory" form:"costCategory" gorm:"column:cost_category;comment:;"binding:"required"` //计费类型
	BugAmount       *int           `json:"bugAmount" form:"bugAmount" gorm:"column:bug_amount;comment:;"binding:"required"`          //购买量
	CostImpressions *int           `json:"costImpressions" form:"costImpressions" gorm:"column:cost_impressions;comment:;"`          //已展示量
	CostClicks      *int           `json:"costClicks" form:"costClicks" gorm:"column:cost_clicks;comment:;"`                         //已点击数
	Cost            *int           `json:"cost" form:"cost" gorm:"column:cost;comment:;"binding:"required"`                          //花费费用
	Status          *int           `json:"status" form:"status" gorm:"column:status;comment:;"binding:"required"`                    //状态
}

// TableName 广告 Ad自定义表名 ads
func (Ad) TableName() string {
	return "ads"
}
