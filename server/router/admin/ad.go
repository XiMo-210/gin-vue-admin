package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AdRouter struct {
}

// InitAdRouter 初始化 广告 路由信息
func (s *AdRouter) InitAdRouter(Router *gin.RouterGroup) {
	adRouter := Router.Group("ad").Use(middleware.OperationRecord())
	adRouterWithoutRecord := Router.Group("ad")
	var adApi = v1.ApiGroupApp.AdminApiGroup.AdApi
	{
		adRouter.POST("createAd", adApi.CreateAd)             // 新建广告
		adRouter.DELETE("deleteAd", adApi.DeleteAd)           // 删除广告
		adRouter.DELETE("deleteAdByIds", adApi.DeleteAdByIds) // 批量删除广告
		adRouter.PUT("updateAd", adApi.UpdateAd)              // 更新广告
	}
	{
		adRouterWithoutRecord.GET("findAd", adApi.FindAd)       // 根据ID获取广告
		adRouterWithoutRecord.GET("getAdList", adApi.GetAdList) // 获取广告列表
	}
}
