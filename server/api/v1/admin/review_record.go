package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type ReviewRecordsApi struct {
}

var reviewRecordService = service.ServiceGroupApp.AdminServiceGroup.ReviewRecordsService

// CreateReviewRecords 创建审核记录
// @Tags ReviewRecords
// @Summary 创建审核记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.ReviewRecords true "创建审核记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /reviewRecord/createReviewRecords [post]
func (reviewRecordApi *ReviewRecordsApi) CreateReviewRecords(c *gin.Context) {
	var reviewRecord admin.ReviewRecords
	err := c.ShouldBindJSON(&reviewRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := reviewRecordService.CreateReviewRecords(&reviewRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteReviewRecords 删除审核记录
// @Tags ReviewRecords
// @Summary 删除审核记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.ReviewRecords true "删除审核记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /reviewRecord/deleteReviewRecords [delete]
func (reviewRecordApi *ReviewRecordsApi) DeleteReviewRecords(c *gin.Context) {
	ID := c.Query("ID")
	if err := reviewRecordService.DeleteReviewRecords(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteReviewRecordsByIds 批量删除审核记录
// @Tags ReviewRecords
// @Summary 批量删除审核记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /reviewRecord/deleteReviewRecordsByIds [delete]
func (reviewRecordApi *ReviewRecordsApi) DeleteReviewRecordsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := reviewRecordService.DeleteReviewRecordsByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateReviewRecords 更新审核记录
// @Tags ReviewRecords
// @Summary 更新审核记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.ReviewRecords true "更新审核记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /reviewRecord/updateReviewRecords [put]
func (reviewRecordApi *ReviewRecordsApi) UpdateReviewRecords(c *gin.Context) {
	var reviewRecord admin.ReviewRecords
	err := c.ShouldBindJSON(&reviewRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userTask, err := userTaskService.GetUserTask(strconv.Itoa(*reviewRecord.UserTaskId))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	if userTask.CurStage != reviewRecord.Stage {
		response.FailWithMessage("用户已完成该阶段, 可删除该申请记录", c)
		return
	}

	task, err := taskService.GetTask(strconv.Itoa(int(userTask.ID)))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	if task.EndTime.Before(time.Now()) {
		response.FailWithMessage("任务时间已截至, 可删除该申请记录", c)
		return
	}

	if err := reviewRecordService.UpdateReviewRecords(reviewRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindReviewRecords 用id查询审核记录
// @Tags ReviewRecords
// @Summary 用id查询审核记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.ReviewRecords true "用id查询审核记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /reviewRecord/findReviewRecords [get]
func (reviewRecordApi *ReviewRecordsApi) FindReviewRecords(c *gin.Context) {
	ID := c.Query("ID")
	if rereviewRecord, err := reviewRecordService.GetReviewRecords(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rereviewRecord": rereviewRecord}, c)
	}
}

// GetReviewRecordsList 分页获取审核记录列表
// @Tags ReviewRecords
// @Summary 分页获取审核记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.ReviewRecordsSearch true "分页获取审核记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /reviewRecord/getReviewRecordsList [get]
func (reviewRecordApi *ReviewRecordsApi) GetReviewRecordsList(c *gin.Context) {
	var pageInfo adminReq.ReviewRecordsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := reviewRecordService.GetReviewRecordsInfoList(pageInfo); err != nil {
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
