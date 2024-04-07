package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TaskSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Category  *int       `json:"category" form:"category"`
	Title     string     `json:"title" form:"title"`
	Campus    string     `json:"campus" form:"campus"`
	College   string     `json:"college" form:"college"`
	NeedMain  *bool      `json:"needMain" form:"needMain"`
	StartTime *time.Time `json:"startTime" form:"startTime"`
	EndTime   *time.Time `json:"endTime" form:"endTime"`
	request.PageInfo
}

type CompleteRecords struct {
	TaskId   *int   `form:"taskId"`
	UserId   *int   `form:"userId"`
	Username string `form:"username"`
	request.PageInfo
}
