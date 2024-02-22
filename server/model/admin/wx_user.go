// 自动生成模板WxUser
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 用户 结构体  WxUser
type WxUser struct {
	global.GVA_MODEL
	Openid          string `json:"openid" form:"openid" gorm:"column:openid;comment:;"`                              //微信用户唯一标识
	StudentInfoId   *int   `json:"studentInfoId" form:"studentInfoId" gorm:"column:student_info_id;comment:;"`       //关联学生信息
	FaceId          string `json:"faceId" form:"faceId" gorm:"column:face_id;comment:;"`                             //人脸识别标识
	Username        string `json:"username" form:"username" gorm:"column:username;comment:;"`                        //用户名
	Password        string `json:"password" form:"password" gorm:"column:password;comment:;"`                        //密码
	Avatar          string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:;"`                              //人像
	Signature       string `json:"signature" form:"signature" gorm:"column:signature;comment:;"`                     //个性签名
	Hobby           string `json:"hobby" form:"hobby" gorm:"column:hobby;comment:;"`                                 //爱好
	Contact         string `json:"contact" form:"contact" gorm:"column:contact;comment:;"`                           //联系方式
	Phone           string `json:"phone" form:"phone" gorm:"column:phone;comment:;"`                                 //手机
	Email           string `json:"email" form:"email" gorm:"column:email;comment:;"`                                 //邮箱
	Exp             *int   `json:"exp" form:"exp" gorm:"column:exp;comment:;"`                                       //智慧种子
	Points          *int   `json:"points" form:"points" gorm:"column:points;comment:;"`                              //求索石
	IsCompletedMain *bool  `json:"isCompletedMain" form:"isCompletedMain" gorm:"column:is_completed_main;comment:;"` //是否完成主线任务
	CurTask         *int   `json:"curTask" form:"curTask" gorm:"column:cur_task;comment:;"`                          //正在进行的任务
}

// TableName 用户 WxUser自定义表名 wx_users
func (WxUser) TableName() string {
	return "wx_users"
}
