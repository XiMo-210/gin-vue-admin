package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type DormitoryBuildingSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Campus         string `json:"campus" form:"campus" `
	Name           string `json:"name" form:"name" `
	Gender         string `json:"gender" form:"gender" `
	ManagerName    string `json:"managerName" form:"managerName" `
	ManagerContact string `json:"managerContact" form:"managerContact" `
	request.PageInfo
}
