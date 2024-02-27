package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type OrganizationService struct {
}

// CreateOrganization 创建组织社团记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationService *OrganizationService) CreateOrganization(organization *admin.Organization) (err error) {
	if err = global.GVA_DB.Create(organization).Error; err != nil {
		return err
	}
	err = global.GVA_DB.Create(&admin.CommentScore{
		Category: admin.ORGANIZATION_COMMENT,
		TargetId: organization.ID,
	}).Error
	return err
}

// DeleteOrganization 删除组织社团记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationService *OrganizationService) DeleteOrganization(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.Organization{}, "id = ?", ID).Error
	return err
}

// DeleteOrganizationByIds 批量删除组织社团记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationService *OrganizationService) DeleteOrganizationByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.Organization{}, "id in ?", IDs).Error
	return err
}

// UpdateOrganization 更新组织社团记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationService *OrganizationService) UpdateOrganization(organization admin.Organization) (err error) {
	err = global.GVA_DB.Save(&organization).Error
	return err
}

// GetOrganization 根据ID获取组织社团记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationService *OrganizationService) GetOrganization(sysUserId uint) (organization admin.Organization, err error) {
	err = global.GVA_DB.Where("sys_user_id = ?", sysUserId).First(&organization).Error
	return
}

// GetOrganizationInfoList 分页获取组织社团记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationService *OrganizationService) GetOrganizationInfoList(info adminReq.OrganizationSearch) (list []admin.Organization, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.Organization{})
	var organizations []admin.Organization
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

	err = db.Find(&organizations).Error
	return organizations, total, err
}
