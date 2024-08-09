package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DormRoomRouter struct {
}

// InitDormRoomRouter 初始化 宿舍 路由信息
func (s *DormRoomRouter) InitDormRoomRouter(Router *gin.RouterGroup) {
	dormRoomRouter := Router.Group("dormRoom").Use(middleware.OperationRecord())
	dormRoomRouterWithoutRecord := Router.Group("dormRoom")
	var dormRoomApi = v1.ApiGroupApp.AdminApiGroup.DormRoomApi
	{
		dormRoomRouter.POST("createDormRoom", dormRoomApi.CreateDormRoom)             // 新建宿舍
		dormRoomRouter.DELETE("deleteDormRoom", dormRoomApi.DeleteDormRoom)           // 删除宿舍
		dormRoomRouter.DELETE("deleteDormRoomByIds", dormRoomApi.DeleteDormRoomByIds) // 批量删除宿舍
		dormRoomRouter.PUT("updateDormRoom", dormRoomApi.UpdateDormRoom)              // 更新宿舍
		dormRoomRouter.GET("alloc", dormRoomApi.DormAlloc)
		dormRoomRouter.GET("student", dormRoomApi.GetDormStudents)
		dormRoomRouter.GET("remind", dormRoomApi.Remind)
	}
	{
		dormRoomRouterWithoutRecord.GET("findDormRoom", dormRoomApi.FindDormRoom)       // 根据ID获取宿舍
		dormRoomRouterWithoutRecord.GET("getDormRoomList", dormRoomApi.GetDormRoomList) // 获取宿舍列表
	}
}
