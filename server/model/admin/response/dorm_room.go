package response

type GetDormRoomList struct {
	ID       uint   `json:"ID" form:"ID" gorm:"column:id"`
	Campus   string `json:"campus" form:"campus" gorm:"column:campus;"`
	Name     string `json:"name" form:"name" gorm:"column:name;"`
	Gender   string `json:"gender" form:"gender" gorm:"column:gender;"`
	RoomNum  int    `json:"roomNum" form:"roomNum" gorm:"column:room_num;"`
	Capacity int    `json:"capacity" form:"capacity" gorm:"column:capacity;"`
	Occupied int    `json:"occupied" form:"occupied" gorm:"column:occupied;"`
	Remark   string `json:"remark" form:"remark" gorm:"column:remark;"`
}
