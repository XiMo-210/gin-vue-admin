package admin

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type PlatformParamApi struct {
}

var platformParamService = service.ServiceGroupApp.AdminServiceGroup.PlatformParamService

// CreatePlatformParam 创建平台参数
// @Tags PlatformParam
// @Summary 创建平台参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.PlatformParam true "创建平台参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /platformParam/createPlatformParam [post]
func (platformParamApi *PlatformParamApi) CreatePlatformParam(c *gin.Context) {
	var platformParam admin.PlatformParam
	err := c.ShouldBindJSON(&platformParam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	*platformParam.RegisterEndDate = platformParam.RegisterEndDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	*platformParam.TermEndDate = platformParam.TermEndDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	if err := platformParamService.CreatePlatformParam(&platformParam); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePlatformParam 删除平台参数
// @Tags PlatformParam
// @Summary 删除平台参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.PlatformParam true "删除平台参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /platformParam/deletePlatformParam [delete]
func (platformParamApi *PlatformParamApi) DeletePlatformParam(c *gin.Context) {
	ID := c.Query("ID")
	if err := platformParamService.DeletePlatformParam(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePlatformParamByIds 批量删除平台参数
// @Tags PlatformParam
// @Summary 批量删除平台参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /platformParam/deletePlatformParamByIds [delete]
func (platformParamApi *PlatformParamApi) DeletePlatformParamByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := platformParamService.DeletePlatformParamByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePlatformParam 更新平台参数
// @Tags PlatformParam
// @Summary 更新平台参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.PlatformParam true "更新平台参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /platformParam/updatePlatformParam [put]
func (platformParamApi *PlatformParamApi) UpdatePlatformParam(c *gin.Context) {
	var platformParam admin.PlatformParam
	err := c.ShouldBindJSON(&platformParam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	*platformParam.RegisterEndDate = platformParam.RegisterEndDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	*platformParam.TermEndDate = platformParam.TermEndDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	if err := platformParamService.UpdatePlatformParam(platformParam); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPlatformParam 用id查询平台参数
// @Tags PlatformParam
// @Summary 用id查询平台参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.PlatformParam true "用id查询平台参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /platformParam/findPlatformParam [get]
func (platformParamApi *PlatformParamApi) FindPlatformParam(c *gin.Context) {
	if replatformParam, err := platformParamService.GetPlatformParam(); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"replatformParam": replatformParam}, c)
	}
}

// GetPlatformParamList 分页获取平台参数列表
// @Tags PlatformParam
// @Summary 分页获取平台参数列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.PlatformParamSearch true "分页获取平台参数列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /platformParam/getPlatformParamList [get]
func (platformParamApi *PlatformParamApi) GetPlatformParamList(c *gin.Context) {
	var pageInfo adminReq.PlatformParamSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := platformParamService.GetPlatformParamInfoList(pageInfo); err != nil {
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
