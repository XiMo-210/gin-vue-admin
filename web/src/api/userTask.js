import service from '@/utils/request'

export const getUserTaskCondition = (params) => {
  return service({
    url: '/userTask/getUserTaskCondition',
    method: 'get',
    params
  })
}

// @Tags UserTask
// @Summary 创建用户任务记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserTask true "创建用户任务记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userTask/createUserTask [post]
export const createUserTask = (data) => {
  return service({
    url: '/userTask/createUserTask',
    method: 'post',
    data
  })
}

// @Tags UserTask
// @Summary 删除用户任务记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserTask true "删除用户任务记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userTask/deleteUserTask [delete]
export const deleteUserTask = (params) => {
  return service({
    url: '/userTask/deleteUserTask',
    method: 'delete',
    params
  })
}

// @Tags UserTask
// @Summary 批量删除用户任务记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户任务记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userTask/deleteUserTask [delete]
export const deleteUserTaskByIds = (params) => {
  return service({
    url: '/userTask/deleteUserTaskByIds',
    method: 'delete',
    params
  })
}

// @Tags UserTask
// @Summary 更新用户任务记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserTask true "更新用户任务记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userTask/updateUserTask [put]
export const updateUserTask = (data) => {
  return service({
    url: '/userTask/updateUserTask',
    method: 'put',
    data
  })
}

// @Tags UserTask
// @Summary 用id查询用户任务记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserTask true "用id查询用户任务记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userTask/findUserTask [get]
export const findUserTask = (params) => {
  return service({
    url: '/userTask/findUserTask',
    method: 'get',
    params
  })
}

// @Tags UserTask
// @Summary 分页获取用户任务记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户任务记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userTask/getUserTaskList [get]
export const getUserTaskList = (params) => {
  return service({
    url: '/userTask/getUserTaskList',
    method: 'get',
    params
  })
}
