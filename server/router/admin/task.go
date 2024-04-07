package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskRouter struct {
}

// InitTaskRouter 初始化 任务 路由信息
func (s *TaskRouter) InitTaskRouter(Router *gin.RouterGroup) {
	taskRouter := Router.Group("task").Use(middleware.OperationRecord())
	taskRouterWithoutRecord := Router.Group("task")
	var taskApi = v1.ApiGroupApp.AdminApiGroup.TaskApi
	{
		taskRouter.POST("createTask", taskApi.CreateTask)             // 新建任务
		taskRouter.DELETE("deleteTask", taskApi.DeleteTask)           // 删除任务
		taskRouter.DELETE("deleteTaskByIds", taskApi.DeleteTaskByIds) // 批量删除任务
		taskRouter.PUT("updateTask", taskApi.UpdateTask)              // 更新任务
	}
	{
		taskRouterWithoutRecord.GET("findTask", taskApi.FindTask)                   // 根据ID获取任务
		taskRouterWithoutRecord.GET("getTaskList", taskApi.GetTaskList)             // 获取任务列表
		taskRouterWithoutRecord.GET("completeCondition", taskApi.CompleteCondition) // 任务完成情况
		taskRouterWithoutRecord.GET("dayCompleteNum", taskApi.DayCompleteNum)       // 任务每日完成数
		taskRouterWithoutRecord.GET("completeRecords", taskApi.CompleteRecords)     // 任务完成记录
		taskRouterWithoutRecord.POST("taskReset", taskApi.TaskReset)                // 任务重置
	}
}
