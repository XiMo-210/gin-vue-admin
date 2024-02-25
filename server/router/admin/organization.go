package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrganizationRouter struct {
}

// InitOrganizationRouter 初始化 组织社团 路由信息
func (s *OrganizationRouter) InitOrganizationRouter(Router *gin.RouterGroup) {
	organizationRouter := Router.Group("organization").Use(middleware.OperationRecord())
	organizationRouterWithoutRecord := Router.Group("organization")
	var organizationApi = v1.ApiGroupApp.AdminApiGroup.OrganizationApi
	{
		organizationRouter.POST("createOrganization", organizationApi.CreateOrganization)             // 新建组织社团
		organizationRouter.DELETE("deleteOrganization", organizationApi.DeleteOrganization)           // 删除组织社团
		organizationRouter.DELETE("deleteOrganizationByIds", organizationApi.DeleteOrganizationByIds) // 批量删除组织社团
		organizationRouter.PUT("updateOrganization", organizationApi.UpdateOrganization)              // 更新组织社团
	}
	{
		organizationRouterWithoutRecord.GET("findOrganization", organizationApi.FindOrganization)       // 根据ID获取组织社团
		organizationRouterWithoutRecord.GET("getOrganizationList", organizationApi.GetOrganizationList) // 获取组织社团列表
	}
}
