package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type BusinessService struct {
}

// CreateBusiness 创建店铺商家记录
// Author [piexlmax](https://github.com/piexlmax)
func (businessService *BusinessService) CreateBusiness(business *admin.Business) (err error) {
	if err = global.GVA_DB.Create(business).Error; err != nil {
		return err
	}
	err = global.GVA_DB.Create(&admin.CommentScore{
		Category: admin.BUSINESS_COMMENT,
		TargetId: business.ID,
	}).Error
	return err
}

// DeleteBusiness 删除店铺商家记录
// Author [piexlmax](https://github.com/piexlmax)
func (businessService *BusinessService) DeleteBusiness(ID string) (err error) {
	if err = global.GVA_DB.Delete(&admin.Business{}, "id = ?", ID).Error; err != nil {
		return err
	}
	err = global.GVA_DB.Where(&admin.CommentScore{
		Category: admin.BUSINESS_COMMENT,
	}).Where("target_id = ?", ID).
		Delete(&admin.CommentScore{}).Error
	return err
}

// DeleteBusinessByIds 批量删除店铺商家记录
// Author [piexlmax](https://github.com/piexlmax)
func (businessService *BusinessService) DeleteBusinessByIds(IDs []string) (err error) {
	if err = global.GVA_DB.Delete(&[]admin.Business{}, "id in ?", IDs).Error; err != nil {
		return err
	}
	err = global.GVA_DB.Where(&admin.CommentScore{
		Category: admin.BUSINESS_COMMENT,
	}).Where("target_id IN ?", IDs).
		Delete(&admin.CommentScore{}).Error
	return err
}

// UpdateBusiness 更新店铺商家记录
// Author [piexlmax](https://github.com/piexlmax)
func (businessService *BusinessService) UpdateBusiness(business admin.Business) (err error) {
	err = global.GVA_DB.Save(&business).Error
	return err
}

// GetBusiness 根据ID获取店铺商家记录
// Author [piexlmax](https://github.com/piexlmax)
func (businessService *BusinessService) GetBusiness(sysUserId uint) (business admin.Business, err error) {
	err = global.GVA_DB.Where("sys_user_id = ?", sysUserId).First(&business).Error
	return
}

// GetBusinessInfoList 分页获取店铺商家记录
// Author [piexlmax](https://github.com/piexlmax)
func (businessService *BusinessService) GetBusinessInfoList(info adminReq.BusinessSearch) (list []admin.Business, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.Business{})
	var businesss []admin.Business
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&businesss).Error
	return businesss, total, err
}

func (businessService *BusinessService) GetBusinessByAdmin(ID string) (business admin.Business, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&business).Error
	return
}
