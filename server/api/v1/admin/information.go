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

type InformationApi struct {
}

var informationService = service.ServiceGroupApp.AdminServiceGroup.InformationService

// CreateInformation 创建信息
// @Tags Information
// @Summary 创建信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Information true "创建信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /information/createInformation [post]
func (informationApi *InformationApi) CreateInformation(c *gin.Context) {
	var information admin.Information
	err := c.ShouldBindJSON(&information)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := informationService.CreateInformation(&information); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteInformation 删除信息
// @Tags Information
// @Summary 删除信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Information true "删除信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /information/deleteInformation [delete]
func (informationApi *InformationApi) DeleteInformation(c *gin.Context) {
	ID := c.Query("ID")
	if err := informationService.DeleteInformation(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteInformationByIds 批量删除信息
// @Tags Information
// @Summary 批量删除信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /information/deleteInformationByIds [delete]
func (informationApi *InformationApi) DeleteInformationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := informationService.DeleteInformationByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateInformation 更新信息
// @Tags Information
// @Summary 更新信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Information true "更新信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /information/updateInformation [put]
func (informationApi *InformationApi) UpdateInformation(c *gin.Context) {
	var information admin.Information
	err := c.ShouldBindJSON(&information)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := informationService.UpdateInformation(information); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindInformation 用id查询信息
// @Tags Information
// @Summary 用id查询信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.Information true "用id查询信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /information/findInformation [get]
func (informationApi *InformationApi) FindInformation(c *gin.Context) {
	ID := c.Query("ID")
	if reinformation, err := informationService.GetInformation(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinformation": reinformation}, c)
	}
}

// GetInformationList 分页获取信息列表
// @Tags Information
// @Summary 分页获取信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.InformationSearch true "分页获取信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /information/getInformationList [get]
func (informationApi *InformationApi) GetInformationList(c *gin.Context) {
	var pageInfo adminReq.InformationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := informationService.GetInformationInfoList(pageInfo); err != nil {
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
