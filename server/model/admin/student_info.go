// 自动生成模板StudentInfo
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 学生信息 结构体  StudentInfo
type StudentInfo struct {
	global.GVA_MODEL
	Name              string     `json:"name" form:"name" gorm:"column:name;comment:;"binding:"required"`                                          //姓名
	Gender            string     `json:"gender" form:"gender" gorm:"column:gender;comment:;"binding:"required"`                                    //性别
	Birthday          *time.Time `json:"birthday" form:"birthday" gorm:"column:birthday;comment:;"binding:"required"`                              //出生日期
	IdCard            string     `json:"idCard" form:"idCard" gorm:"column:id_card;comment:;"binding:"required"`                                   //身份证
	AdmissionLetterId string     `json:"admissionLetterId" form:"admissionLetterId" gorm:"column:admission_letter_id;comment:;"binding:"required"` //录取通知书编号
	AdmissionYear     *int       `json:"admissionYear" form:"admissionYear" gorm:"column:admission_year;comment:;"binding:"required"`              //入学年份
	OriginPlace       string     `json:"originPlace" form:"originPlace" gorm:"column:origin_place;comment:;"binding:"required"`                    //生源地
	Campus            string     `json:"campus" form:"campus" gorm:"column:campus;comment:;"binding:"required"`                                    //校区
	StudentId         string     `json:"studentId" form:"studentId" gorm:"column:student_id;comment:;"binding:"required"`                          //学号
	College           string     `json:"college" form:"college" gorm:"column:college;comment:;"binding:"required"`                                 //学院
	Major             string     `json:"major" form:"major" gorm:"column:major;comment:;"binding:"required"`                                       //专业
	Class             string     `json:"class" form:"class" gorm:"column:class;comment:;"binding:"required"`                                       //班级
	Dormitory         string     `json:"dormitory" form:"dormitory" gorm:"column:dormitory;comment:;"binding:"required"`                           //寝室
	Portrait          string     `json:"portrait" form:"portrait" gorm:"column:portrait;comment:;"binding:"required"`                              //人像
	UserId            *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:;"binding:"required"`                                   //用户ID
}

// TableName 学生信息 StudentInfo自定义表名 student_infos
func (StudentInfo) TableName() string {
	return "student_infos"
}
