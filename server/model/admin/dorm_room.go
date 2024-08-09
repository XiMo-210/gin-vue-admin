// 自动生成模板DormRoom
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 宿舍 结构体  DormRoom
type DormRoom struct {
	global.GVA_MODEL
	DormitoryBuildingId int    `json:"dormitoryBuildingId" form:"dormitoryBuildingId" gorm:"column:dormitory_building_id;comment:;"binding:"required"` //所属寝室楼ID
	RoomNum             int    `json:"roomNum" form:"roomNum" gorm:"column:room_num;comment:;"binding:"required"`                                      //宿舍房间号
	Capacity            int    `json:"capacity" form:"capacity" gorm:"column:capacity;comment:;"binding:"required"`                                    //容纳人数
	Occupied            int    `json:"occupied" form:"occupied" gorm:"column:occupied;comment:;"`                                                      //已入住人数
	Remark              string `json:"remark" form:"remark" gorm:"column:remark;comment:;"`                                                            //备注
}

// TableName 宿舍 DormRoom自定义表名 dorm_rooms
func (DormRoom) TableName() string {
	return "dorm_rooms"
}
