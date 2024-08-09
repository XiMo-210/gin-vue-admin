import service from '@/utils/request'

// @Tags Questionnaire
// @Summary 创建寝室分配问卷
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Questionnaire true "创建寝室分配问卷"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /questionnaire/createQuestionnaire [post]
export const createQuestionnaire = (data) => {
  return service({
    url: '/questionnaire/createQuestionnaire',
    method: 'post',
    data
  })
}

// @Tags Questionnaire
// @Summary 删除寝室分配问卷
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Questionnaire true "删除寝室分配问卷"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionnaire/deleteQuestionnaire [delete]
export const deleteQuestionnaire = (params) => {
  return service({
    url: '/questionnaire/deleteQuestionnaire',
    method: 'delete',
    params
  })
}

// @Tags Questionnaire
// @Summary 批量删除寝室分配问卷
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除寝室分配问卷"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionnaire/deleteQuestionnaire [delete]
export const deleteQuestionnaireByIds = (params) => {
  return service({
    url: '/questionnaire/deleteQuestionnaireByIds',
    method: 'delete',
    params
  })
}

// @Tags Questionnaire
// @Summary 更新寝室分配问卷
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Questionnaire true "更新寝室分配问卷"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionnaire/updateQuestionnaire [put]
export const updateQuestionnaire = (data) => {
  return service({
    url: '/questionnaire/updateQuestionnaire',
    method: 'put',
    data
  })
}

// @Tags Questionnaire
// @Summary 用id查询寝室分配问卷
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Questionnaire true "用id查询寝室分配问卷"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionnaire/findQuestionnaire [get]
export const findQuestionnaire = (params) => {
  return service({
    url: '/questionnaire/findQuestionnaire',
    method: 'get',
    params
  })
}

// @Tags Questionnaire
// @Summary 分页获取寝室分配问卷列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取寝室分配问卷列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionnaire/getQuestionnaireList [get]
export const getQuestionnaireList = (params) => {
  return service({
    url: '/questionnaire/getQuestionnaireList',
    method: 'get',
    params
  })
}
