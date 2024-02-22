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

type WxUserApi struct {
}

var wxUserService = service.ServiceGroupApp.AdminServiceGroup.WxUserService

// CreateWxUser 创建用户
// @Tags WxUser
// @Summary 创建用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.WxUser true "创建用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wxUser/createWxUser [post]
func (wxUserApi *WxUserApi) CreateWxUser(c *gin.Context) {
	var wxUser admin.WxUser
	err := c.ShouldBindJSON(&wxUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wxUserService.CreateWxUser(&wxUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWxUser 删除用户
// @Tags WxUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.WxUser true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wxUser/deleteWxUser [delete]
func (wxUserApi *WxUserApi) DeleteWxUser(c *gin.Context) {
	ID := c.Query("ID")
	if err := wxUserService.DeleteWxUser(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWxUserByIds 批量删除用户
// @Tags WxUser
// @Summary 批量删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wxUser/deleteWxUserByIds [delete]
func (wxUserApi *WxUserApi) DeleteWxUserByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := wxUserService.DeleteWxUserByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWxUser 更新用户
// @Tags WxUser
// @Summary 更新用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.WxUser true "更新用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wxUser/updateWxUser [put]
func (wxUserApi *WxUserApi) UpdateWxUser(c *gin.Context) {
	var wxUser admin.WxUser
	err := c.ShouldBindJSON(&wxUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := wxUserService.UpdateWxUser(wxUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWxUser 用id查询用户
// @Tags WxUser
// @Summary 用id查询用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.WxUser true "用id查询用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wxUser/findWxUser [get]
func (wxUserApi *WxUserApi) FindWxUser(c *gin.Context) {
	ID := c.Query("ID")
	if rewxUser, err := wxUserService.GetWxUser(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewxUser": rewxUser}, c)
	}
}

// GetWxUserList 分页获取用户列表
// @Tags WxUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.WxUserSearch true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wxUser/getWxUserList [get]
func (wxUserApi *WxUserApi) GetWxUserList(c *gin.Context) {
	var pageInfo adminReq.WxUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := wxUserService.GetWxUserInfoList(pageInfo); err != nil {
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
