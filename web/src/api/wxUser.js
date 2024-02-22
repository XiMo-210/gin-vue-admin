import service from '@/utils/request'

// @Tags WxUser
// @Summary 创建用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WxUser true "创建用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /wxUser/createWxUser [post]
export const createWxUser = (data) => {
  return service({
    url: '/wxUser/createWxUser',
    method: 'post',
    data
  })
}

// @Tags WxUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WxUser true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wxUser/deleteWxUser [delete]
export const deleteWxUser = (params) => {
  return service({
    url: '/wxUser/deleteWxUser',
    method: 'delete',
    params
  })
}

// @Tags WxUser
// @Summary 批量删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wxUser/deleteWxUser [delete]
export const deleteWxUserByIds = (params) => {
  return service({
    url: '/wxUser/deleteWxUserByIds',
    method: 'delete',
    params
  })
}

// @Tags WxUser
// @Summary 更新用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WxUser true "更新用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wxUser/updateWxUser [put]
export const updateWxUser = (data) => {
  return service({
    url: '/wxUser/updateWxUser',
    method: 'put',
    data
  })
}

// @Tags WxUser
// @Summary 用id查询用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WxUser true "用id查询用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wxUser/findWxUser [get]
export const findWxUser = (params) => {
  return service({
    url: '/wxUser/findWxUser',
    method: 'get',
    params
  })
}

// @Tags WxUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wxUser/getWxUserList [get]
export const getWxUserList = (params) => {
  return service({
    url: '/wxUser/getWxUserList',
    method: 'get',
    params
  })
}
