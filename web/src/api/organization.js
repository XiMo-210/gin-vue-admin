import service from '@/utils/request'

// @Tags Organization
// @Summary 创建组织社团
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Organization true "创建组织社团"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /organization/createOrganization [post]
export const createOrganization = (data) => {
  return service({
    url: '/organization/createOrganization',
    method: 'post',
    data
  })
}

// @Tags Organization
// @Summary 删除组织社团
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Organization true "删除组织社团"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /organization/deleteOrganization [delete]
export const deleteOrganization = (params) => {
  return service({
    url: '/organization/deleteOrganization',
    method: 'delete',
    params
  })
}

// @Tags Organization
// @Summary 批量删除组织社团
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除组织社团"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /organization/deleteOrganization [delete]
export const deleteOrganizationByIds = (params) => {
  return service({
    url: '/organization/deleteOrganizationByIds',
    method: 'delete',
    params
  })
}

// @Tags Organization
// @Summary 更新组织社团
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Organization true "更新组织社团"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /organization/updateOrganization [put]
export const updateOrganization = (data) => {
  return service({
    url: '/organization/updateOrganization',
    method: 'put',
    data
  })
}

// @Tags Organization
// @Summary 用id查询组织社团
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Organization true "用id查询组织社团"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /organization/findOrganization [get]
export const findOrganization = (params) => {
  return service({
    url: '/organization/findOrganization',
    method: 'get',
    params
  })
}

// @Tags Organization
// @Summary 分页获取组织社团列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取组织社团列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /organization/getOrganizationList [get]
export const getOrganizationList = (params) => {
  return service({
    url: '/organization/getOrganizationList',
    method: 'get',
    params
  })
}

export const createOrganizationByAdmin = (data) => {
  return service({
    url: '/organization/createOrganizationByAdmin',
    method: 'post',
    data
  })
}

export const updateOrganizationByAdmin = (data) => {
  return service({
    url: '/organization/updateOrganizationByAdmin',
    method: 'put',
    data
  })
}

export const findOrganizationByAdmin = (params) => {
  return service({
    url: '/organization/findOrganizationByAdmin',
    method: 'get',
    params
  })
}
