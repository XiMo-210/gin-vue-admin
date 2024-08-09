package admin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/admin"
    adminReq "github.com/flipped-aurora/gin-vue-admin/server/model/admin/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type QuestionnaireApi struct {
}

var questionnaireService = service.ServiceGroupApp.AdminServiceGroup.QuestionnaireService


// CreateQuestionnaire 创建寝室分配问卷
// @Tags Questionnaire
// @Summary 创建寝室分配问卷
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Questionnaire true "创建寝室分配问卷"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /questionnaire/createQuestionnaire [post]
func (questionnaireApi *QuestionnaireApi) CreateQuestionnaire(c *gin.Context) {
	var questionnaire admin.Questionnaire
	err := c.ShouldBindJSON(&questionnaire)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := questionnaireService.CreateQuestionnaire(&questionnaire); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQuestionnaire 删除寝室分配问卷
// @Tags Questionnaire
// @Summary 删除寝室分配问卷
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Questionnaire true "删除寝室分配问卷"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionnaire/deleteQuestionnaire [delete]
func (questionnaireApi *QuestionnaireApi) DeleteQuestionnaire(c *gin.Context) {
	ID := c.Query("ID")
	if err := questionnaireService.DeleteQuestionnaire(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQuestionnaireByIds 批量删除寝室分配问卷
// @Tags Questionnaire
// @Summary 批量删除寝室分配问卷
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /questionnaire/deleteQuestionnaireByIds [delete]
func (questionnaireApi *QuestionnaireApi) DeleteQuestionnaireByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := questionnaireService.DeleteQuestionnaireByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQuestionnaire 更新寝室分配问卷
// @Tags Questionnaire
// @Summary 更新寝室分配问卷
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body admin.Questionnaire true "更新寝室分配问卷"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionnaire/updateQuestionnaire [put]
func (questionnaireApi *QuestionnaireApi) UpdateQuestionnaire(c *gin.Context) {
	var questionnaire admin.Questionnaire
	err := c.ShouldBindJSON(&questionnaire)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := questionnaireService.UpdateQuestionnaire(questionnaire); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQuestionnaire 用id查询寝室分配问卷
// @Tags Questionnaire
// @Summary 用id查询寝室分配问卷
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query admin.Questionnaire true "用id查询寝室分配问卷"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionnaire/findQuestionnaire [get]
func (questionnaireApi *QuestionnaireApi) FindQuestionnaire(c *gin.Context) {
	ID := c.Query("ID")
	if requestionnaire, err := questionnaireService.GetQuestionnaire(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requestionnaire": requestionnaire}, c)
	}
}

// GetQuestionnaireList 分页获取寝室分配问卷列表
// @Tags Questionnaire
// @Summary 分页获取寝室分配问卷列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query adminReq.QuestionnaireSearch true "分页获取寝室分配问卷列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionnaire/getQuestionnaireList [get]
func (questionnaireApi *QuestionnaireApi) GetQuestionnaireList(c *gin.Context) {
	var pageInfo adminReq.QuestionnaireSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := questionnaireService.GetQuestionnaireInfoList(pageInfo); err != nil {
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
