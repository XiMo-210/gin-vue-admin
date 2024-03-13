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
	strconv2 "strconv"
)

type CouponApi struct {
}

var couponService = service.ServiceGroupApp.AdminServiceGroup.CouponService

// CreateCoupon 创建优惠券
// @Tags Coupon
// @Summary 创建优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Coupon true "创建优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /coupon/createCoupon [post]
func (couponApi *CouponApi) CreateCoupon(c *gin.Context) {
	var coupon admin.Coupon
	err := c.ShouldBindJSON(&coupon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	*coupon.RemainCount = *coupon.TotalCount

	if err := couponService.CreateCoupon(&coupon); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCoupon 删除优惠券
// @Tags Coupon
// @Summary 删除优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Coupon true "删除优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /coupon/deleteCoupon [delete]
func (couponApi *CouponApi) DeleteCoupon(c *gin.Context) {
	ID := c.Query("ID")
	if err := couponService.DeleteCoupon(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCouponByIds 批量删除优惠券
// @Tags Coupon
// @Summary 批量删除优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /coupon/deleteCouponByIds [delete]
func (couponApi *CouponApi) DeleteCouponByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := couponService.DeleteCouponByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCoupon 更新优惠券
// @Tags Coupon
// @Summary 更新优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Coupon true "更新优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /coupon/updateCoupon [put]
func (couponApi *CouponApi) UpdateCoupon(c *gin.Context) {
	var coupon admin.Coupon
	err := c.ShouldBindJSON(&coupon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	oldCoupon, err := couponService.GetCoupon(strconv2.Itoa(int(coupon.ID)))
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	*coupon.RemainCount += *coupon.TotalCount - *oldCoupon.TotalCount

	if err := couponService.UpdateCoupon(coupon); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCoupon 用id查询优惠券
// @Tags Coupon
// @Summary 用id查询优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.Coupon true "用id查询优惠券"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /coupon/findCoupon [get]
func (couponApi *CouponApi) FindCoupon(c *gin.Context) {
	ID := c.Query("ID")
	if recoupon, err := couponService.GetCoupon(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recoupon": recoupon}, c)
	}
}

// GetCouponList 分页获取优惠券列表
// @Tags Coupon
// @Summary 分页获取优惠券列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.CouponSearch true "分页获取优惠券列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /coupon/getCouponList [get]
func (couponApi *CouponApi) GetCouponList(c *gin.Context) {
	var pageInfo adminReq.CouponSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, _ := c.Get("claims")
	customClaims, _ := claims.(*request.CustomClaims)

	business, err := businessService.GetBusiness(customClaims.BaseClaims.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	if list, total, err := couponService.GetCouponInfoList(pageInfo, business.ID); err != nil {
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
