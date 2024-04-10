package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PlatformParamRouter struct {
}

// InitPlatformParamRouter 初始化 平台参数 路由信息
func (s *PlatformParamRouter) InitPlatformParamRouter(Router *gin.RouterGroup) {
	platformParamRouter := Router.Group("platformParam").Use(middleware.OperationRecord())
	platformParamRouterWithoutRecord := Router.Group("platformParam")
	var platformParamApi = v1.ApiGroupApp.AdminApiGroup.PlatformParamApi
	{
		platformParamRouter.POST("createPlatformParam", platformParamApi.CreatePlatformParam)             // 新建平台参数
		platformParamRouter.DELETE("deletePlatformParam", platformParamApi.DeletePlatformParam)           // 删除平台参数
		platformParamRouter.DELETE("deletePlatformParamByIds", platformParamApi.DeletePlatformParamByIds) // 批量删除平台参数
		platformParamRouter.PUT("updatePlatformParam", platformParamApi.UpdatePlatformParam)              // 更新平台参数
	}
	{
		platformParamRouterWithoutRecord.GET("findPlatformParam", platformParamApi.FindPlatformParam)       // 根据ID获取平台参数
		platformParamRouterWithoutRecord.GET("getPlatformParamList", platformParamApi.GetPlatformParamList) // 获取平台参数列表
	}
}
