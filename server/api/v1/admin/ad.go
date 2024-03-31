package admin

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
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

	if time.Now().Before(*ad.StartDate) {
		*ad.Status = 1
	} else if time.Now().Before(*ad.EndDate) {
		*ad.Status = 2
	} else {
		*ad.Status = 3
	}

	claims, _ := c.Get("claims")
	customClaims, _ := claims.(*request.CustomClaims)

	advertiser, err := advertiserService.GetAdvertiser(customClaims.BaseClaims.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if advertiser.Points < uint(*ad.Cost) {
		response.FailWithMessage("积分不足", c)
		return
	}
	advertiser.Points -= uint(*ad.Cost)
	if err := advertiserService.UpdateAdvertiser(advertiser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
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

	if time.Now().Before(*ad.StartDate) {
		*ad.Status = 1
	} else if time.Now().Before(*ad.EndDate) {
		*ad.Status = 2
	} else {
		*ad.Status = 3
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

	claims, _ := c.Get("claims")
	customClaims, _ := claims.(*request.CustomClaims)

	advertiser, err := advertiserService.GetAdvertiser(customClaims.BaseClaims.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	err = adService.UpdateAdStatus()
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := adService.GetAdInfoList(pageInfo, advertiser.ID); err != nil {
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

func (adApi *AdApi) GetKeywordStats(c *gin.Context) {
	keyword := c.Query("keyword")
	prefixKey := "search_keywords_"

	var today, week float64
	for i := 0; i < 7; i++ {
		key := prefixKey + time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		value, err := global.GVA_REDIS.ZScore(context.Background(), key, keyword).Result()
		if err != nil && err != redis.Nil {
			global.GVA_LOG.Error("获取失败!", zap.Error(err))
			response.FailWithMessage("获取失败", c)
		}
		week += value
		if i == 0 {
			today = value
		}
	}

	response.OkWithDetailed(gin.H{
		"today": today,
		"week":  week,
	}, "获取成功", c)
}
