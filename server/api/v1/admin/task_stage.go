package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TaskStageApi struct {
}

var taskStageService = service.ServiceGroupApp.AdminServiceGroup.TaskStageService

// CreateTaskStage 创建任务阶段
// @Tags TaskStage
// @Summary 创建任务阶段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.TaskStage true "创建任务阶段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /taskStage/createTaskStage [post]
func (taskStageApi *TaskStageApi) CreateTaskStage(c *gin.Context) {
	var taskStage admin.TaskStage
	err := c.ShouldBindJSON(&taskStage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := taskStageService.CreateTaskStage(&taskStage); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTaskStage 删除任务阶段
// @Tags TaskStage
// @Summary 删除任务阶段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.TaskStage true "删除任务阶段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskStage/deleteTaskStage [delete]
func (taskStageApi *TaskStageApi) DeleteTaskStage(c *gin.Context) {
	ID := c.Query("ID")
	if err := taskStageService.DeleteTaskStage(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTaskStageByIds 批量删除任务阶段
// @Tags TaskStage
// @Summary 批量删除任务阶段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /taskStage/deleteTaskStageByIds [delete]
func (taskStageApi *TaskStageApi) DeleteTaskStageByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := taskStageService.DeleteTaskStageByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTaskStage 更新任务阶段
// @Tags TaskStage
// @Summary 更新任务阶段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.TaskStage true "更新任务阶段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /taskStage/updateTaskStage [put]
func (taskStageApi *TaskStageApi) UpdateTaskStage(c *gin.Context) {
	var taskStage admin.TaskStage
	err := c.ShouldBindJSON(&taskStage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := taskStageService.UpdateTaskStage(taskStage); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTaskStage 用id查询任务阶段
// @Tags TaskStage
// @Summary 用id查询任务阶段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.TaskStage true "用id查询任务阶段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /taskStage/findTaskStage [get]
func (taskStageApi *TaskStageApi) FindTaskStage(c *gin.Context) {
	ID := c.Query("ID")
	if retaskStage, err := taskStageService.GetTaskStage(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retaskStage": retaskStage}, c)
	}
}

// GetTaskStageList 分页获取任务阶段列表
// @Tags TaskStage
// @Summary 分页获取任务阶段列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.TaskStageSearch true "分页获取任务阶段列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /taskStage/getTaskStageList [get]
func (taskStageApi *TaskStageApi) GetTaskStageList(c *gin.Context) {
	var pageInfo adminReq.TaskStageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := taskStageService.GetTaskStageInfoList(pageInfo); err != nil {
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
