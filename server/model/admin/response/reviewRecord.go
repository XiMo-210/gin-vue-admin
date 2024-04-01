package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	"gorm.io/datatypes"
	"time"
)

type FindReviewRecords struct {
	admin.ReviewRecords
	Username       string         `json:"username"`
	FaceId         string         `json:"faceId"`
	Category       int            `json:"category"`
	TaskTitle      string         `json:"taskTitle"`
	TaskDesc       string         `json:"taskDesc"`
	Campus         string         `json:"campus"`
	College        string         `json:"college"`
	StartTime      time.Time      `json:"startTime"`
	EndTime        time.Time      `json:"endTime"`
	TaskStageTitle string         `json:"taskStageTitle"`
	TaskStageDesc  string         `json:"taskStageDesc"`
	RequiredItem   string         `json:"requiredItem"`
	Imgs           datatypes.JSON `json:"imgs"`
	NeedFace       bool           `json:"needFace"`
	TaskStagePic   string         `json:"taskStagePic"`
}

type GetReviewRecordsList struct {
	admin.ReviewRecords
	Username       string    `json:"username"`
	Category       int       `json:"category"`
	TaskTitle      string    `json:"taskTitle"`
	EndTime        time.Time `json:"endTime"`
	TaskStageTitle string    `json:"taskStageTitle"`
}
