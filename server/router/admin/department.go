package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DepartmentRouter struct {
}

// InitDepartmentRouter 初始化 部门 路由信息
func (s *DepartmentRouter) InitDepartmentRouter(Router *gin.RouterGroup) {
	departmentRouter := Router.Group("department").Use(middleware.OperationRecord())
	departmentRouterWithoutRecord := Router.Group("department")
	var departmentApi = v1.ApiGroupApp.AdminApiGroup.DepartmentApi
	{
		departmentRouter.POST("createDepartment", departmentApi.CreateDepartment)             // 新建部门
		departmentRouter.DELETE("deleteDepartment", departmentApi.DeleteDepartment)           // 删除部门
		departmentRouter.DELETE("deleteDepartmentByIds", departmentApi.DeleteDepartmentByIds) // 批量删除部门
		departmentRouter.PUT("updateDepartment", departmentApi.UpdateDepartment)              // 更新部门
	}
	{
		departmentRouterWithoutRecord.GET("findDepartment", departmentApi.FindDepartment)       // 根据ID获取部门
		departmentRouterWithoutRecord.GET("getDepartmentList", departmentApi.GetDepartmentList) // 获取部门列表
	}
}
