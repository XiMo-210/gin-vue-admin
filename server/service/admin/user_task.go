package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type UserTaskService struct {
}

// CreateUserTask 创建用户任务记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (userTaskService *UserTaskService) CreateUserTask(userTask *admin.UserTask) (err error) {
	err = global.GVA_DB.Create(userTask).Error
	return err
}

// DeleteUserTask 删除用户任务记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (userTaskService *UserTaskService) DeleteUserTask(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.UserTask{}, "id = ?", ID).Error
	return err
}

// DeleteUserTaskByIds 批量删除用户任务记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (userTaskService *UserTaskService) DeleteUserTaskByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.UserTask{}, "id in ?", IDs).Error
	return err
}

// UpdateUserTask 更新用户任务记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (userTaskService *UserTaskService) UpdateUserTask(userTask admin.UserTask) (err error) {
	err = global.GVA_DB.Save(&userTask).Error
	return err
}

// GetUserTask 根据ID获取用户任务记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (userTaskService *UserTaskService) GetUserTask(ID string) (userTask admin.UserTask, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&userTask).Error
	return
}

// GetUserTaskInfoList 分页获取用户任务记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (userTaskService *UserTaskService) GetUserTaskInfoList(info adminReq.UserTaskSearch) (list []admin.UserTask, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.UserTask{})
	var userTasks []admin.UserTask
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

	err = db.Find(&userTasks).Error
	return userTasks, total, err
}
