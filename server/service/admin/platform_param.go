package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type PlatformParamService struct {
}

// CreatePlatformParam 创建平台参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformParamService *PlatformParamService) CreatePlatformParam(platformParam *admin.PlatformParam) (err error) {
	err = global.GVA_DB.Create(platformParam).Error
	return err
}

// DeletePlatformParam 删除平台参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformParamService *PlatformParamService) DeletePlatformParam(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.PlatformParam{}, "id = ?", ID).Error
	return err
}

// DeletePlatformParamByIds 批量删除平台参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformParamService *PlatformParamService) DeletePlatformParamByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.PlatformParam{}, "id in ?", IDs).Error
	return err
}

// UpdatePlatformParam 更新平台参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformParamService *PlatformParamService) UpdatePlatformParam(platformParam admin.PlatformParam) (err error) {
	err = global.GVA_DB.Save(&platformParam).Error
	return err
}

// GetPlatformParam 根据ID获取平台参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformParamService *PlatformParamService) GetPlatformParam(ID string) (platformParam admin.PlatformParam, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&platformParam).Error
	return
}

// GetPlatformParamInfoList 分页获取平台参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (platformParamService *PlatformParamService) GetPlatformParamInfoList(info adminReq.PlatformParamSearch) (list []admin.PlatformParam, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.PlatformParam{})
	var platformParams []admin.PlatformParam
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

	err = db.Find(&platformParams).Error
	return platformParams, total, err
}
