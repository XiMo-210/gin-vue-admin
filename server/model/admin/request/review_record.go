package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ReviewRecordsSearch struct {
	Username  string `form:"username"`
	Category  *int   `form:"category"`
	TaskTitle string `form:"taskTitle"`
	Status    *int   `json:"status" form:"status" `
	request.PageInfo
}
