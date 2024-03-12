package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/datatypes"
	"time"
)

type AdSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Name         string `json:"name" form:"name" `
	CostCategory *int   `json:"costCategory" form:"costCategory" `
	Status       *int   `json:"status" form:"status" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
