// 自动生成模板Questionnaire
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 寝室分配问卷 结构体  Questionnaire
type Questionnaire struct {
 global.GVA_MODEL 
      StudentInfoId  *int `json:"studentInfoId" form:"studentInfoId" gorm:"column:student_info_id;comment:;"binding:"required"`  //studentInfoId字段 
      Param1  *int `json:"param1" form:"param1" gorm:"column:param1;comment:;"`  //参数1 
      Param2  *int `json:"param2" form:"param2" gorm:"column:param2;comment:;"`  //参数2 
      Param3  *int `json:"param3" form:"param3" gorm:"column:param3;comment:;"`  //参数3 
      Param4  *int `json:"param4" form:"param4" gorm:"column:param4;comment:;"`  //参数4 
      Param5  *int `json:"param5" form:"param5" gorm:"column:param5;comment:;"`  //参数5 
      Param6  *int `json:"param6" form:"param6" gorm:"column:param6;comment:;"`  //参数6 
      Param7  *int `json:"param7" form:"param7" gorm:"column:param7;comment:;"`  //参数7 
      Param8  *int `json:"param8" form:"param8" gorm:"column:param8;comment:;"`  //参数8 
      Param9  *int `json:"param9" form:"param9" gorm:"column:param9;comment:;"`  //参数9 
      Param10  *int `json:"param10" form:"param10" gorm:"column:param10;comment:;"`  //参数10 
}


// TableName 寝室分配问卷 Questionnaire自定义表名 questionnaires
func (Questionnaire) TableName() string {
  return "questionnaires"
}

