package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type WxUserSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	StudentInfoId   *int   `json:"studentInfoId" form:"studentInfoId" `
	Username        string `json:"username" form:"username" `
	StartExp        *int   `json:"startExp" form:"startExp"`
	EndExp          *int   `json:"endExp" form:"endExp"`
	StartPoints     *int   `json:"startPoints" form:"startPoints"`
	EndPoints       *int   `json:"endPoints" form:"endPoints"`
	IsCompletedMain *bool  `json:"isCompletedMain" form:"isCompletedMain" `
	CurTask         *int   `json:"curTask" form:"curTask" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
