package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CouponRouter struct {
}

// InitCouponRouter 初始化 优惠券 路由信息
func (s *CouponRouter) InitCouponRouter(Router *gin.RouterGroup) {
	couponRouter := Router.Group("coupon").Use(middleware.OperationRecord())
	couponRouterWithoutRecord := Router.Group("coupon")
	var couponApi = v1.ApiGroupApp.AdminApiGroup.CouponApi
	{
		couponRouter.POST("createCoupon", couponApi.CreateCoupon)             // 新建优惠券
		couponRouter.DELETE("deleteCoupon", couponApi.DeleteCoupon)           // 删除优惠券
		couponRouter.DELETE("deleteCouponByIds", couponApi.DeleteCouponByIds) // 批量删除优惠券
		couponRouter.PUT("updateCoupon", couponApi.UpdateCoupon)              // 更新优惠券
	}
	{
		couponRouterWithoutRecord.GET("findCoupon", couponApi.FindCoupon)       // 根据ID获取优惠券
		couponRouterWithoutRecord.GET("getCouponList", couponApi.GetCouponList) // 获取优惠券列表
	}
}
