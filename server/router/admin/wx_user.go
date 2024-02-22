package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WxUserRouter struct {
}

// InitWxUserRouter 初始化 用户 路由信息
func (s *WxUserRouter) InitWxUserRouter(Router *gin.RouterGroup) {
	wxUserRouter := Router.Group("wxUser").Use(middleware.OperationRecord())
	wxUserRouterWithoutRecord := Router.Group("wxUser")
	var wxUserApi = v1.ApiGroupApp.AdminApiGroup.WxUserApi
	{
		wxUserRouter.POST("createWxUser", wxUserApi.CreateWxUser)             // 新建用户
		wxUserRouter.DELETE("deleteWxUser", wxUserApi.DeleteWxUser)           // 删除用户
		wxUserRouter.DELETE("deleteWxUserByIds", wxUserApi.DeleteWxUserByIds) // 批量删除用户
		wxUserRouter.PUT("updateWxUser", wxUserApi.UpdateWxUser)              // 更新用户
	}
	{
		wxUserRouterWithoutRecord.GET("findWxUser", wxUserApi.FindWxUser)       // 根据ID获取用户
		wxUserRouterWithoutRecord.GET("getWxUserList", wxUserApi.GetWxUserList) // 获取用户列表
	}
}
