package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
    adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type QuestionnaireService struct {
}

// CreateQuestionnaire 创建寝室分配问卷记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionnaireService *QuestionnaireService) CreateQuestionnaire(questionnaire *admin.Questionnaire) (err error) {
	err = global.GVA_DB.Create(questionnaire).Error
	return err
}

// DeleteQuestionnaire 删除寝室分配问卷记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionnaireService *QuestionnaireService)DeleteQuestionnaire(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.Questionnaire{},"id = ?",ID).Error
	return err
}

// DeleteQuestionnaireByIds 批量删除寝室分配问卷记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionnaireService *QuestionnaireService)DeleteQuestionnaireByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.Questionnaire{},"id in ?",IDs).Error
	return err
}

// UpdateQuestionnaire 更新寝室分配问卷记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionnaireService *QuestionnaireService)UpdateQuestionnaire(questionnaire admin.Questionnaire) (err error) {
	err = global.GVA_DB.Save(&questionnaire).Error
	return err
}

// GetQuestionnaire 根据ID获取寝室分配问卷记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionnaireService *QuestionnaireService)GetQuestionnaire(ID string) (questionnaire admin.Questionnaire, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&questionnaire).Error
	return
}

// GetQuestionnaireInfoList 分页获取寝室分配问卷记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionnaireService *QuestionnaireService)GetQuestionnaireInfoList(info adminReq.QuestionnaireSearch) (list []admin.Questionnaire, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&admin.Questionnaire{})
    var questionnaires []admin.Questionnaire
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.StudentInfoId != nil {
        db = db.Where("student_info_id = ?",info.StudentInfoId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&questionnaires).Error
	return  questionnaires, total, err
}
