import service from '@/utils/request'

// @Tags Information
// @Summary 创建信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "创建信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /information/createInformation [post]
export const createInformation = (data) => {
  return service({
    url: '/information/createInformation',
    method: 'post',
    data
  })
}

// @Tags Information
// @Summary 删除信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "删除信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /information/deleteInformation [delete]
export const deleteInformation = (params) => {
  return service({
    url: '/information/deleteInformation',
    method: 'delete',
    params
  })
}

// @Tags Information
// @Summary 批量删除信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /information/deleteInformation [delete]
export const deleteInformationByIds = (params) => {
  return service({
    url: '/information/deleteInformationByIds',
    method: 'delete',
    params
  })
}

// @Tags Information
// @Summary 更新信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Information true "更新信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /information/updateInformation [put]
export const updateInformation = (data) => {
  return service({
    url: '/information/updateInformation',
    method: 'put',
    data
  })
}

// @Tags Information
// @Summary 用id查询信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Information true "用id查询信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /information/findInformation [get]
export const findInformation = (params) => {
  return service({
    url: '/information/findInformation',
    method: 'get',
    params
  })
}

// @Tags Information
// @Summary 分页获取信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /information/getInformationList [get]
export const getInformationList = (params) => {
  return service({
    url: '/information/getInformationList',
    method: 'get',
    params
  })
}
