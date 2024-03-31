package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ReviewRecordsRouter struct {
}

// InitReviewRecordsRouter 初始化 审核记录 路由信息
func (s *ReviewRecordsRouter) InitReviewRecordsRouter(Router *gin.RouterGroup) {
	reviewRecordRouter := Router.Group("reviewRecord").Use(middleware.OperationRecord())
	reviewRecordRouterWithoutRecord := Router.Group("reviewRecord")
	var reviewRecordApi = v1.ApiGroupApp.AdminApiGroup.ReviewRecordsApi
	{
		reviewRecordRouter.POST("createReviewRecords", reviewRecordApi.CreateReviewRecords)             // 新建审核记录
		reviewRecordRouter.DELETE("deleteReviewRecords", reviewRecordApi.DeleteReviewRecords)           // 删除审核记录
		reviewRecordRouter.DELETE("deleteReviewRecordsByIds", reviewRecordApi.DeleteReviewRecordsByIds) // 批量删除审核记录
		reviewRecordRouter.PUT("updateReviewRecords", reviewRecordApi.UpdateReviewRecords)              // 更新审核记录
	}
	{
		reviewRecordRouterWithoutRecord.GET("findReviewRecords", reviewRecordApi.FindReviewRecords)       // 根据ID获取审核记录
		reviewRecordRouterWithoutRecord.GET("getReviewRecordsList", reviewRecordApi.GetReviewRecordsList) // 获取审核记录列表
	}
}
