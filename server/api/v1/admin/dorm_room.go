package admin

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/subscribeMessage/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
	adminResp "github.com/flipped-aurora/gin-vue-admin/server/model/admin/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"os"
	"sort"
)

type DormRoomApi struct {
}

var dormRoomService = service.ServiceGroupApp.AdminServiceGroup.DormRoomService

// CreateDormRoom 创建宿舍
// @Tags DormRoom
// @Summary 创建宿舍
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.DormRoom true "创建宿舍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dormRoom/createDormRoom [post]
func (dormRoomApi *DormRoomApi) CreateDormRoom(c *gin.Context) {
	var dormRoom adminReq.CreateDormRoom
	err := c.ShouldBindJSON(&dormRoom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var rooms []admin.DormRoom
	for i := dormRoom.RoomNumStart; i <= dormRoom.RoomNumEnd; i++ {
		rooms = append(rooms, admin.DormRoom{
			DormitoryBuildingId: dormRoom.DormitoryBuildingId,
			RoomNum:             i,
			Capacity:            dormRoom.Capacity,
			Occupied:            *dormRoom.Occupied,
		})
	}

	if err := dormRoomService.CreateDormRoom(rooms); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDormRoom 删除宿舍
// @Tags DormRoom
// @Summary 删除宿舍
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.DormRoom true "删除宿舍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dormRoom/deleteDormRoom [delete]
func (dormRoomApi *DormRoomApi) DeleteDormRoom(c *gin.Context) {
	ID := c.Query("ID")
	if err := dormRoomService.DeleteDormRoom(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDormRoomByIds 批量删除宿舍
// @Tags DormRoom
// @Summary 批量删除宿舍
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dormRoom/deleteDormRoomByIds [delete]
func (dormRoomApi *DormRoomApi) DeleteDormRoomByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := dormRoomService.DeleteDormRoomByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDormRoom 更新宿舍
// @Tags DormRoom
// @Summary 更新宿舍
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.DormRoom true "更新宿舍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dormRoom/updateDormRoom [put]
func (dormRoomApi *DormRoomApi) UpdateDormRoom(c *gin.Context) {
	var dormRoom admin.DormRoom
	err := c.ShouldBindJSON(&dormRoom)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dormRoomService.UpdateDormRoom(dormRoom); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDormRoom 用id查询宿舍
// @Tags DormRoom
// @Summary 用id查询宿舍
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.DormRoom true "用id查询宿舍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dormRoom/findDormRoom [get]
func (dormRoomApi *DormRoomApi) FindDormRoom(c *gin.Context) {
	ID := c.Query("ID")
	if redormRoom, err := dormRoomService.GetDormRoom(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redormRoom": redormRoom}, c)
	}
}

// GetDormRoomList 分页获取宿舍列表
// @Tags DormRoom
// @Summary 分页获取宿舍列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.DormRoomSearch true "分页获取宿舍列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dormRoom/getDormRoomList [get]
func (dormRoomApi *DormRoomApi) GetDormRoomList(c *gin.Context) {
	var pageInfo adminReq.DormRoomSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dormRoomService.GetDormRoomInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (dormRoomApi *DormRoomApi) GetDormStudents(c *gin.Context) {
	ID := c.Query("ID")
	var dormAllocs []admin.DormAlloc
	global.GVA_DB.Model(&admin.DormAlloc{}).Where("dorm_room_id = ?", ID).Find(&dormAllocs)
	var ids []uint
	for _, v := range dormAllocs {
		ids = append(ids, v.StudentInfoId)
	}
	var students []admin.StudentInfo
	global.GVA_DB.Model(&admin.StudentInfo{}).Where("id IN ?", ids).Find(&students)

	response.OkWithData(students, c)
}

type Param struct {
	ID      uint   `gorm:"column:id"`
	Name    string `gorm:"column:name"`
	Gender  string `gorm:"column:gender"`
	Campus  string `gorm:"column:campus"`
	College string `gorm:"column:college"`
	Major   string `gorm:"column:major"`
	Class   string `gorm:"column:class"`
	Param1  uint8  `gorm:"column:param1"`
	Param2  uint8  `gorm:"column:param2"`
	Param3  uint8  `gorm:"column:param3"`
	Param4  uint8  `gorm:"column:param4"`
	Param5  uint8  `gorm:"column:param5"`
}

type Result struct {
	Female []Alloc `json:"female"`
	Male   []Alloc `json:"male"`
}

type Alloc struct {
	ID      uint   `json:"id"`
	Campus  string `json:"campus"`
	DormNum uint   `json:"dormNum"`
}

type Key struct {
	Campus string
	Gender string
}

func (dormRoomApi *DormRoomApi) DormAlloc(c *gin.Context) {
	var data []Param
	if global.GVA_DB.Model(&admin.StudentInfo{}).Select("student_infos.id, student_infos.name, student_infos.gender, student_infos.campus, student_infos.college, student_infos.major, student_infos.class, questionnaires.param1, questionnaires.param2, questionnaires.param3, questionnaires.param4, questionnaires.param5").
		Joins("LEFT JOIN questionnaires ON student_infos.id = questionnaires.student_info_id").Find(&data).Error != nil {
		response.FailWithMessage("分配失败", c)
		return
	}

	// 创建 CSV 文件
	file, err := os.Create("temp.csv")
	if err != nil {
		response.FailWithMessage("分配失败", c)
		fmt.Println("Error creating CSV file:", err)
		return
	}

	// 创建 CSV 写入器
	writer := csv.NewWriter(file)

	// 准备头部
	header := []string{
		"ID", "Name", "Gender", "Campus", "College", "Major", "Class", "Param1", "Param2", "Param3", "Param4", "Param5",
	}

	// 写入头部
	err = writer.Write(header)
	if err != nil {
		response.FailWithMessage("分配失败", c)
		fmt.Println("error writing header to CSV file:", err)
		return
	}

	// 遍历 Param 结构体切片，并写入 CSV
	for _, param := range data {
		row := []string{
			fmt.Sprintf("%d", param.ID),
			param.Name,
			param.Gender,
			param.Campus,
			param.College,
			param.Major,
			param.Class,
			fmt.Sprintf("%d", param.Param1),
			fmt.Sprintf("%d", param.Param2),
			fmt.Sprintf("%d", param.Param3),
			fmt.Sprintf("%d", param.Param4),
			fmt.Sprintf("%d", param.Param5),
		}
		err = writer.Write(row)
		if err != nil {
			response.FailWithMessage("分配失败", c)
			fmt.Println("error writing row to CSV file:", err)
		}
	}

	writer.Flush()

	// 创建 resty 客户端
	client := resty.New()

	// 设置文件路径
	filePath := "temp.csv"

	// 创建一个文件读取器
	fileReader, err := os.Open(filePath)
	if err != nil {
		response.FailWithMessage("分配失败", c)
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileReader.Close()

	result := &Result{}
	// 发送 POST 请求
	_, err = client.R().
		SetResult(result).
		SetFile("file", filePath).
		Post("http://localhost:8123/api/dorm-alloc")
	// 检查错误
	if err != nil {
		response.FailWithMessage("分配失败", c)
		fmt.Println("Error sending request:", err)
		return
	}

	// 创建一个映射，以存储每个 DormNum 对应的学生列表
	dormAllocMap := make(map[Key]map[uint][]uint)

	// 将学生按 DormNum 分组
	for gender, allocs := range map[string][]Alloc{"female": result.Female, "male": result.Male} {
		for _, alloc := range allocs {
			key := Key{
				Campus: alloc.Campus,
				Gender: gender,
			}
			if dormAllocMap[key] == nil {
				dormAllocMap[key] = make(map[uint][]uint)
			}
			if dormAllocMap[key][alloc.DormNum] == nil {
				dormAllocMap[key][alloc.DormNum] = make([]uint, 0)
			}
			dormAllocMap[key][alloc.DormNum] = append(dormAllocMap[key][alloc.DormNum], alloc.ID)
		}
	}

	var dormRooms []adminResp.GetDormRoomList
	if err := global.GVA_DB.Model(&admin.DormRoom{}).Select("dormitory_buildings.campus, dormitory_buildings.name, dormitory_buildings.gender, dorm_rooms.*").Joins("LEFT JOIN dormitory_buildings ON dormitory_buildings.id = dorm_rooms.dormitory_building_id").Find(&dormRooms).Error; err != nil {
		response.FailWithMessage("分配失败", c)
		return
	}

	// 创建一个映射，以存储每个 DormNum 对应的学生列表
	dormToAllocMap := make(map[Key][]uint)
	// 将学生按 DormNum 分组
	for _, room := range dormRooms {
		if room.Gender == "男" {
			room.Gender = "male"
		} else {
			room.Gender = "female"
		}
		key := Key{
			Campus: room.Campus,
			Gender: room.Gender,
		}
		if dormToAllocMap[key] == nil {
			dormToAllocMap[key] = make([]uint, 0)
		}
		dormToAllocMap[key] = append(dormToAllocMap[key], room.ID)
	}

	count := make(map[Key]uint)
	count[Key{
		Campus: "朝晖校区",
		Gender: "male",
	}] = 0
	count[Key{
		Campus: "朝晖校区",
		Gender: "female",
	}] = 0
	count[Key{
		Campus: "屏峰校区",
		Gender: "male",
	}] = 0
	count[Key{
		Campus: "屏峰校区",
		Gender: "female",
	}] = 0

	var allocs []admin.DormAlloc
	var dormed []uint
	for key, dormAllocs := range dormAllocMap {
		ks := make([]int, 0, len(dormAllocs))
		for k := range dormAllocs {
			ks = append(ks, int(k))
		}
		sort.Ints(ks)
		for _, v := range ks {
			for _, studentInfoId := range dormAllocs[uint(v)] {
				allocs = append(allocs, admin.DormAlloc{
					StudentInfoId: studentInfoId,
					DormRoomId:    dormToAllocMap[key][count[key]],
				})
			}
			dormed = append(dormed, dormToAllocMap[key][count[key]])
			count[key]++
		}
	}

	global.GVA_DB.Model(&admin.DormRoom{}).Where("id in ?", dormed).Update("occupied", 4)

	if err := global.GVA_DB.Create(&allocs).Error; err != nil {
		response.FailWithMessage("分配失败", c)
		return
	}

	data2 := &power.HashMap{
		// 问卷名
		"thing1": power.StringMap{
			"value": "寝室分配",
		},
		// 截止时间
		"phrase4": power.StringMap{
			"value": "已完成",
		},
		// 备注
		"thing9": power.StringMap{
			"value": "您的寝室分配结果已出炉, 点击查看室友。",
		},
	}

	mpResp, err := global.MiniProgram.SubscribeMessage.Send(context.Background(), &request.RequestSubscribeMessageSend{
		ToUser:           "oBgza4gX9b7MioFdrzWWOOu9yJnE",
		TemplateID:       "Ee-ANBeWeQougnSw8ipyoEW-SkRYBeWvZNgLE4GllJo",
		Page:             "pages/mine/detail",
		MiniProgramState: global.GVA_CONFIG.MiniProgram.State,
		Lang:             "zh_CN",
		Data:             data2,
	})
	if err != nil || mpResp.ErrCode != 0 {
		global.GVA_LOG.Error("消息推送失败!", zap.Error(err))
		response.FailWithMessage("消息推送失败", c)
		return
	}

	response.Ok(c)
}

func (dormRoomApi *DormRoomApi) Remind(c *gin.Context) {
	data := &power.HashMap{
		// 问卷名
		"thing1": power.StringMap{
			"value": "住宿习惯调查问卷",
		},
		// 截止时间
		"time2": power.StringMap{
			"value": "2024年8月9日",
		},
		// 备注
		"thing4": power.StringMap{
			"value": "您还未填写寝室分配问卷, 请尽快填写",
		},
	}

	mpResp, err := global.MiniProgram.SubscribeMessage.Send(context.Background(), &request.RequestSubscribeMessageSend{
		ToUser:           "oBgza4gX9b7MioFdrzWWOOu9yJnE",
		TemplateID:       "PdvAqfTP3H2qWep-H9QkKYXnf5cLb4Pk-_2JRWhkMEg",
		Page:             "pages/mine/roomDistribution",
		MiniProgramState: global.GVA_CONFIG.MiniProgram.State,
		Lang:             "zh_CN",
		Data:             data,
	})
	if err != nil || mpResp.ErrCode != 0 {
		global.GVA_LOG.Error("消息推送失败!", zap.Error(err))
		response.FailWithMessage("消息推送失败", c)
		return
	}

	response.Ok(c)
}
