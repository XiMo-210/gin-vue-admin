package response

import "time"

type CompleteRecords struct {
	UserTaskId  int       `json:"userTaskId"`
	UserId      int       `json:"userId"`
	Avatar      string    `json:"avatar"`
	Username    string    `json:"username"`
	CompletedAt time.Time `json:"completedAt"`
}
