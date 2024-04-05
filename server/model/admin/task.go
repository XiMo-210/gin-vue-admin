// 自动生成模板Task
package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
	"time"
)

type TaskWithStages struct {
	Task       Task        `json:"task" binding:"required"`
	TaskStages []TaskStage `json:"taskStages" binding:"required"`
}

// 任务 结构体  Task
type Task struct {
	global.GVA_MODEL
	Category  *int       `json:"category" form:"category" gorm:"column:category;comment:;"binding:"required"`     //任务类型
	Title     string     `json:"title" form:"title" gorm:"column:title;comment:;"binding:"required"`              //任务名称
	Desc      string     `json:"desc" form:"desc" gorm:"column:desc;comment:;"binding:"required"`                 //任务描述
	Campus    string     `json:"campus" form:"campus" gorm:"column:campus;comment:;"binding:"required"`           //任务校区
	College   string     `json:"college" form:"college" gorm:"column:college;comment:;"binding:"required"`        //任务学院
	Reward    *int       `json:"reward" form:"reward" gorm:"column:reward;comment:;"binding:"required"`           //奖励积分
	NeedMain  *bool      `json:"needMain" form:"needMain" gorm:"column:need_main;comment:;"binding:"required"`    //是否需要完成主线任务
	StartTime *time.Time `json:"startTime" form:"startTime" gorm:"column:start_time;comment:;"binding:"required"` //开始时间
	EndTime   *time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:;"binding:"required"`       //结束时间
	StageNum  int        `json:"stageNum" form:"stageNum" gorm:"column:stage_num;comment:;"`                      //阶段数
}

// TableName 任务 Task自定义表名 tasks
func (Task) TableName() string {
	return "tasks"
}

// 任务阶段 结构体  TaskStage
type TaskStage struct {
	global.GVA_MODEL
	TaskId       uint           `json:"taskId" form:"taskId" gorm:"column:task_id;comment:;"`                                 //所属任务ID
	Stage        *int           `json:"stage" form:"stage" gorm:"column:stage;comment:;"binding:"required"`                   //阶段
	Title        string         `json:"title" form:"title" gorm:"column:title;comment:;"binding:"required"`                   //阶段任务名称
	Desc         string         `json:"desc" form:"desc" gorm:"column:desc;comment:;"`                                        //阶段任务描述
	RequiredItem string         `json:"requiredItem" form:"requiredItem" gorm:"column:required_item;comment:;"`               //需要准备的物品
	Imgs         datatypes.JSON `json:"imgs" form:"imgs" gorm:"column:imgs;comment:;"`                                        //图片
	NeedPic      *bool          `json:"needPic" form:"needPic" gorm:"column:need_pic;comment:;"binding:"required"`            //是否需要上传指定内容图片
	NeedFace     *bool          `json:"needFace" form:"needFace" gorm:"column:need_face;comment:;"binding:"required"`         //图片是否需要人脸
	Pic          string         `json:"pic" form:"pic" gorm:"column:pic;comment:;"binding:"required"`                         //指定内容图片
	NeedLoc      *bool          `json:"needLoc" form:"needLoc" gorm:"column:need_loc;comment:;"binding:"required"`            //是否需要指定位置
	AllowDist    *int           `json:"allowDist" form:"allowDist" gorm:"column:allow_dist;comment:;"binding:"required"`      //允许距离
	NeedNav      *bool          `json:"needNav" form:"needNav" gorm:"column:need_nav;comment:;"binding:"required"`            //是否需要导航
	Loc          string         `json:"loc" form:"loc" gorm:"column:loc;comment:;"binding:"required"`                         //指定位置
	NeedCamera   *bool          `json:"needCamera" form:"needCamera" gorm:"column:need_camera;comment:;"binding:"required"`   //是否关联摄像头
	CameraId     *int           `json:"cameraId" form:"cameraId" gorm:"column:camera_id;comment:;size:20;"binding:"required"` //摄像头编号
}

// TableName 任务阶段 TaskStage自定义表名 task_stages
func (TaskStage) TableName() string {
	return "task_stages"
}
