package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/admin"
	"time"
)

type GetUserTaskCondition struct {
	CurTask        admin.TaskWithStages `json:"curTask"`
	CurUserTask    admin.UserTask       `json:"curUserTask"`
	CompletedTasks []CompletedTask      `json:"completedTasks"`
}

type CompletedTask struct {
	Category   uint8     `json:"category"`
	Title      string    `json:"title"`
	Desc       string    `json:"desc"`
	Campus     string    `json:"campus"`
	College    string    `json:"college"`
	Reward     uint      `json:"reward"`
	FinishTime time.Time `json:"finishTime" gorm:"column:updated_at"`
}
