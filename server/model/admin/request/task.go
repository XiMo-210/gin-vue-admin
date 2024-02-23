package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TaskSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Category       *int       `json:"category" form:"category" `
	Title          string     `json:"title" form:"title" `
	Campus         string     `json:"campus" form:"campus" `
	College        string     `json:"college" form:"college" `
	NeedMain       *bool      `json:"needMain" form:"needMain" `
	StartStartTime *time.Time `json:"startStartTime" form:"startStartTime"`
	EndStartTime   *time.Time `json:"endStartTime" form:"endStartTime"`
	StartEndTime   *time.Time `json:"startEndTime" form:"startEndTime"`
	EndEndTime     *time.Time `json:"endEndTime" form:"endEndTime"`
	request.PageInfo
}
