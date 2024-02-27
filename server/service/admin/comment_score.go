package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type CommentScoreService struct {
}

// CreateCommentScore 创建评分记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentScoreService *CommentScoreService) CreateCommentScore(commentScore *admin.CommentScore) (err error) {
	err = global.GVA_DB.Create(commentScore).Error
	return err
}

// DeleteCommentScore 删除评分记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentScoreService *CommentScoreService) DeleteCommentScore(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.CommentScore{}, "id = ?", ID).Error
	return err
}

// DeleteCommentScoreByIds 批量删除评分记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentScoreService *CommentScoreService) DeleteCommentScoreByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.CommentScore{}, "id in ?", IDs).Error
	return err
}

// UpdateCommentScore 更新评分记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentScoreService *CommentScoreService) UpdateCommentScore(commentScore admin.CommentScore) (err error) {
	err = global.GVA_DB.Save(&commentScore).Error
	return err
}

// GetCommentScore 根据ID获取评分记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentScoreService *CommentScoreService) GetCommentScore(ID string) (commentScore admin.CommentScore, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&commentScore).Error
	return
}

// GetCommentScoreInfoList 分页获取评分记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentScoreService *CommentScoreService) GetCommentScoreInfoList(info adminReq.CommentScoreSearch) (list []admin.CommentScore, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.CommentScore{})
	var commentScores []admin.CommentScore
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&commentScores).Error
	return commentScores, total, err
}
