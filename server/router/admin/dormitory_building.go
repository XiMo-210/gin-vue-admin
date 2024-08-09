package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DormitoryBuildingRouter struct {
}

// InitDormitoryBuildingRouter 初始化 寝室楼 路由信息
func (s *DormitoryBuildingRouter) InitDormitoryBuildingRouter(Router *gin.RouterGroup) {
	dormitoryBuildingRouter := Router.Group("dormitoryBuilding").Use(middleware.OperationRecord())
	dormitoryBuildingRouterWithoutRecord := Router.Group("dormitoryBuilding")
	var dormitoryBuildingApi = v1.ApiGroupApp.AdminApiGroup.DormitoryBuildingApi
	{
		dormitoryBuildingRouter.POST("createDormitoryBuilding", dormitoryBuildingApi.CreateDormitoryBuilding)   // 新建寝室楼
		dormitoryBuildingRouter.DELETE("deleteDormitoryBuilding", dormitoryBuildingApi.DeleteDormitoryBuilding) // 删除寝室楼
		dormitoryBuildingRouter.DELETE("deleteDormitoryBuildingByIds", dormitoryBuildingApi.DeleteDormitoryBuildingByIds) // 批量删除寝室楼
		dormitoryBuildingRouter.PUT("updateDormitoryBuilding", dormitoryBuildingApi.UpdateDormitoryBuilding)    // 更新寝室楼
	}
	{
		dormitoryBuildingRouterWithoutRecord.GET("findDormitoryBuilding", dormitoryBuildingApi.FindDormitoryBuilding)        // 根据ID获取寝室楼
		dormitoryBuildingRouterWithoutRecord.GET("getDormitoryBuildingList", dormitoryBuildingApi.GetDormitoryBuildingList)  // 获取寝室楼列表
	}
}
