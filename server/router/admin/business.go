package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BusinessRouter struct {
}

// InitBusinessRouter 初始化 店铺商家 路由信息
func (s *BusinessRouter) InitBusinessRouter(Router *gin.RouterGroup) {
	businessRouter := Router.Group("business").Use(middleware.OperationRecord())
	businessRouterWithoutRecord := Router.Group("business")
	var businessApi = v1.ApiGroupApp.AdminApiGroup.BusinessApi
	{
		businessRouter.POST("createBusiness", businessApi.CreateBusiness)             // 新建店铺商家
		businessRouter.DELETE("deleteBusiness", businessApi.DeleteBusiness)           // 删除店铺商家
		businessRouter.DELETE("deleteBusinessByIds", businessApi.DeleteBusinessByIds) // 批量删除店铺商家
		businessRouter.PUT("updateBusiness", businessApi.UpdateBusiness)              // 更新店铺商家
	}
	{
		businessRouterWithoutRecord.GET("findBusiness", businessApi.FindBusiness)       // 根据ID获取店铺商家
		businessRouterWithoutRecord.GET("getBusinessList", businessApi.GetBusinessList) // 获取店铺商家列表
	}
	{
		businessRouter.POST("createBusinessByAdmin", businessApi.CreateBusinessByAdmin)         // 新建店铺商家ByAdmin
		businessRouter.PUT("updateBusinessByAdmin", businessApi.UpdateBusinessByAdmin)          // 更新店铺商家ByAdmin
		businessRouterWithoutRecord.GET("findBusinessByAdmin", businessApi.FindBusinessByAdmin) // 根据ID获取店铺商家ByAdmin
	}
}
