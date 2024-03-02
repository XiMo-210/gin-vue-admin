package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
)

type CouponService struct {
}

// CreateCoupon 创建优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (couponService *CouponService) CreateCoupon(coupon *admin.Coupon) (err error) {
	err = global.GVA_DB.Create(coupon).Error
	return err
}

// DeleteCoupon 删除优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (couponService *CouponService) DeleteCoupon(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.Coupon{}, "id = ?", ID).Error
	return err
}

// DeleteCouponByIds 批量删除优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (couponService *CouponService) DeleteCouponByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.Coupon{}, "id in ?", IDs).Error
	return err
}

// UpdateCoupon 更新优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (couponService *CouponService) UpdateCoupon(coupon admin.Coupon) (err error) {
	err = global.GVA_DB.Save(&coupon).Error
	return err
}

// GetCoupon 根据ID获取优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (couponService *CouponService) GetCoupon(ID string) (coupon admin.Coupon, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&coupon).Error
	return
}

// GetCouponInfoList 分页获取优惠券记录
// Author [piexlmax](https://github.com/piexlmax)
func (couponService *CouponService) GetCouponInfoList(info adminReq.CouponSearch) (list []admin.Coupon, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.Coupon{})
	var coupons []admin.Coupon
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.CouponCategory != nil {
		db = db.Where("coupon_category = ?", info.CouponCategory)
	}
	if info.StartRedeemPoints != nil && info.EndRedeemPoints != nil {
		db = db.Where("redeem_points BETWEEN ? AND ? ", info.StartRedeemPoints, info.EndRedeemPoints)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["total_count"] = true
	orderMap["remain_count"] = true
	orderMap["used_count"] = true
	orderMap["redeem_points"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&coupons).Error
	return coupons, total, err
}
