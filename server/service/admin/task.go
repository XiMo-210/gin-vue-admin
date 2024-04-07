package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin/response"
)

type TaskService struct {
}

// CreateTask 创建任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) CreateTask(task *admin.Task) (err error) {
	if err = global.GVA_DB.Create(task).Error; err != nil {
		return err
	}
	err = global.GVA_DB.Create(&admin.CommentScore{
		Category: admin.TASK_COMMENT,
		TargetId: task.ID,
	}).Error
	return err
}

func (taskService *TaskService) CreateTaskStages(taskStages []admin.TaskStage) (err error) {
	err = global.GVA_DB.Omit("id").Create(taskStages).Error
	return err
}

// DeleteTask 删除任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) DeleteTask(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.Task{}, "id = ?", ID).Error
	return err
}

func (taskService *TaskService) DeleteTaskStages(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.TaskStage{}, "task_id = ?", ID).Error
	return err
}

// DeleteTaskByIds 批量删除任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) DeleteTaskByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.Task{}, "id in ?", IDs).Error
	return err
}

func (taskService *TaskService) DeleteTaskStagesByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.TaskStage{}, "task_id in ?", IDs).Error
	return err
}

// UpdateTask 更新任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) UpdateTask(task admin.Task) (err error) {
	err = global.GVA_DB.Save(&task).Error
	return err
}

// GetTask 根据ID获取任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) GetTask(ID string) (task admin.Task, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&task).Error
	return
}

func (taskService *TaskService) GetTaskStages(ID string) (taskStages []admin.TaskStage, err error) {
	err = global.GVA_DB.Where("task_id = ?", ID).Order("stage").Find(&taskStages).Error
	return
}

// GetTaskInfoList 分页获取任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) GetTaskInfoList(info adminReq.TaskSearch) (list []admin.Task, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.Task{})
	var tasks []admin.Task
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Category != nil {
		db = db.Where("category = ?", info.Category)
	}
	if info.Title != "" {
		db = db.Where("title = ?", info.Title)
	}
	if info.Campus != "" {
		db = db.Where("campus = ?", info.Campus)
	}
	if info.College != "" {
		db = db.Where("college = ?", info.College)
	}
	if info.NeedMain != nil {
		db = db.Where("need_main = ?", info.NeedMain)
	}
	if info.StartTime != nil {
		db = db.Where("start_time = ? ", info.StartTime)
	}
	if info.EndTime != nil {
		db = db.Where("end_time = ? ", info.EndTime)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tasks).Error
	return tasks, total, err
}

func (taskService *TaskService) GetCompleteRecords(info adminReq.CompleteRecords) (list []response.CompleteRecords, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&admin.UserTask{}).
		Select("user_tasks.updated_at AS completed_at, user_tasks.id AS user_task_id ,wx_users.avatar, wx_users.username, wx_users.id AS user_id").
		Joins("JOIN wx_users ON user_tasks.user_id = wx_users.id").
		Where("task_id = ? AND cur_stage = 0", info.TaskId)

	if info.UserId != nil {
		db.Where("user_id = ?", info.UserId)
	}
	if info.Username != "" {
		db.Where("username = ?", info.Username)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("completed_at desc").Find(&list).Error
	return list, total, err
}
