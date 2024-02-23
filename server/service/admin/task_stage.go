package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type TaskStageService struct {
}

// CreateTaskStage 创建任务阶段记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskStageService *TaskStageService) CreateTaskStage(taskStage *admin.TaskStage) (err error) {
	err = global.GVA_DB.Create(taskStage).Error
	return err
}

// DeleteTaskStage 删除任务阶段记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskStageService *TaskStageService) DeleteTaskStage(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.TaskStage{}, "id = ?", ID).Error
	return err
}

// DeleteTaskStageByIds 批量删除任务阶段记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskStageService *TaskStageService) DeleteTaskStageByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.TaskStage{}, "id in ?", IDs).Error
	return err
}

// UpdateTaskStage 更新任务阶段记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskStageService *TaskStageService) UpdateTaskStage(taskStage admin.TaskStage) (err error) {
	err = global.GVA_DB.Save(&taskStage).Error
	return err
}

// GetTaskStage 根据ID获取任务阶段记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskStageService *TaskStageService) GetTaskStage(ID string) (taskStage admin.TaskStage, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&taskStage).Error
	return
}

// GetTaskStageInfoList 分页获取任务阶段记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskStageService *TaskStageService) GetTaskStageInfoList(info adminReq.TaskStageSearch) (list []admin.TaskStage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.TaskStage{})
	var taskStages []admin.TaskStage
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

	err = db.Find(&taskStages).Error
	return taskStages, total, err
}
