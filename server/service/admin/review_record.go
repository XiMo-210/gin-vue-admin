package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin/response"
)

type ReviewRecordsService struct {
}

// CreateReviewRecords 创建审核记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (reviewRecordService *ReviewRecordsService) CreateReviewRecords(reviewRecord *admin.ReviewRecords) (err error) {
	err = global.GVA_DB.Create(reviewRecord).Error
	return err
}

// DeleteReviewRecords 删除审核记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (reviewRecordService *ReviewRecordsService) DeleteReviewRecords(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.ReviewRecords{}, "id = ?", ID).Error
	return err
}

// DeleteReviewRecordsByIds 批量删除审核记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (reviewRecordService *ReviewRecordsService) DeleteReviewRecordsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.ReviewRecords{}, "id in ?", IDs).Error
	return err
}

// UpdateReviewRecords 更新审核记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (reviewRecordService *ReviewRecordsService) UpdateReviewRecords(reviewRecord admin.ReviewRecords) (err error) {
	err = global.GVA_DB.Save(&reviewRecord).Error
	return err
}

// GetReviewRecords 根据ID获取审核记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (reviewRecordService *ReviewRecordsService) GetReviewRecords(ID string) (reviewRecord response.FindReviewRecords, err error) {
	err = global.GVA_DB.Model(&admin.ReviewRecords{}).
		Select("review_records.*, wx_users.username, wx_users.face_id, tasks.category, tasks.title AS task_title, tasks.desc AS task_desc, tasks.campus, tasks.college, tasks.start_time, tasks.end_time, task_stages.title AS task_stage_title, task_stages.desc AS task_stage_desc, task_stages.required_item, task_stages.imgs, task_stages.need_face, task_stages.pic AS task_stage_pic").
		Joins("JOIN user_tasks ON review_records.user_task_id = user_tasks.id").
		Joins("JOIN wx_users ON wx_users.id = user_tasks.user_id").
		Joins("JOIN tasks ON tasks.id = user_tasks.task_id").
		Joins("JOIN task_stages ON task_stages.task_id = tasks.id AND task_stages.stage = review_records.stage AND task_stages.deleted_at IS NOT NULL").
		Where("review_records.id = ?", ID).First(&reviewRecord).Error

	return
}

// GetReviewRecordsInfoList 分页获取审核记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (reviewRecordService *ReviewRecordsService) GetReviewRecordsInfoList(info adminReq.ReviewRecordsSearch) (list []response.GetReviewRecordsList, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.ReviewRecords{}).
		Select("review_records.*, wx_users.username, tasks.category, tasks.title AS task_title, tasks.end_time, task_stages.title AS task_stage_title").
		Joins("JOIN user_tasks ON review_records.user_task_id = user_tasks.id").
		Joins("JOIN wx_users ON wx_users.id = user_tasks.user_id").
		Joins("JOIN tasks ON tasks.id = user_tasks.task_id").
		Joins("JOIN task_stages ON task_stages.task_id = tasks.id AND task_stages.stage = review_records.stage AND task_stages.deleted_at IS NULL")

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Username != "" {
		db = db.Where("username = ?", info.Username)
	}
	if info.Category != nil {
		db = db.Where("category = ?", info.Category)
	}
	if info.TaskTitle != "" {
		db = db.Where("tasks.title = ?", info.TaskTitle)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&list).Error
	return list, total, err
}
