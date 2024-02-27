// 自动生成模板CommentScore
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 评分 结构体  CommentScore
type CommentScore struct {
	global.GVA_MODEL
	Category *bool    `json:"category" form:"category" gorm:"column:category;comment:;"`  //种类
	TargetId *int     `json:"targetId" form:"targetId" gorm:"column:target_id;comment:;"` //评论对象ID
	One      *int     `json:"one" form:"one" gorm:"column:one;comment:;"`                 //一星数量
	Two      *int     `json:"two" form:"two" gorm:"column:two;comment:;"`                 //二星数量
	Three    *int     `json:"three" form:"three" gorm:"column:three;comment:;"`           //三星数量
	Four     *int     `json:"four" form:"four" gorm:"column:four;comment:;"`              //四星数量
	Five     *int     `json:"five" form:"five" gorm:"column:five;comment:;"`              //五星数量
	Score    *float64 `json:"score" form:"score" gorm:"column:score;comment:;"`           //评分
}

// TableName 评分 CommentScore自定义表名 comment_scores
func (CommentScore) TableName() string {
	return "comment_scores"
}
