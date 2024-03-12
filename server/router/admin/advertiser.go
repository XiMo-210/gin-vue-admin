package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AdvertiserRouter struct {
}

// InitAdvertiserRouter 初始化 广告主 路由信息
func (s *AdvertiserRouter) InitAdvertiserRouter(Router *gin.RouterGroup) {
	advertiserRouter := Router.Group("advertiser").Use(middleware.OperationRecord())
	advertiserRouterWithoutRecord := Router.Group("advertiser")
	var advertiserApi = v1.ApiGroupApp.AdminApiGroup.AdvertiserApi
	{
		advertiserRouter.POST("createAdvertiser", advertiserApi.CreateAdvertiser)             // 新建广告主
		advertiserRouter.DELETE("deleteAdvertiser", advertiserApi.DeleteAdvertiser)           // 删除广告主
		advertiserRouter.DELETE("deleteAdvertiserByIds", advertiserApi.DeleteAdvertiserByIds) // 批量删除广告主
		advertiserRouter.PUT("updateAdvertiser", advertiserApi.UpdateAdvertiser)              // 更新广告主
	}
	{
		advertiserRouterWithoutRecord.GET("findAdvertiser", advertiserApi.FindAdvertiser)       // 根据ID获取广告主
		advertiserRouterWithoutRecord.GET("getAdvertiserList", advertiserApi.GetAdvertiserList) // 获取广告主列表
	}
}
