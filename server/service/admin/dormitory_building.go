package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
    adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type DormitoryBuildingService struct {
}

// CreateDormitoryBuilding 创建寝室楼记录
// Author [piexlmax](https://github.com/piexlmax)
func (dormitoryBuildingService *DormitoryBuildingService) CreateDormitoryBuilding(dormitoryBuilding *admin.DormitoryBuilding) (err error) {
	err = global.GVA_DB.Create(dormitoryBuilding).Error
	return err
}

// DeleteDormitoryBuilding 删除寝室楼记录
// Author [piexlmax](https://github.com/piexlmax)
func (dormitoryBuildingService *DormitoryBuildingService)DeleteDormitoryBuilding(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.DormitoryBuilding{},"id = ?",ID).Error
	return err
}

// DeleteDormitoryBuildingByIds 批量删除寝室楼记录
// Author [piexlmax](https://github.com/piexlmax)
func (dormitoryBuildingService *DormitoryBuildingService)DeleteDormitoryBuildingByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.DormitoryBuilding{},"id in ?",IDs).Error
	return err
}

// UpdateDormitoryBuilding 更新寝室楼记录
// Author [piexlmax](https://github.com/piexlmax)
func (dormitoryBuildingService *DormitoryBuildingService)UpdateDormitoryBuilding(dormitoryBuilding admin.DormitoryBuilding) (err error) {
	err = global.GVA_DB.Save(&dormitoryBuilding).Error
	return err
}

// GetDormitoryBuilding 根据ID获取寝室楼记录
// Author [piexlmax](https://github.com/piexlmax)
func (dormitoryBuildingService *DormitoryBuildingService)GetDormitoryBuilding(ID string) (dormitoryBuilding admin.DormitoryBuilding, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dormitoryBuilding).Error
	return
}

// GetDormitoryBuildingInfoList 分页获取寝室楼记录
// Author [piexlmax](https://github.com/piexlmax)
func (dormitoryBuildingService *DormitoryBuildingService)GetDormitoryBuildingInfoList(info adminReq.DormitoryBuildingSearch) (list []admin.DormitoryBuilding, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&admin.DormitoryBuilding{})
    var dormitoryBuildings []admin.DormitoryBuilding
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Campus != "" {
        db = db.Where("campus = ?",info.Campus)
    }
    if info.Name != "" {
        db = db.Where("name = ?",info.Name)
    }
    if info.Gender != "" {
        db = db.Where("gender = ?",info.Gender)
    }
    if info.ManagerName != "" {
        db = db.Where("manager_name = ?",info.ManagerName)
    }
    if info.ManagerContact != "" {
        db = db.Where("manager_contact = ?",info.ManagerContact)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&dormitoryBuildings).Error
	return  dormitoryBuildings, total, err
}
