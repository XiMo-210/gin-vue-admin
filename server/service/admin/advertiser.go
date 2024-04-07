package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type AdvertiserService struct {
}

// CreateAdvertiser 创建广告主记录
// Author [piexlmax](https://github.com/piexlmax)
func (advertiserService *AdvertiserService) CreateAdvertiser(advertiser *admin.Advertiser) (err error) {
	err = global.GVA_DB.Create(advertiser).Error
	return err
}

// DeleteAdvertiser 删除广告主记录
// Author [piexlmax](https://github.com/piexlmax)
func (advertiserService *AdvertiserService) DeleteAdvertiser(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.Advertiser{}, "id = ?", ID).Error
	return err
}

// DeleteAdvertiserByIds 批量删除广告主记录
// Author [piexlmax](https://github.com/piexlmax)
func (advertiserService *AdvertiserService) DeleteAdvertiserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.Advertiser{}, "id in ?", IDs).Error
	return err
}

// UpdateAdvertiser 更新广告主记录
// Author [piexlmax](https://github.com/piexlmax)
func (advertiserService *AdvertiserService) UpdateAdvertiser(advertiser admin.Advertiser) (err error) {
	err = global.GVA_DB.Save(&advertiser).Error
	return err
}

// GetAdvertiser 根据ID获取广告主记录
// Author [piexlmax](https://github.com/piexlmax)
func (advertiserService *AdvertiserService) GetAdvertiser(sysUserId uint) (advertiser admin.Advertiser, err error) {
	err = global.GVA_DB.Where("sys_user_id = ?", sysUserId).First(&advertiser).Error
	return
}

// GetAdvertiserInfoList 分页获取广告主记录
// Author [piexlmax](https://github.com/piexlmax)
func (advertiserService *AdvertiserService) GetAdvertiserInfoList(info adminReq.AdvertiserSearch) (list []admin.Advertiser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.Advertiser{})
	var advertisers []admin.Advertiser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.SubjectInfo != "" {
		db = db.Where("subject_info = ?", info.SubjectInfo)
	}
	if info.License != "" {
		db = db.Where("license = ?", info.License)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&advertisers).Error
	return advertisers, total, err
}

func (advertiserService *AdvertiserService) GetAdvertiserByAdmin(ID string) (advertiser admin.Advertiser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&advertiser).Error
	return
}
