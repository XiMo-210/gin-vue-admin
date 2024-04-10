import service from '@/utils/request'

// @Tags PlatformParam
// @Summary 创建平台参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlatformParam true "创建平台参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /platformParam/createPlatformParam [post]
export const createPlatformParam = (data) => {
  return service({
    url: '/platformParam/createPlatformParam',
    method: 'post',
    data
  })
}

// @Tags PlatformParam
// @Summary 删除平台参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlatformParam true "删除平台参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /platformParam/deletePlatformParam [delete]
export const deletePlatformParam = (params) => {
  return service({
    url: '/platformParam/deletePlatformParam',
    method: 'delete',
    params
  })
}

// @Tags PlatformParam
// @Summary 批量删除平台参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除平台参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /platformParam/deletePlatformParam [delete]
export const deletePlatformParamByIds = (params) => {
  return service({
    url: '/platformParam/deletePlatformParamByIds',
    method: 'delete',
    params
  })
}

// @Tags PlatformParam
// @Summary 更新平台参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlatformParam true "更新平台参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /platformParam/updatePlatformParam [put]
export const updatePlatformParam = (data) => {
  return service({
    url: '/platformParam/updatePlatformParam',
    method: 'put',
    data
  })
}

// @Tags PlatformParam
// @Summary 用id查询平台参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PlatformParam true "用id查询平台参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /platformParam/findPlatformParam [get]
export const findPlatformParam = (params) => {
  return service({
    url: '/platformParam/findPlatformParam',
    method: 'get',
    params
  })
}

// @Tags PlatformParam
// @Summary 分页获取平台参数列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取平台参数列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /platformParam/getPlatformParamList [get]
export const getPlatformParamList = (params) => {
  return service({
    url: '/platformParam/getPlatformParamList',
    method: 'get',
    params
  })
}
