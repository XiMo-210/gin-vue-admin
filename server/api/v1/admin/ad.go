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

type AdApi struct {
}

var adService = service.ServiceGroupApp.AdminServiceGroup.AdService

// CreateAd 创建广告
// @Tags Ad
// @Summary 创建广告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Ad true "创建广告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ad/createAd [post]
func (adApi *AdApi) CreateAd(c *gin.Context) {
	var ad admin.Ad
	err := c.ShouldBindJSON(&ad)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := adService.CreateAd(&ad); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAd 删除广告
// @Tags Ad
// @Summary 删除广告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Ad true "删除广告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ad/deleteAd [delete]
func (adApi *AdApi) DeleteAd(c *gin.Context) {
	ID := c.Query("ID")
	if err := adService.DeleteAd(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAdByIds 批量删除广告
// @Tags Ad
// @Summary 批量删除广告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ad/deleteAdByIds [delete]
func (adApi *AdApi) DeleteAdByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := adService.DeleteAdByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAd 更新广告
// @Tags Ad
// @Summary 更新广告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Ad true "更新广告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ad/updateAd [put]
func (adApi *AdApi) UpdateAd(c *gin.Context) {
	var ad admin.Ad
	err := c.ShouldBindJSON(&ad)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := adService.UpdateAd(ad); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAd 用id查询广告
// @Tags Ad
// @Summary 用id查询广告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.Ad true "用id查询广告"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ad/findAd [get]
func (adApi *AdApi) FindAd(c *gin.Context) {
	ID := c.Query("ID")
	if read, err := adService.GetAd(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"read": read}, c)
	}
}

// GetAdList 分页获取广告列表
// @Tags Ad
// @Summary 分页获取广告列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.AdSearch true "分页获取广告列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ad/getAdList [get]
func (adApi *AdApi) GetAdList(c *gin.Context) {
	var pageInfo adminReq.AdSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := adService.GetAdInfoList(pageInfo); err != nil {
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
