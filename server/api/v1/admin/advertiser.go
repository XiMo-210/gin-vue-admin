package admin

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AdvertiserApi struct {
}

var advertiserService = service.ServiceGroupApp.AdminServiceGroup.AdvertiserService

// CreateAdvertiser 创建广告主
// @Tags Advertiser
// @Summary 创建广告主
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Advertiser true "创建广告主"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /advertiser/createAdvertiser [post]
func (advertiserApi *AdvertiserApi) CreateAdvertiser(c *gin.Context) {
	var advertiser admin.Advertiser
	err := c.ShouldBindJSON(&advertiser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, _ := c.Get("claims")
	customClaims, _ := claims.(*request.CustomClaims)
	advertiser.SysUserId = customClaims.BaseClaims.ID

	if err := advertiserService.CreateAdvertiser(&advertiser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAdvertiser 删除广告主
// @Tags Advertiser
// @Summary 删除广告主
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Advertiser true "删除广告主"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /advertiser/deleteAdvertiser [delete]
func (advertiserApi *AdvertiserApi) DeleteAdvertiser(c *gin.Context) {
	ID := c.Query("ID")
	if err := advertiserService.DeleteAdvertiser(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAdvertiserByIds 批量删除广告主
// @Tags Advertiser
// @Summary 批量删除广告主
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /advertiser/deleteAdvertiserByIds [delete]
func (advertiserApi *AdvertiserApi) DeleteAdvertiserByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := advertiserService.DeleteAdvertiserByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAdvertiser 更新广告主
// @Tags Advertiser
// @Summary 更新广告主
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Advertiser true "更新广告主"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /advertiser/updateAdvertiser [put]
func (advertiserApi *AdvertiserApi) UpdateAdvertiser(c *gin.Context) {
	var advertiser admin.Advertiser
	err := c.ShouldBindJSON(&advertiser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := advertiserService.UpdateAdvertiser(advertiser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAdvertiser 用id查询广告主
// @Tags Advertiser
// @Summary 用id查询广告主
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.Advertiser true "用id查询广告主"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /advertiser/findAdvertiser [get]
func (advertiserApi *AdvertiserApi) FindAdvertiser(c *gin.Context) {
	claims, _ := c.Get("claims")
	customClaims, _ := claims.(*request.CustomClaims)

	if readvertiser, err := advertiserService.GetAdvertiser(customClaims.BaseClaims.ID); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"readvertiser": readvertiser}, c)
	}
}

// GetAdvertiserList 分页获取广告主列表
// @Tags Advertiser
// @Summary 分页获取广告主列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.AdvertiserSearch true "分页获取广告主列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /advertiser/getAdvertiserList [get]
func (advertiserApi *AdvertiserApi) GetAdvertiserList(c *gin.Context) {
	var pageInfo adminReq.AdvertiserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := advertiserService.GetAdvertiserInfoList(pageInfo); err != nil {
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

func (advertiserApi *AdvertiserApi) CreateAdvertiserByAdmin(c *gin.Context) {
	var advertiser admin.Advertiser
	err := c.ShouldBindJSON(&advertiser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := advertiserService.CreateAdvertiser(&advertiser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (advertiserApi *AdvertiserApi) UpdateAdvertiserByAdmin(c *gin.Context) {
	var advertiser admin.Advertiser
	err := c.ShouldBindJSON(&advertiser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := advertiserService.UpdateAdvertiser(advertiser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (advertiserApi *AdvertiserApi) FindAdvertiserByAdmin(c *gin.Context) {
	ID := c.Query("ID")
	if readvertiser, err := advertiserService.GetAdvertiserByAdmin(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"readvertiser": readvertiser}, c)
	}
}
