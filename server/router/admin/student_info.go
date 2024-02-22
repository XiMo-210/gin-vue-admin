package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StudentInfoRouter struct {
}

// InitStudentInfoRouter 初始化 学生信息 路由信息
func (s *StudentInfoRouter) InitStudentInfoRouter(Router *gin.RouterGroup) {
	studentInfoRouter := Router.Group("studentInfo").Use(middleware.OperationRecord())
	studentInfoRouterWithoutRecord := Router.Group("studentInfo")
	var studentInfoApi = v1.ApiGroupApp.AdminApiGroup.StudentInfoApi
	{
		studentInfoRouter.POST("createStudentInfo", studentInfoApi.CreateStudentInfo)             // 新建学生信息
		studentInfoRouter.DELETE("deleteStudentInfo", studentInfoApi.DeleteStudentInfo)           // 删除学生信息
		studentInfoRouter.DELETE("deleteStudentInfoByIds", studentInfoApi.DeleteStudentInfoByIds) // 批量删除学生信息
		studentInfoRouter.PUT("updateStudentInfo", studentInfoApi.UpdateStudentInfo)              // 更新学生信息
	}
	{
		studentInfoRouterWithoutRecord.GET("findStudentInfo", studentInfoApi.FindStudentInfo)       // 根据ID获取学生信息
		studentInfoRouterWithoutRecord.GET("getStudentInfoList", studentInfoApi.GetStudentInfoList) // 获取学生信息列表
	}
}
