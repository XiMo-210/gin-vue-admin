package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	adminResp "github.com/flipped-aurora/gin-vue-admin/server/model/admin/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type UserTaskApi struct {
}

var userTaskService = service.ServiceGroupApp.AdminServiceGroup.UserTaskService

func (userTaskApi *UserTaskApi) GetUserTaskCondition(c *gin.Context) {
	ID := c.Query("ID")

	wxUser, err := wxUserService.GetWxUser(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	var task admin.Task
	var taskStages []admin.TaskStage
	var curUserTask admin.UserTask
	if *wxUser.CurTask != 0 {
		task, err = taskService.GetTask(strconv.Itoa(*wxUser.CurTask))
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
			return
		}

		taskStages, err = taskService.GetTaskStages(strconv.Itoa(*wxUser.CurTask))
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
			return
		}

		curUserTask, err = userTaskService.GetUserTaskByUserAndTask(wxUser.ID, uint(*wxUser.CurTask))
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
			return
		}
	}

	completedTasks, err := userTaskService.GetUserCompletedTasks(ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(adminResp.GetUserTaskCondition{
		CurTask: admin.TaskWithStages{
			Task:       task,
			TaskStages: taskStages,
		},
		CurUserTask:    curUserTask,
		CompletedTasks: completedTasks,
	}, "获取成功", c)

}

// CreateUserTask 创建用户任务记录
// @Tags UserTask
// @Summary 创建用户任务记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.UserTask true "创建用户任务记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userTask/createUserTask [post]
func (userTaskApi *UserTaskApi) CreateUserTask(c *gin.Context) {
	var userTask admin.UserTask
	err := c.ShouldBindJSON(&userTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userTaskService.CreateUserTask(&userTask); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserTask 删除用户任务记录
// @Tags UserTask
// @Summary 删除用户任务记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.UserTask true "删除用户任务记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userTask/deleteUserTask [delete]
func (userTaskApi *UserTaskApi) DeleteUserTask(c *gin.Context) {
	ID := c.Query("ID")
	if err := userTaskService.DeleteUserTask(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserTaskByIds 批量删除用户任务记录
// @Tags UserTask
// @Summary 批量删除用户任务记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userTask/deleteUserTaskByIds [delete]
func (userTaskApi *UserTaskApi) DeleteUserTaskByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := userTaskService.DeleteUserTaskByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserTask 更新用户任务记录
// @Tags UserTask
// @Summary 更新用户任务记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.UserTask true "更新用户任务记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userTask/updateUserTask [put]
func (userTaskApi *UserTaskApi) UpdateUserTask(c *gin.Context) {
	var userTask admin.UserTask
	err := c.ShouldBindJSON(&userTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userTaskService.UpdateUserTask(userTask); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserTask 用id查询用户任务记录
// @Tags UserTask
// @Summary 用id查询用户任务记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.UserTask true "用id查询用户任务记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userTask/findUserTask [get]
func (userTaskApi *UserTaskApi) FindUserTask(c *gin.Context) {
	ID := c.Query("ID")
	if reuserTask, err := userTaskService.GetUserTask(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserTask": reuserTask}, c)
	}
}

// GetUserTaskList 分页获取用户任务记录列表
// @Tags UserTask
// @Summary 分页获取用户任务记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.UserTaskSearch true "分页获取用户任务记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userTask/getUserTaskList [get]
func (userTaskApi *UserTaskApi) GetUserTaskList(c *gin.Context) {
	var pageInfo adminReq.UserTaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userTaskService.GetUserTaskInfoList(pageInfo); err != nil {
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
