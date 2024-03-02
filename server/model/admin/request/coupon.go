package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CouponSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Name              string `json:"name" form:"name" `
	CouponCategory    *int   `json:"couponCategory" form:"couponCategory" `
	StartRedeemPoints *int   `json:"startRedeemPoints" form:"startRedeemPoints"`
	EndRedeemPoints   *int   `json:"endRedeemPoints" form:"endRedeemPoints"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
