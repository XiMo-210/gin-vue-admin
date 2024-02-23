import service from '@/utils/request'

// @Tags TaskStage
// @Summary 创建任务阶段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskStage true "创建任务阶段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /taskStage/createTaskStage [post]
export const createTaskStage = (data) => {
  return service({
    url: '/taskStage/createTaskStage',
    method: 'post',
    data
  })
}

// @Tags TaskStage
// @Summary 删除任务阶段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskStage true "删除任务阶段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskStage/deleteTaskStage [delete]
export const deleteTaskStage = (params) => {
  return service({
    url: '/taskStage/deleteTaskStage',
    method: 'delete',
    params
  })
}

// @Tags TaskStage
// @Summary 批量删除任务阶段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除任务阶段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskStage/deleteTaskStage [delete]
export const deleteTaskStageByIds = (params) => {
  return service({
    url: '/taskStage/deleteTaskStageByIds',
    method: 'delete',
    params
  })
}

// @Tags TaskStage
// @Summary 更新任务阶段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskStage true "更新任务阶段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /taskStage/updateTaskStage [put]
export const updateTaskStage = (data) => {
  return service({
    url: '/taskStage/updateTaskStage',
    method: 'put',
    data
  })
}

// @Tags TaskStage
// @Summary 用id查询任务阶段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TaskStage true "用id查询任务阶段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /taskStage/findTaskStage [get]
export const findTaskStage = (params) => {
  return service({
    url: '/taskStage/findTaskStage',
    method: 'get',
    params
  })
}

// @Tags TaskStage
// @Summary 分页获取任务阶段列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取任务阶段列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /taskStage/getTaskStageList [get]
export const getTaskStageList = (params) => {
  return service({
    url: '/taskStage/getTaskStageList',
    method: 'get',
    params
  })
}
