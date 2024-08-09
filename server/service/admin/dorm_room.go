package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin/response"
)

type DormRoomService struct {
}

// CreateDormRoom 创建宿舍记录
// Author [piexlmax](https://github.com/piexlmax)
func (dormRoomService *DormRoomService) CreateDormRoom(dormRoom []admin.DormRoom) (err error) {
	err = global.GVA_DB.Create(&dormRoom).Error
	return err
}

// DeleteDormRoom 删除宿舍记录
// Author [piexlmax](https://github.com/piexlmax)
func (dormRoomService *DormRoomService) DeleteDormRoom(ID string) (err error) {
	err = global.GVA_DB.Delete(&admin.DormRoom{}, "id = ?", ID).Error
	return err
}

// DeleteDormRoomByIds 批量删除宿舍记录
// Author [piexlmax](https://github.com/piexlmax)
func (dormRoomService *DormRoomService) DeleteDormRoomByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]admin.DormRoom{}, "id in ?", IDs).Error
	return err
}

// UpdateDormRoom 更新宿舍记录
// Author [piexlmax](https://github.com/piexlmax)
func (dormRoomService *DormRoomService) UpdateDormRoom(dormRoom admin.DormRoom) (err error) {
	err = global.GVA_DB.Save(&dormRoom).Error
	return err
}

// GetDormRoom 根据ID获取宿舍记录
// Author [piexlmax](https://github.com/piexlmax)
func (dormRoomService *DormRoomService) GetDormRoom(ID string) (dormRoom admin.DormRoom, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dormRoom).Error
	return
}

// GetDormRoomInfoList 分页获取宿舍记录
// Author [piexlmax](https://github.com/piexlmax)
func (dormRoomService *DormRoomService) GetDormRoomInfoList(info adminReq.DormRoomSearch) (list []response.GetDormRoomList, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&admin.DormRoom{}).Select("dormitory_buildings.campus, dormitory_buildings.name, dormitory_buildings.gender, dorm_rooms.*").Joins("LEFT JOIN dormitory_buildings ON dormitory_buildings.id = dorm_rooms.dormitory_building_id")
	var dormRooms []response.GetDormRoomList
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Campus != nil {
		db = db.Where("campus = ?", info.Campus)
	}
	if info.Gender != nil {
		db = db.Where("gender = ?", info.Gender)
	}
	if info.Name != nil {
		db = db.Where("name = ?", info.Name)
	}
	if info.RoomNum != nil {
		db = db.Where("room_num = ?", info.RoomNum)
	}
	if info.Capacity != nil {
		db = db.Where("capacity = ?", info.Capacity)
	}
	if info.Occupied != nil {
		db = db.Where("occupied = ?", info.Occupied)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&dormRooms).Error
	return dormRooms, total, err
}
