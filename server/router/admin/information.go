package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InformationRouter struct {
}

// InitInformationRouter 初始化 信息 路由信息
func (s *InformationRouter) InitInformationRouter(Router *gin.RouterGroup) {
	informationRouter := Router.Group("information").Use(middleware.OperationRecord())
	informationRouterWithoutRecord := Router.Group("information")
	var informationApi = v1.ApiGroupApp.AdminApiGroup.InformationApi
	{
		informationRouter.POST("createInformation", informationApi.CreateInformation)             // 新建信息
		informationRouter.DELETE("deleteInformation", informationApi.DeleteInformation)           // 删除信息
		informationRouter.DELETE("deleteInformationByIds", informationApi.DeleteInformationByIds) // 批量删除信息
		informationRouter.PUT("updateInformation", informationApi.UpdateInformation)              // 更新信息
	}
	{
		informationRouterWithoutRecord.GET("findInformation", informationApi.FindInformation)       // 根据ID获取信息
		informationRouterWithoutRecord.GET("getInformationList", informationApi.GetInformationList) // 获取信息列表
	}
}
