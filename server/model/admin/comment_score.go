// 自动生成模板CommentScore
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 评分 结构体  CommentScore
type CommentScore struct {
	global.GVA_MODEL
	Category uint8   `json:"category" form:"category" gorm:"column:category;"`  //种类
	TargetId uint    `json:"targetId" form:"targetId" gorm:"column:target_id;"` //评论对象ID
	One      int     `json:"one" form:"one" gorm:"column:one;default:0;"`       //一星数量
	Two      int     `json:"two" form:"two" gorm:"column:two;default:0;"`       //二星数量
	Three    int     `json:"three" form:"three" gorm:"column:three;default:0;"` //三星数量
	Four     int     `json:"four" form:"four" gorm:"column:four;default:0;"`    //四星数量
	Five     int     `json:"five" form:"five" gorm:"column:five;default:0;"`    //五星数量
	Score    float64 `json:"score" form:"score" gorm:"column:score;default:0;"` //评分
}

// TableName 评分 CommentScore自定义表名 comment_scores
func (CommentScore) TableName() string {
	return "comment_scores"
}

const (
	TASK_COMMENT = iota + 1
	ORGANIZATION_COMMENT
	BUSINESS_COMMENT
)
