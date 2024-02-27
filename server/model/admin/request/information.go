package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type InformationSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Category uint8  `json:"category" form:"category" `
	Title    string `json:"title" form:"title" `
	Source   string `json:"source" form:"source" `
	request.PageInfo
}
