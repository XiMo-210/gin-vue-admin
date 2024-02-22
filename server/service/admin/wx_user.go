package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type WxUserService struct {
}

// CreateWxUser 创建用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (wxUserService *WxUserService) CreateWxUser(wxUser *admin.WxUser) (err error) {
	err = global.GVA_DB.Create(wxUser).Error
	return err
}

// DeleteWxUser 删除用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (wxUserService *WxUserService) DeleteWxUser(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.WxUser{}, "id = ?", ID).Error
	return err
}

// DeleteWxUserByIds 批量删除用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (wxUserService *WxUserService) DeleteWxUserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.WxUser{}, "id in ?", IDs).Error
	return err
}

// UpdateWxUser 更新用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (wxUserService *WxUserService) UpdateWxUser(wxUser admin.WxUser) (err error) {
	err = global.GVA_DB.Save(&wxUser).Error
	return err
}

// GetWxUser 根据ID获取用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (wxUserService *WxUserService) GetWxUser(ID string) (wxUser admin.WxUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&wxUser).Error
	return
}

// GetWxUserInfoList 分页获取用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (wxUserService *WxUserService) GetWxUserInfoList(info adminReq.WxUserSearch) (list []admin.WxUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.WxUser{})
	var wxUsers []admin.WxUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StudentInfoId != nil {
		db = db.Where("student_info_id = ?", info.StudentInfoId)
	}
	if info.Username != "" {
		db = db.Where("username = ?", info.Username)
	}
	if info.StartExp != nil && info.EndExp != nil {
		db = db.Where("exp BETWEEN ? AND ? ", info.StartExp, info.EndExp)
	}
	if info.StartPoints != nil && info.EndPoints != nil {
		db = db.Where("points BETWEEN ? AND ? ", info.StartPoints, info.EndPoints)
	}
	if info.IsCompletedMain != nil {
		db = db.Where("is_completed_main = ?", info.IsCompletedMain)
	}
	if info.CurTask != nil {
		db = db.Where("cur_task = ?", info.CurTask)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["exp"] = true
	orderMap["points"] = true
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

	err = db.Find(&wxUsers).Error
	return wxUsers, total, err
}
