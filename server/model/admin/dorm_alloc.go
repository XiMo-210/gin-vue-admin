package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DormAlloc struct {
	global.GVA_MODEL
	StudentInfoId uint
	DormRoomId    uint
}
