package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type TaskApi struct {
}

var taskService = service.ServiceGroupApp.AdminServiceGroup.TaskService

// CreateTask 创建任务
// @Tags Task
// @Summary 创建任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Task true "创建任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /task/createTask [post]
func (taskApi *TaskApi) CreateTask(c *gin.Context) {
	var taskWithStages admin.TaskWithStages
	err := c.ShouldBindJSON(&taskWithStages)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := taskService.CreateTask(&taskWithStages.Task); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	for i, _ := range taskWithStages.TaskStages {
		taskWithStages.TaskStages[i].TaskId = taskWithStages.Task.ID
	}

	if err := taskService.CreateTaskStages(taskWithStages.TaskStages); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// DeleteTask 删除任务
// @Tags Task
// @Summary 删除任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Task true "删除任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task/deleteTask [delete]
func (taskApi *TaskApi) DeleteTask(c *gin.Context) {
	ID := c.Query("ID")
	if err := taskService.DeleteTask(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}

	if err := taskService.DeleteTaskStages(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// DeleteTaskByIds 批量删除任务
// @Tags Task
// @Summary 批量删除任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /task/deleteTaskByIds [delete]
func (taskApi *TaskApi) DeleteTaskByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := taskService.DeleteTaskByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}

	if err := taskService.DeleteTaskStagesByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}

	response.OkWithMessage("批量删除成功", c)
}

// UpdateTask 更新任务
// @Tags Task
// @Summary 更新任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Task true "更新任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /task/updateTask [put]
func (taskApi *TaskApi) UpdateTask(c *gin.Context) {
	var taskWithStages admin.TaskWithStages
	err := c.ShouldBindJSON(&taskWithStages)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := taskService.UpdateTask(taskWithStages.Task); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	if err := taskService.DeleteTaskStages(strconv.Itoa(int(taskWithStages.Task.ID))); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	for i, _ := range taskWithStages.TaskStages {
		taskWithStages.TaskStages[i].TaskId = taskWithStages.Task.ID

	}

	if err := taskService.CreateTaskStages(taskWithStages.TaskStages); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// FindTask 用id查询任务
// @Tags Task
// @Summary 用id查询任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.Task true "用id查询任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /task/findTask [get]
func (taskApi *TaskApi) FindTask(c *gin.Context) {
	ID := c.Query("ID")
	task, err := taskService.GetTask(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	taskStages, err := taskService.GetTaskStages(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	response.OkWithData(gin.H{"retask": admin.TaskWithStages{
		Task:       task,
		TaskStages: taskStages,
	}}, c)
}

// GetTaskList 分页获取任务列表
// @Tags Task
// @Summary 分页获取任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.TaskSearch true "分页获取任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/getTaskList [get]
func (taskApi *TaskApi) GetTaskList(c *gin.Context) {
	var pageInfo adminReq.TaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := taskService.GetTaskInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
