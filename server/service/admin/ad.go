package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"time"
)

type AdService struct {
}

// CreateAd 创建广告记录
// Author [piexlmax](https://github.com/piexlmax)
func (adService *AdService) CreateAd(ad *admin.Ad) (err error) {
	err = global.GVA_DB.Create(ad).Error
	return err
}

// DeleteAd 删除广告记录
// Author [piexlmax](https://github.com/piexlmax)
func (adService *AdService) DeleteAd(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.Ad{}, "id = ?", ID).Error
	return err
}

// DeleteAdByIds 批量删除广告记录
// Author [piexlmax](https://github.com/piexlmax)
func (adService *AdService) DeleteAdByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.Ad{}, "id in ?", IDs).Error
	return err
}

// UpdateAd 更新广告记录
// Author [piexlmax](https://github.com/piexlmax)
func (adService *AdService) UpdateAd(ad admin.Ad) (err error) {
	err = global.GVA_DB.Save(&ad).Error
	return err
}

// GetAd 根据ID获取广告记录
// Author [piexlmax](https://github.com/piexlmax)
func (adService *AdService) GetAd(ID string) (ad admin.Ad, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ad).Error
	return
}

// GetAdInfoList 分页获取广告记录
// Author [piexlmax](https://github.com/piexlmax)
func (adService *AdService) GetAdInfoList(info adminReq.AdSearch, advertiserId uint) (list []admin.Ad, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.Ad{}).Where(&admin.Ad{AdvertiserId: advertiserId})
	var ads []admin.Ad
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.AdCategory != nil {
		db = db.Where("ad_category = ?", info.AdCategory)
	}
	if info.CostCategory != nil {
		db = db.Where("cost_category = ?", info.CostCategory)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["cost_impressions"] = true
	orderMap["cost_clicks"] = true
	orderMap["cost"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&ads).Error
	return ads, total, err
}

func (adService *AdService) UpdateAdStatus() (err error) {
	if err = global.GVA_DB.Model(&admin.Ad{}).Where("end_date < ? AND status <> 4", time.Now()).Update("status", 3).Error; err != nil {
		return err
	}
	if err = global.GVA_DB.Model(&admin.Ad{}).Where("cost_category = 1 AND buy_amount = cost_impressions").Update("status", 4).Error; err != nil {
		return err
	}
	if err = global.GVA_DB.Model(&admin.Ad{}).Where("cost_category = 2 AND buy_amount = cost_clicks").Update("status", 4).Error; err != nil {
		return err
	}
	return nil
}
