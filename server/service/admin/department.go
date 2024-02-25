package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type DepartmentService struct {
}

// CreateDepartment 创建部门记录
// Author [piexlmax](https://github.com/piexlmax)
func (departmentService *DepartmentService) CreateDepartment(department *admin.Department) (err error) {
	err = global.GVA_DB.Create(department).Error
	return err
}

// DeleteDepartment 删除部门记录
// Author [piexlmax](https://github.com/piexlmax)
func (departmentService *DepartmentService) DeleteDepartment(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.Department{}, "id = ?", ID).Error
	return err
}

// DeleteDepartmentByIds 批量删除部门记录
// Author [piexlmax](https://github.com/piexlmax)
func (departmentService *DepartmentService) DeleteDepartmentByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.Department{}, "id in ?", IDs).Error
	return err
}

// UpdateDepartment 更新部门记录
// Author [piexlmax](https://github.com/piexlmax)
func (departmentService *DepartmentService) UpdateDepartment(department admin.Department) (err error) {
	err = global.GVA_DB.Save(&department).Error
	return err
}

// GetDepartment 根据ID获取部门记录
// Author [piexlmax](https://github.com/piexlmax)
func (departmentService *DepartmentService) GetDepartment(ID string) (department admin.Department, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&department).Error
	return
}

// GetDepartmentInfoList 分页获取部门记录
// Author [piexlmax](https://github.com/piexlmax)
func (departmentService *DepartmentService) GetDepartmentInfoList(info adminReq.DepartmentSearch) (list []admin.Department, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.Department{})
	var departments []admin.Department
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

	err = db.Find(&departments).Error
	return departments, total, err
}
