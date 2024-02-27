package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type InformationService struct {
}

// CreateInformation 创建信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (informationService *InformationService) CreateInformation(information *admin.Information) (err error) {
	err = global.GVA_DB.Create(information).Error
	return err
}

// DeleteInformation 删除信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (informationService *InformationService) DeleteInformation(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.Information{}, "id = ?", ID).Error
	return err
}

// DeleteInformationByIds 批量删除信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (informationService *InformationService) DeleteInformationByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.Information{}, "id in ?", IDs).Error
	return err
}

// UpdateInformation 更新信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (informationService *InformationService) UpdateInformation(information admin.Information) (err error) {
	err = global.GVA_DB.Save(&information).Error
	return err
}

// GetInformation 根据ID获取信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (informationService *InformationService) GetInformation(ID string) (information admin.Information, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&information).Error
	return
}

// GetInformationInfoList 分页获取信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (informationService *InformationService) GetInformationInfoList(info adminReq.InformationSearch) (list []admin.Information, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.Information{})
	var informations []admin.Information
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Category != 0 {
		db = db.Where("category = ?", info.Category)
	}
	if info.Title != "" {
		db = db.Where("title = ?", info.Title)
	}
	if info.Source != "" {
		db = db.Where("source = ?", info.Source)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&informations).Error
	return informations, total, err
}
