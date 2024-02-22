package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type StudentInfoSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Name          string `json:"name" form:"name" `
	Gender        string `json:"gender" form:"gender" `
	AdmissionYear *int   `json:"admissionYear" form:"admissionYear" `
	OriginPlace   string `json:"originPlace" form:"originPlace" `
	Campus        string `json:"campus" form:"campus" `
	StudentId     string `json:"studentId" form:"studentId" `
	College       string `json:"college" form:"college" `
	Major         string `json:"major" form:"major" `
	Class         string `json:"class" form:"class" `
	Dormitory     string `json:"dormitory" form:"dormitory" `
	IsRegister    *bool  `json:"isRegister" form:"isRegister" `
	request.PageInfo
}
