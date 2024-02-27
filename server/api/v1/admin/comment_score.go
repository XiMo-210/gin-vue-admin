package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommentScoreApi struct {
}

var commentScoreService = service.ServiceGroupApp.AdminServiceGroup.CommentScoreService

// CreateCommentScore 创建评分
// @Tags CommentScore
// @Summary 创建评分
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.CommentScore true "创建评分"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /commentScore/createCommentScore [post]
func (commentScoreApi *CommentScoreApi) CreateCommentScore(c *gin.Context) {
	var commentScore admin.CommentScore
	err := c.ShouldBindJSON(&commentScore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := commentScoreService.CreateCommentScore(&commentScore); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCommentScore 删除评分
// @Tags CommentScore
// @Summary 删除评分
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.CommentScore true "删除评分"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /commentScore/deleteCommentScore [delete]
func (commentScoreApi *CommentScoreApi) DeleteCommentScore(c *gin.Context) {
	ID := c.Query("ID")
	if err := commentScoreService.DeleteCommentScore(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCommentScoreByIds 批量删除评分
// @Tags CommentScore
// @Summary 批量删除评分
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /commentScore/deleteCommentScoreByIds [delete]
func (commentScoreApi *CommentScoreApi) DeleteCommentScoreByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := commentScoreService.DeleteCommentScoreByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCommentScore 更新评分
// @Tags CommentScore
// @Summary 更新评分
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.CommentScore true "更新评分"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /commentScore/updateCommentScore [put]
func (commentScoreApi *CommentScoreApi) UpdateCommentScore(c *gin.Context) {
	var commentScore admin.CommentScore
	err := c.ShouldBindJSON(&commentScore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := commentScoreService.UpdateCommentScore(commentScore); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCommentScore 用id查询评分
// @Tags CommentScore
// @Summary 用id查询评分
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.CommentScore true "用id查询评分"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /commentScore/findCommentScore [get]
func (commentScoreApi *CommentScoreApi) FindCommentScore(c *gin.Context) {
	ID := c.Query("ID")
	if recommentScore, err := commentScoreService.GetCommentScore(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recommentScore": recommentScore}, c)
	}
}

// GetCommentScoreList 分页获取评分列表
// @Tags CommentScore
// @Summary 分页获取评分列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.CommentScoreSearch true "分页获取评分列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /commentScore/getCommentScoreList [get]
func (commentScoreApi *CommentScoreApi) GetCommentScoreList(c *gin.Context) {
	var pageInfo adminReq.CommentScoreSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := commentScoreService.GetCommentScoreInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
