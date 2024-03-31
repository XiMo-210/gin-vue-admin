package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
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
func (reviewRecordService *ReviewRecordsService) GetReviewRecords(ID string) (reviewRecord admin.ReviewRecords, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&reviewRecord).Error
	return
}

// GetReviewRecordsInfoList 分页获取审核记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (reviewRecordService *ReviewRecordsService) GetReviewRecordsInfoList(info adminReq.ReviewRecordsSearch) (list []admin.ReviewRecords, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.ReviewRecords{})
	var reviewRecords []admin.ReviewRecords
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
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

	err = db.Find(&reviewRecords).Error
	return reviewRecords, total, err
}
