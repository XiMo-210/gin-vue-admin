package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CommentScoreRouter struct {
}

// InitCommentScoreRouter 初始化 评分 路由信息
func (s *CommentScoreRouter) InitCommentScoreRouter(Router *gin.RouterGroup) {
	commentScoreRouter := Router.Group("commentScore").Use(middleware.OperationRecord())
	commentScoreRouterWithoutRecord := Router.Group("commentScore")
	var commentScoreApi = v1.ApiGroupApp.AdminApiGroup.CommentScoreApi
	{
		commentScoreRouter.POST("createCommentScore", commentScoreApi.CreateCommentScore)             // 新建评分
		commentScoreRouter.DELETE("deleteCommentScore", commentScoreApi.DeleteCommentScore)           // 删除评分
		commentScoreRouter.DELETE("deleteCommentScoreByIds", commentScoreApi.DeleteCommentScoreByIds) // 批量删除评分
		commentScoreRouter.PUT("updateCommentScore", commentScoreApi.UpdateCommentScore)              // 更新评分
	}
	{
		commentScoreRouterWithoutRecord.GET("findCommentScore", commentScoreApi.FindCommentScore)       // 根据ID获取评分
		commentScoreRouterWithoutRecord.GET("getCommentScoreList", commentScoreApi.GetCommentScoreList) // 获取评分列表
	}
}
