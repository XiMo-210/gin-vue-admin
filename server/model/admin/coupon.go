// 自动生成模板Coupon
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 优惠券 结构体  Coupon
type Coupon struct {
	global.GVA_MODEL
	BusinessId        *int       `json:"businessId" form:"businessId" gorm:"column:business_id;comment:;"binding:"required"`                      //所属店铺商家
	Name              string     `json:"name" form:"name" gorm:"column:name;comment:;"binding:"required"`                                         //优惠券名称
	TotalCount        *int       `json:"totalCount" form:"totalCount" gorm:"column:total_count;comment:;"binding:"required"`                      //总数量
	RemainCount       *int       `json:"remainCount" form:"remainCount" gorm:"column:remain_count;comment:;"binding:"required"`                   //剩余数量
	UsedCount         *int       `json:"usedCount" form:"usedCount" gorm:"column:used_count;comment:;"binding:"required"`                         //已使用数量
	CouponCategory    *int       `json:"couponCategory" form:"couponCategory" gorm:"column:coupon_category;comment:;"binding:"required"`          //优惠券类型
	FullMoney         *int       `json:"fullMoney" form:"fullMoney" gorm:"column:full_money;comment:;"binding:"required"`                         //满
	MinusMoney        *int       `json:"minusMoney" form:"minusMoney" gorm:"column:minus_money;comment:;"binding:"required"`                      //减
	Cash              *int       `json:"cash" form:"cash" gorm:"column:cash;comment:;"binding:"required"`                                         //现金金额
	Discount          *int       `json:"discount" form:"discount" gorm:"column:discount;comment:;"binding:"required"`                             //折扣
	UsageInstructions string     `json:"usageInstructions" form:"usageInstructions" gorm:"column:usage_instructions;comment:;"binding:"required"` //使用说明
	RedeemCount       *int       `json:"redeemCount" form:"redeemCount" gorm:"column:redeem_count;comment:;"binding:"required"`                   //每个用户可兑换数量
	RedeemStartTime   *time.Time `json:"redeemStartTime" form:"redeemStartTime" gorm:"column:redeem_start_time;comment:;"binding:"required"`      //兑换开始时间
	RedeemEndTime     *time.Time `json:"redeemEndTime" form:"redeemEndTime" gorm:"column:redeem_end_time;comment:;"binding:"required"`            //兑换结束时间
	ExpCategory       *int       `json:"expCategory" form:"expCategory" gorm:"column:exp_category;comment:;"binding:"required"`                   //有效期类型
	FixedStartTime    *time.Time `json:"fixedStartTime" form:"fixedStartTime" gorm:"column:fixed_start_time;comment:;"binding:"required"`         //固定开始时间
	FixedEndTime      *time.Time `json:"fixedEndTime" form:"fixedEndTime" gorm:"column:fixed_end_time;comment:;"binding:"required"`               //固定结束时间
	ValidDay          *int       `json:"validDay" form:"validDay" gorm:"column:valid_day;comment:;"binding:"required"`                            //有效天数
	RedeemPoints      *int       `json:"redeemPoints" form:"redeemPoints" gorm:"column:redeem_points;comment:;"binding:"required"`                //兑换所需积分
}

// TableName 优惠券 Coupon自定义表名 coupons
func (Coupon) TableName() string {
	return "coupons"
}
