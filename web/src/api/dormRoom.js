import service from '@/utils/request'

// @Tags DormRoom
// @Summary 创建宿舍
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DormRoom true "创建宿舍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dormRoom/createDormRoom [post]
export const createDormRoom = (data) => {
  return service({
    url: '/dormRoom/createDormRoom',
    method: 'post',
    data
  })
}

// @Tags DormRoom
// @Summary 删除宿舍
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DormRoom true "删除宿舍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dormRoom/deleteDormRoom [delete]
export const deleteDormRoom = (params) => {
  return service({
    url: '/dormRoom/deleteDormRoom',
    method: 'delete',
    params
  })
}

// @Tags DormRoom
// @Summary 批量删除宿舍
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除宿舍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dormRoom/deleteDormRoom [delete]
export const deleteDormRoomByIds = (params) => {
  return service({
    url: '/dormRoom/deleteDormRoomByIds',
    method: 'delete',
    params
  })
}

// @Tags DormRoom
// @Summary 更新宿舍
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DormRoom true "更新宿舍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dormRoom/updateDormRoom [put]
export const updateDormRoom = (data) => {
  return service({
    url: '/dormRoom/updateDormRoom',
    method: 'put',
    data
  })
}

// @Tags DormRoom
// @Summary 用id查询宿舍
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DormRoom true "用id查询宿舍"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dormRoom/findDormRoom [get]
export const findDormRoom = (params) => {
  return service({
    url: '/dormRoom/findDormRoom',
    method: 'get',
    params
  })
}

// @Tags DormRoom
// @Summary 分页获取宿舍列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取宿舍列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dormRoom/getDormRoomList [get]
export const getDormRoomList = (params) => {
  return service({
    url: '/dormRoom/getDormRoomList',
    method: 'get',
    params
  })
}

export const dormAlloc = (params) => {
  return service({
    url: '/dormRoom/alloc',
    method: 'get',
    params
  })
}

export const getDormStudents = (params) => {
  return service({
    url: '/dormRoom/student',
    method: 'get',
    params
  })
}

export const remind = (params) => {
  return service({
    url: '/dormRoom/remind',
    method: 'get',
    params
  })
}
