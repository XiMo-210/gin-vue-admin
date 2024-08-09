package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type DormRoomSearch struct {
	Campus   *string `json:"campus" form:"campus"`
	Gender   *string `json:"gender" form:"gender"`
	Name     *string `json:"name" form:"name"`
	RoomNum  *int    `json:"roomNum" form:"roomNum"`
	Capacity *int    `json:"capacity" form:"capacity"`
	Occupied *int    `json:"occupied" form:"occupied"`
	request.PageInfo
}

type CreateDormRoom struct {
	global.GVA_MODEL
	DormitoryBuildingId int  `json:"dormitoryBuildingId" form:"dormitoryBuildingId" binding:"required"`
	RoomNumStart        int  `json:"roomNumStart" form:"roomNumStart" binding:"required"`
	RoomNumEnd          int  `json:"roomNumEnd" form:"roomNumEnd" binding:"required"`
	Capacity            int  `json:"capacity" form:"capacity" binding:"required"`
	Occupied            *int `json:"occupied" form:"occupied" binding:"required"`
}
