package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskStageRouter struct {
}

// InitTaskStageRouter 初始化 任务阶段 路由信息
func (s *TaskStageRouter) InitTaskStageRouter(Router *gin.RouterGroup) {
	taskStageRouter := Router.Group("taskStage").Use(middleware.OperationRecord())
	taskStageRouterWithoutRecord := Router.Group("taskStage")
	var taskStageApi = v1.ApiGroupApp.AdminApiGroup.TaskStageApi
	{
		taskStageRouter.POST("createTaskStage", taskStageApi.CreateTaskStage)             // 新建任务阶段
		taskStageRouter.DELETE("deleteTaskStage", taskStageApi.DeleteTaskStage)           // 删除任务阶段
		taskStageRouter.DELETE("deleteTaskStageByIds", taskStageApi.DeleteTaskStageByIds) // 批量删除任务阶段
		taskStageRouter.PUT("updateTaskStage", taskStageApi.UpdateTaskStage)              // 更新任务阶段
	}
	{
		taskStageRouterWithoutRecord.GET("findTaskStage", taskStageApi.FindTaskStage)       // 根据ID获取任务阶段
		taskStageRouterWithoutRecord.GET("getTaskStageList", taskStageApi.GetTaskStageList) // 获取任务阶段列表
	}
}
