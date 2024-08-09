package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type QuestionnaireRouter struct {
}

// InitQuestionnaireRouter 初始化 寝室分配问卷 路由信息
func (s *QuestionnaireRouter) InitQuestionnaireRouter(Router *gin.RouterGroup) {
	questionnaireRouter := Router.Group("questionnaire").Use(middleware.OperationRecord())
	questionnaireRouterWithoutRecord := Router.Group("questionnaire")
	var questionnaireApi = v1.ApiGroupApp.AdminApiGroup.QuestionnaireApi
	{
		questionnaireRouter.POST("createQuestionnaire", questionnaireApi.CreateQuestionnaire)   // 新建寝室分配问卷
		questionnaireRouter.DELETE("deleteQuestionnaire", questionnaireApi.DeleteQuestionnaire) // 删除寝室分配问卷
		questionnaireRouter.DELETE("deleteQuestionnaireByIds", questionnaireApi.DeleteQuestionnaireByIds) // 批量删除寝室分配问卷
		questionnaireRouter.PUT("updateQuestionnaire", questionnaireApi.UpdateQuestionnaire)    // 更新寝室分配问卷
	}
	{
		questionnaireRouterWithoutRecord.GET("findQuestionnaire", questionnaireApi.FindQuestionnaire)        // 根据ID获取寝室分配问卷
		questionnaireRouterWithoutRecord.GET("getQuestionnaireList", questionnaireApi.GetQuestionnaireList)  // 获取寝室分配问卷列表
	}
}
