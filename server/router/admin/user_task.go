package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserTaskRouter struct {
}

// InitUserTaskRouter 初始化 用户任务记录 路由信息
func (s *UserTaskRouter) InitUserTaskRouter(Router *gin.RouterGroup) {
	userTaskRouter := Router.Group("userTask").Use(middleware.OperationRecord())
	userTaskRouterWithoutRecord := Router.Group("userTask")
	var userTaskApi = v1.ApiGroupApp.AdminApiGroup.UserTaskApi
	{
		userTaskRouter.POST("createUserTask", userTaskApi.CreateUserTask)             // 新建用户任务记录
		userTaskRouter.DELETE("deleteUserTask", userTaskApi.DeleteUserTask)           // 删除用户任务记录
		userTaskRouter.DELETE("deleteUserTaskByIds", userTaskApi.DeleteUserTaskByIds) // 批量删除用户任务记录
		userTaskRouter.PUT("updateUserTask", userTaskApi.UpdateUserTask)              // 更新用户任务记录
	}
	{
		userTaskRouterWithoutRecord.GET("findUserTask", userTaskApi.FindUserTask)       // 根据ID获取用户任务记录
		userTaskRouterWithoutRecord.GET("getUserTaskList", userTaskApi.GetUserTaskList) // 获取用户任务记录列表
	}
}
